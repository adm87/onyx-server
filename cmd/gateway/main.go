package main

import (
	"strconv"

	"github.com/adm87/onyx-server/pkg/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	addr := cfg.Gateway.Host + ":" + strconv.Itoa(cfg.Gateway.Port)
	println("Starting gateway server on " + addr)
}
