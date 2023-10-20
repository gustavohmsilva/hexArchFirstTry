package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"

	"github.com/gustavohmsilva/ports/internal/adapters/handler"
	"github.com/gustavohmsilva/ports/internal/adapters/repository"
	"github.com/gustavohmsilva/ports/internal/core/port"
	"github.com/gustavohmsilva/ports/internal/core/service"
)

const (
	EnvServerAddress = "SERVER_ADDRESS"
)

var (
	harborService    port.HarborServiceIface
	harborRepository port.HarborRepositoryIface
)

func main() {
	// retrieve web address from container environment variables
	serverAddress, ok := os.LookupEnv(EnvServerAddress)
	if !ok {
		log.Print("failed to get the server address from environment variables")
	}

	// Initialize Database
	harborRepository = repository.NewHarbor()

	// Initialize service with given database
	harborService = service.NewHarbor(harborRepository)

	// Create the router, handlers, and take care of server initialization
	router := newRouter()
	server := newServer(serverAddress, router)
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Print(err.Error())
		}
	}()

	// graceful shutdown
	waitForShutdown(server)
}

// waitForShutdown graceful shutdown of the server in case of OS signal
func waitForShutdown(server *http.Server) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Print("failed to gracefully shut down server")
	} else {
		log.Print("closed server gracefully")
	}

}

// newServer initiates the http server with received routes
func newServer(addr string, r *mux.Router) *http.Server {
	return &http.Server{
		Addr:    addr,
		Handler: r,
	}
}

// newRouter define the accepted endpoints and methods for the server
func newRouter() *mux.Router {
	r := mux.NewRouter()
	h := handler.NewHarbor(harborService)

	r.HandleFunc("/api/harbor", h.CreateHarbor).Methods(http.MethodPost)
	r.HandleFunc("/api/harbor", h.UpdateHarbor).Methods(http.MethodPut)
	return r
}
