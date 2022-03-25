package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

type Controller interface {
	Mux() *mux.Router
}

type Server struct {
	controller Controller
	server     http.Server
}

func New(controller Controller, address string) Server {
	return Server{
		controller: controller,
		server: http.Server{
			Addr:    address,
			Handler: controller.Mux(),
		},
	}
}

func (c *Server) ServeHTTP() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := c.server.ListenAndServe()

		if err != http.ErrServerClosed {
			log.Printf("Server terminated unexpectedly: %v", err)
		}
	}()

	log.Println("Server started.")

	<-done
	log.Println("Server closed.")

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(5)*time.Second)
	defer func() {
		cancel()
	}()

	err := c.server.Shutdown(ctx)
	if err != nil {
		log.Fatalf("Server shutdown failed: %+v", err)
	}
	log.Println("Server exited properly.")
}
