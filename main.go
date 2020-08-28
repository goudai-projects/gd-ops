package main

import (
	"fmt"
	v1 "github.com/goudai-projects/gd-ops/api/v1"
	"github.com/goudai-projects/gd-ops/app"
	"github.com/goudai-projects/gd-ops/config"
	"github.com/goudai-projects/gd-ops/utils"
)

func main() {
	fmt.Println("Hello GD DevOps.")
	cfg := config.GetConfig()
	srv := app.NewServer(cfg)
	utils.InitI18n()
	v1.Init(srv, srv.Router())
	srv.Start()
}
