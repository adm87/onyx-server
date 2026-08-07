package main

import (
	"log"

	"github.com/adm87/onyx-server/internal/services/auth"
)

func main() {
	if err := auth.Run(); err != nil {
		log.Fatalf("exited with error: %v", err)
	}
}
