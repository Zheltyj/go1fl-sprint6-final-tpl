package main

import (
	"log"
	"os"

	"github.com/Zheltyj/go1fl-sprint6-final-tpl/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "server: ", log.LstdFlags|log.Lshortfile)
	serv := server.NewServer(logger)

	if err := serv.HttpServer.ListenAndServe(); err != nil {
		logger.Fatalf("Server error: %v", err)
	}
}
