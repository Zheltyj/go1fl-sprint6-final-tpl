package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Zheltyj/go1fl-sprint6-final-tpl/internal/handlers"
)

type Server struct {
	Logger     *log.Logger
	HttpServer *http.Server
}

func NewServer(logger *log.Logger) *Server {
	router := http.NewServeMux()

	router.HandleFunc("/", handlers.IndexHandler)
	router.HandleFunc("/upload", handlers.UploadHandler)

	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		Logger:     logger,
		HttpServer: httpServer,
	}
}
