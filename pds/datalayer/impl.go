package datalayer

import (
	"errors"
	"fmt"
	"log"

	"t_task/proto"
)

// InMemory implementation of the datalayer
type InMemoryPortRepository struct {
	DB map[string]proto.Port
}

// Creates new instance of InMemoryRepository with given DB as map[string]Port
func CreateInMemoryPortRepository(db map[string]proto.Port) *InMemoryPortRepository {
	return &InMemoryPortRepository{DB: db}
}

// Stores/Updates Proto in database
func (r *InMemoryPortRepository) Create(p proto.Port) (*proto.Port, error) {
	log.Println("Creating new Port instance in memory...")
	if r.DB == nil {
		return nil, errors.New("something wrong with database")
	}
	if _, inDB := r.DB[p.ID]; inDB {
		log.Printf("Port #%s is already in database. Overwriting the KEY.", p.ID)
		// If port is already in database -- update it.
		r.DB[p.ID] = p
		return &p, nil
	}
	r.DB[p.ID] = p
	log.Printf("Successfully added new Port instance into DB. ID: %s", p.ID)
	return &p, nil
}

// Deletes Proto from Database with given ID
func (r *InMemoryPortRepository) Delete(id string) (*proto.Port, error) {
	log.Printf("Deleting Port instance by ID: %s", id)
	if r.DB == nil {
		return nil, errors.New("something wrong with database")
	}
	if val, ok := r.DB[id]; ok {
		delete(r.DB, id)
		log.Printf("Successfully deleted Port #%s", id)
		return &val, nil
	}
	msg := fmt.Sprintf("no such entry in database: %s", id)
	log.Println(msg)
	return nil, errors.New(msg)
}

// Lists all Ports from Database. Should we handle some kind of a pagination there ?
func (r *InMemoryPortRepository) List() ([]*proto.Port, error) {
	if r.DB == nil {
		return nil, errors.New("something wrong with database")
	}
	log.Println("Listing all Ports available in database zzz..")
	ports := make([]*proto.Port, 0, len(r.DB))
	for k := range r.DB {
		port := r.DB[k]
		ports = append(ports, &port)
	}
	return ports, nil
}

// Gets one Port with given ID from database
func (r *InMemoryPortRepository) FindOne(id string) (*proto.Port, error) {
	if r.DB == nil {
		return nil, errors.New("something wrong with database")
	}
	log.Printf("Looking for Port #%s in database.", id)
	if port, inDB := r.DB[id]; inDB {
		return &port, nil
	}
	msg := errors.New(fmt.Sprintf("port with this ID(%s) does not exist", id))
	log.Println(msg)
	return nil, msg
}
