package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"t_task/pds/api"
	"t_task/pds/datalayer"
	"t_task/proto"
)

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	repo := datalayer.CreateInMemoryPortRepository(make(map[string]proto.Port))
	service := api.CreateNewPDService(repo)
	proto.RegisterPDServiceServer(srv, service)
	reflection.Register(srv)
	log.Println("Started RPC service on 4040 port.")
	if err = srv.Serve(listener); err != nil {
		panic(err)
	}
}
