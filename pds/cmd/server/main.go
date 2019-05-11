package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"t_task/pds/api"
	"t_task/pds/config"
	"t_task/pds/datalayer"
	"t_task/proto"
)

func main() {
	// Reading .env (if present) and creating config instance
	env, err := config.CreateConfig()
	if err != nil {
		log.Printf("Error creating config instance: %s", err)
		return
	}
	// Creating listener for RPC Service
	listener, err := net.Listen("tcp", env.ListenAddress)
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	repo := datalayer.CreateInMemoryPortRepository(make(map[string]proto.Port))
	service := api.CreateNewPDService(repo)
	proto.RegisterPDServiceServer(srv, service)
	reflection.Register(srv)
	log.Println("Started RPC service on 4040 port.")
	idleConnsClosed := make(chan struct{})
	go func() {
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
		<-stop
		srv.GracefulStop()
		log.Println("Gracefully shutdown service!")
		close(idleConnsClosed)
	}()
	if err = srv.Serve(listener); err != nil {
		panic(err)
	}
	<-idleConnsClosed
}
