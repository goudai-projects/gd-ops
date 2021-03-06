package log

import (
	"fmt"
	"os"
	"sync/atomic"
	"time"

	"github.com/mattn/go-isatty"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// none is used to disable logging output as well as to disable stack tracing.
const none zapcore.Level = 100

var levelToZap = map[Level]zapcore.Level{
	DebugLevel: zapcore.DebugLevel,
	InfoLevel:  zapcore.InfoLevel,
	WarnLevel:  zapcore.WarnLevel,
	ErrorLevel: zapcore.ErrorLevel,
	FatalLevel: zapcore.FatalLevel,
	NoneLevel:  none,
}

// functions that can be replaced in a test setting
type patchTable struct {
	write       func(ent zapcore.Entry, fields []zapcore.Field) error
	sync        func() error
	exitProcess func(code int)
	errorSink   zapcore.WriteSyncer
}

// function table that can be replaced by tests
var funcs = &atomic.Value{}

func init() {
	// use our defaults for starters so that logging works even before everything is fully configured
	_ = Configure(DefaultOptions())
}

// prepZap is a utility function used by the Configure function.
func prepZap(options *Options) (zapcore.Core, zapcore.Core, zapcore.WriteSyncer, error) {
	encCfg := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "scope",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stack",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeTime:     formatDate,
	}

	var enc zapcore.Encoder
	if options.JSONEncoding {
		enc = zapcore.NewJSONEncoder(encCfg)
	} else {
		// 在终端环境中开启颜色高亮
		if isatty.IsTerminal(os.Stdout.Fd()) || isatty.IsCygwinTerminal(os.Stdout.Fd()) {
			encCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
		}

		enc = zapcore.NewConsoleEncoder(encCfg)
	}

	var rotaterSink zapcore.WriteSyncer
	if options.RotateOutputPath != "" {
		if options.RotationMaxSize == 0 {
			options.RotationMaxSize = defaultRotationMaxSize
		}
		if options.RotationMaxBackups == 0 {
			options.RotationMaxBackups = defaultRotationMaxBackups
		}
		if options.RotationMaxAge == 0 {
			options.RotationMaxAge = defaultRotationMaxAge
		}

		rotaterSink = zapcore.AddSync(&lumberjack.Logger{
			Filename:   options.RotateOutputPath,
			MaxSize:    options.RotationMaxSize,
			MaxBackups: options.RotationMaxBackups,
			MaxAge:     options.RotationMaxAge,
		})
	}

	errSink, closeErrorSink, err := zap.Open(options.ErrorOutputPaths...)
	if err != nil {
		return nil, nil, nil, err
	}

	var outputSink zapcore.WriteSyncer
	if len(options.OutputPaths) > 0 {
		outputSink, _, err = zap.Open(options.OutputPaths...)
		if err != nil {
			closeErrorSink()
			return nil, nil, nil, err
		}
	}

	var sink zapcore.WriteSyncer
	if rotaterSink != nil && outputSink != nil {
		sink = zapcore.NewMultiWriteSyncer(outputSink, rotaterSink)
	} else if rotaterSink != nil {
		sink = rotaterSink
	} else {
		sink = outputSink
	}

	var enabler zap.LevelEnablerFunc = func(lvl zapcore.Level) bool {
		switch lvl {
		case zapcore.ErrorLevel:
			return defaultScope.ErrorEnabled()
		case zapcore.WarnLevel:
			return defaultScope.WarnEnabled()
		case zapcore.InfoLevel:
			return defaultScope.InfoEnabled()
		}
		return defaultScope.DebugEnabled()
	}

	return zapcore.NewCore(enc, sink, zap.NewAtomicLevelAt(zapcore.DebugLevel)),
		zapcore.NewCore(enc, sink, enabler),
		errSink, nil
}

