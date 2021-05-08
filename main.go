package main

import (
	"conch/server"
	"log"
)

func main() {
	if err := server.NewServer().ListenAndServe(); err != nil {
		log.Fatalf("server run with error: %v:", err)
	}
}
