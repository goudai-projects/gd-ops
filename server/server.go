package server

import (
	"fmt"
	"github.com/goudai-projects/gd-ops/config"
	"github.com/goudai-projects/gd-ops/log"
)

func Init() {
	cfg := config.GetConfig()
	r := NewRouter()
	err := r.Run(fmt.Sprintf("%s:%d", cfg.Server.Address, cfg.Server.Port))
	if err != nil {
		log.Error("http server start fail")
	}
}
