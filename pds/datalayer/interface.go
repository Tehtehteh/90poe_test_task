package datalayer

import "t_task/proto"

// Interface for accessing underlying datalayer
type IPortRepository interface {
	Create(proto.Port) (*proto.Port, error)
	Delete(id string) (*proto.Port, error)
	List() ([]*proto.Port, error)
	FindOne(id string) (*proto.Port, error)
}
