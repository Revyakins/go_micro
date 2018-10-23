package main

import pbf "micro/consignment"

const (
	port = ":50051"
)

type IRepository interface {
	Create(*pbf.Consignment) (*pbf.Consignment, error)
}

type Repository struct {
	consignments []*pbf.Consignment
}

func (repo Repository) Create(consignment *pbf.Consignment) (*pbf.Consignment, error) {
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

func main() {

}
