package datalayer

import (
	"errors"
	"fmt"
	"t_task/proto"
)

type InMemoryPortRepository struct {
	DB map[string]proto.Port
}

func CreateInMemoryPortRepository(db map[string]proto.Port) *InMemoryPortRepository {
	return &InMemoryPortRepository{DB: db}
}

func (r *InMemoryPortRepository) Create(p proto.Port) (*proto.Port, error) {
	if r.DB == nil {
		return nil, errors.New("something wrong with database")
	}
	if _, inDB := r.DB[p.ID]; inDB {
		msg := fmt.Sprintf("port is already present in database with id: %s", p.ID)
		return nil, errors.New(msg)
	}
	r.DB[p.ID] = p
	return &p, nil
}

func (r *InMemoryPortRepository) Delete(id string) (*proto.Port, error) {
	if r.DB == nil {
		return nil, errors.New("something wrong with database")
	}
	if val, ok := r.DB[id]; ok {
		delete(r.DB, id)
		return &val, nil
	}
	msg := fmt.Sprintf("no such entry in database: %s", id)
	return nil, errors.New(msg)
}

func (r *InMemoryPortRepository) List() ([]*proto.Port, error) {
	if r.DB == nil {
		return nil, errors.New("something wrong with database")
	}
	ports := make([]*proto.Port, 0, len(r.DB))
	for _, v := range r.DB {
		ports = append(ports, &v)
	}
	return ports, nil
}
