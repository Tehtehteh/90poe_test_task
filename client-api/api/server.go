package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"t_task/client-api/config"
	"t_task/client-api/trace"
)

type Servable interface {
	Serve()
}

// Server contains the values needed for the server
type Server struct {
	Config *config.Config
	Router *mux.Router
	Client *PDClient
}

func CreateServer(config *config.Config) *Server {
	client := NewPDClient(config.PDServiceAddr)
	s := Server{
		Config: config,
		Router: mux.NewRouter(),
		Client: client,
	}
	return &s
}

// Serve runs the server
func (s *Server) Serve() {

	closer, err := trace.InitGlobalTracer(*s.Config)
	if err != nil {
		log.Fatal(context.Background(), "Failed to init global tracer: ", err)
	} else {
		defer func() {
			err := closer.Close()
			if err != nil {
				log.Printf("error closing jaeger global tracer: %s", err)
			}
		}()
	}
	s.routes()
	// server config
	httpServer := http.Server{
		Addr:         s.Config.ListenAddress,
		ReadTimeout:  600 * time.Second,
		WriteTimeout: 600 * time.Second,
	}
	log.Printf("Client API started listening on port `%s`.\n", s.Config.ListenAddress)
	http.Handle("/", s.Router)
	idleConnsClosed := make(chan struct{})
	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
		<-stop
		if err := httpServer.Shutdown(context.Background()); err != nil {
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()
	if err := httpServer.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("HTTP server ListenAndServe: %v", err)
	}
	<-idleConnsClosed
}
