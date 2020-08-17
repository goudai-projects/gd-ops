package server

import (
	"fmt"
	"github.com/goudai-projects/gd-ops/log"
)

func Init(address string, port int) {
	r := NewRouter()
	err := r.Run(fmt.Sprintf("%s:%d", address, port))
	if err != nil {
		log.Error("http server start fail")
	}
}
