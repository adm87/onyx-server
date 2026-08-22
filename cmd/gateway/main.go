package main

import (
	"log"

	"github.com/adm87/onyx-server/cmd/gateway/internal/server"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatalf("gateway exited with err: %v", err)
	}
}
