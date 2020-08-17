package main

import (
	"fmt"
	"github.com/goudai-projects/gd-ops/config"
	"github.com/goudai-projects/gd-ops/db"
	"github.com/goudai-projects/gd-ops/server"
)

func main() {
	fmt.Println("Hello GD DevOps.")
	cfg := config.GetConfig()
	db.Init(cfg.Database.DSN)
	server.Init(cfg.Server.Address, cfg.Server.Port)
}
