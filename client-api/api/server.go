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
	"google.golang.org/grpc"
	"t_task/client-api/config"
	"t_task/client-api/trace"
	"t_task/proto"
)

type Servable interface {
	Serve()
}

// Server contains the values needed for the server
type Server struct {
	Config *config.Config
	Router *mux.Router
	proto.PDServiceClient
}

func CreateServer(config *config.Config) *Server {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}
	client := proto.NewPDServiceClient(conn)
	s := Server{
		Config:          config,
		Router:          mux.NewRouter(),
		PDServiceClient: client,
	}
	return &s
}

// Serve runs the server
func (s *Server) Serve() {

	closer, err := trace.InitGlobalTracer(*s.Config)
	if err != nil {
		log.Fatal(context.Background(), "Failed to init global tracer: ", err)
	} else {
		defer closer.Close()
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
