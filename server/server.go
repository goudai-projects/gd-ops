package server

import (
	"fmt"
	"github.com/goudai-projects/gdops/config"
	"github.com/goudai-projects/gdops/log"
)

func Init() {
	cfg := config.GetConfig()
	r := NewRouter()
	err := r.Run(fmt.Sprintf("%s:%d", cfg.Server.Address, cfg.Server.Port))
	if err != nil {
		log.Error("http server start fail")
	}
}
