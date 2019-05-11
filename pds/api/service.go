package api

import (
	"context"
	"errors"
	"log"

	"t_task/pds/datalayer"
	"t_task/proto"
)

func CreateNewPDService(repo datalayer.IPortRepository) *PDService {
	return &PDService{
		PortsRepository: repo,
	}
}

type PDService struct {
	PortsRepository datalayer.IPortRepository
}

func (s *PDService) List(ctx context.Context, _ *proto.NoParams) (*proto.Ports, error) {
	ports, err := s.PortsRepository.List()
	if err != nil {
		return nil, err
	}
	resp := proto.Ports{Response: ports}
	return &resp, nil
}

func (s *PDService) Insert(ctx context.Context, p *proto.Port) (*proto.Port, error) {
	if p == nil {
		return nil, errors.New("port is nil, aborting")
	}
	res, err := s.PortsRepository.Create(*p)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *PDService) DeleteByID(ctx context.Context, param *proto.DeletePortByIDMsg) (*proto.Port, error) {
	if param != nil {
		log.Printf("Received following delete by ID param: %s", param.ID)
	}
	return nil, nil
}
