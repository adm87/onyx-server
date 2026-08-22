package main

import (
	"log"

	"github.com/adm87/onyx-server/cmd/auth-service/internal/server"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatalf("auth-service exited with err: %v", err)
	}
}
