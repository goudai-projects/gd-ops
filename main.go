package main

import (
	"fmt"
	"github.com/goudai-projects/gd-ops/pkg/k8s"
)

func main() {
	instance := k8s.GetInstance()
	instance2 := k8s.GetInstance()
	instance3 := k8s.GetInstance()
	println(instance)
	println(instance2)
	println(instance3)
	fmt.Println("Hello GD DevOps.")
	//cfg := config.GetConfig()
	//db.Init(cfg.Database.DSN)
	//server.Init(cfg.Server.Address, cfg.Server.Port)
}
