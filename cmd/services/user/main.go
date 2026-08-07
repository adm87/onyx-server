package main

import (
	"log"

	"github.com/adm87/onyx-server/internal/services/user"
)

func main() {
	if err := user.Run(); err != nil {
		log.Fatalf("exited with error: %v", err)
	}
}
