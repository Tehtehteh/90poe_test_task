package api

import (
	"context"
	"log"

	"t_task/proto"

	"google.golang.org/grpc"
)

type PDClient struct {
	c proto.PDServiceClient
}

type IPDClient interface {
	List(ctx context.Context) (*proto.Ports, error)
	Insert(ctx context.Context, p *proto.Port) (*proto.Port, error)
	DeleteByID(ctx context.Context, id string) (*proto.Port, error)
	GetByID(ctx context.Context, id string) (*proto.Port, error)
}

func NewPDClient(addr string) *PDClient {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}
	client := proto.NewPDServiceClient(conn)
	return &PDClient{c: client}
}

func (c *PDClient) List(ctx context.Context) (*proto.Ports, error) {
	return c.c.List(ctx, &proto.NoParams{})
}

func (c *PDClient) Insert(ctx context.Context, p *proto.Port) (*proto.Port, error) {
	return c.c.Insert(ctx, p)
}

func (c *PDClient) DeleteByID(ctx context.Context, id string) (*proto.Port, error) {
	return c.c.DeleteByID(ctx, &proto.DeletePortByIDMsg{ID: id})
}

func (c *PDClient) GetByID(ctx context.Context, id string) (*proto.Port, error) {
	return c.c.GetByID(ctx, &proto.GetByIDMsg{ID: id})
}
