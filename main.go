package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	pbf "micro/consignment"
	"net"
)

const (
	port = ":50051"
)

type IRepository interface {
	Create(*pbf.Consignment) (*pbf.Consignment, error)
	GetAll() []*pbf.Consignment
}

type Repository struct {
	consignments []*pbf.Consignment
}

func (repo *Repository) Create(consignment *pbf.Consignment) (*pbf.Consignment, error) {
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

func (repo *Repository) GetAll() []*pbf.Consignment {
	return repo.consignments
}

type service struct {
	repo IRepository
}

func (s *service) CreateConsignment(ctx context.Context, req *pbf.Consignment) (*pbf.Response, error) {
	consignment, err := s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	return &pbf.Response{Created: true, Consignment: consignment}, nil
}

func (s *service) GetConsignments(ctx context.Context, req *pbf.GetRequest) (*pbf.Response, error) {
	consignments := s.repo.GetAll()
	return &pbf.Response{Consignments: consignments}, nil
}

func main() {
	repo := &Repository{}

	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pbf.RegisterShippingServiceServer(s, &service{repo})

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