func formatDate(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	t = t.UTC()
	year, month, day := t.Date()
	hour, minute, second := t.Clock()
	micros := t.Nanosecond() / 1000

	buf := make([]byte, 27)

	buf[0] = byte((year/1000)%10) + '0'
	buf[1] = byte((year/100)%10) + '0'
	buf[2] = byte((year/10)%10) + '0'
	buf[3] = byte(year%10) + '0'
	buf[4] = '-'
	buf[5] = byte((month)/10) + '0'
	buf[6] = byte((month)%10) + '0'
	buf[7] = '-'
	buf[8] = byte((day)/10) + '0'
	buf[9] = byte((day)%10) + '0'
	buf[10] = 'T'
	buf[11] = byte((hour)/10) + '0'
	buf[12] = byte((hour)%10) + '0'
	buf[13] = ':'
	buf[14] = byte((minute)/10) + '0'
	buf[15] = byte((minute)%10) + '0'
	buf[16] = ':'
	buf[17] = byte((second)/10) + '0'
	buf[18] = byte((second)%10) + '0'
	buf[19] = '.'
	buf[20] = byte((micros/100000)%10) + '0'
	buf[21] = byte((micros/10000)%10) + '0'
	buf[22] = byte((micros/1000)%10) + '0'
	buf[23] = byte((micros/100)%10) + '0'
	buf[24] = byte((micros/10)%10) + '0'
	buf[25] = byte((micros)%10) + '0'
	buf[26] = 'Z'

	enc.AppendString(string(buf))
}

func updateScopes(options *Options) error {
	// snapshot what's there
	allScopes := Scopes()

	// update the output levels of all listed scopes
	if err := processLevels(allScopes, options.Levels, func(s *Scope, l Level) { s.SetOutputLevel(l) }); err != nil {
		return err
	}

	// update the stack tracing levels of all listed scopes
	if err := processLevels(allScopes, options.StackTraceLevels, func(s *Scope, l Level) { s.SetStackTraceLevel(l) }); err != nil {
		return err
	}

	// update the caller location setting of all listed scopes
	// sc := strings.Split(options.logCallers, ",")
	// for _, s := range sc {
	// 	if s == "" {
	// 		continue
	// 	}

	// 	if s == OverrideScopeName {
	// 		// ignore everything else and just apply the override value
	// 		for _, scope := range allScopes {
	// 			scope.SetLogCallers(true)
	// 		}

	// 		return nil
	// 	}

	// 	if scope, ok := allScopes[s]; ok {
	// 		scope.SetLogCallers(true)
	// 	} else {
	// 		return fmt.Errorf("unknown scope '%s' specified", s)
	// 	}
	// }

	return nil
}

// processLevels breaks down an argument string into a set of scope & levels and then
// tries to apply the result to the scopes. It supports the use of a global override.
func processLevels(allScopes map[string]*Scope, levels map[string]string, setter func(*Scope, Level)) error {
	// levels := strings.Split(arg, ",")
	for s, sl := range levels {
		l := stringToLevel[sl]

		if scope, ok := allScopes[s]; ok {
			setter(scope, l)
		} else if s == OverrideScopeName {
			// override replaces everything
			for _, scope := range allScopes {
				setter(scope, l)
			}
			return nil
		} else {
			return fmt.Errorf("unknown scope '%s' specified", s)
		}
	}

	return nil
}

// Configure initializes Istio's logging subsystem.
//
// You typically call this once at process startup.
// Once this call returns, the logging system is ready to accept data.
// nolint: staticcheck
func Configure(options *Options) error {
	core, captureCore, errSink, err := prepZap(options)
	if err != nil {
		return err
	}

	if err = updateScopes(options); err != nil {
		return err
	}

	pt := patchTable{
		write: func(ent zapcore.Entry, fields []zapcore.Field) error {
			err := core.Write(ent, fields)
			if ent.Level == zapcore.FatalLevel {
				funcs.Load().(patchTable).exitProcess(1)
			}

			return err
		},
		sync:        core.Sync,
		exitProcess: os.Exit,
		errorSink:   errSink,
	}
	funcs.Store(pt)

	opts := []zap.Option{
		zap.ErrorOutput(errSink),
		zap.AddCallerSkip(1),
	}

	if defaultScope.GetLogCallers() {
		opts = append(opts, zap.AddCaller())
	}

	l := defaultScope.GetStackTraceLevel()
	if l != NoneLevel {
		opts = append(opts, zap.AddStacktrace(levelToZap[l]))
	}

	captureLogger := zap.New(captureCore, opts...)

	// capture global zap logging and force it through our logger
	_ = zap.ReplaceGlobals(captureLogger)

	// capture standard golang "log" package output and force it through our logger
	_ = zap.RedirectStdLog(captureLogger)

	return nil
}

// Sync flushes any buffered log entries.
// Processes should normally take care to call Sync before exiting.
func Sync() error {
	return funcs.Load().(patchTable).sync()
}
