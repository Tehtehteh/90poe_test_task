package datalayer

import "t_task/proto"

type IPortRepository interface {
	Create(proto.Port) (*proto.Port, error)
	Delete(id string) (*proto.Port, error)
	List() ([]*proto.Port, error)
}
