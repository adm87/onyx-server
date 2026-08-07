package main

import (
	"log"

	"github.com/adm87/onyx-server/internal/gateway"
)

func main() {
	if err := gateway.Run(); err != nil {
		log.Fatalf("exited with error: %v", err)
	}
}
