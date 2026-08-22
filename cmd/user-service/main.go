package main

import (
	"log"

	"github.com/adm87/onyx-server/cmd/user-service/internal/server"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatalf("exited with err: %v", err)
	}
}
