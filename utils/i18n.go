package utils

import (
	"encoding/json"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"io/ioutil"
	"os"
	"path/filepath"
)

var searchPaths = []string{
	".",
	"..",
	"../..",
	"../../..",
}

var bundle *i18n.Bundle

func InitI18n() {
	bundle = i18n.NewBundle(language.Chinese)
	bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
	var searchedPaths []string
	for _, path := range searchPaths {
		if _, err := os.Stat(path + "/i18n"); !os.IsNotExist(err) {
			searchedPaths = append(searchedPaths, path+"/i18n")
		}
	}
	for _, path := range searchedPaths {
		files, _ := ioutil.ReadDir(path)
		for _, file := range files {
			bundle.MustLoadMessageFile(filepath.Join(path, file.Name()))
		}
	}

}

func Translate(lang []string, msgId string, params map[string]interface{}) string {
	l := i18n.NewLocalizer(bundle, lang...)
	return l.MustLocalize(&i18n.LocalizeConfig{MessageID: msgId, TemplateData: params})
}
