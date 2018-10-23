package main

import (
	"context"
	"encoding/json"
	"google.golang.org/grpc"
	"io/ioutil"
	"log"
	pbf "micro/consignment"
	"os"
)

const (
	address         = "localhost:50051"
	defaultFilename = "consingment.json"
)

func ParseFile(file string) (*pbf.Consignment, error) {
	var consignment *pbf.Consignment

	data, err := ioutil.ReadFile(file)

	if err != nil {
		return nil, err
	}

	json.Unmarshal(data, &consignment)
	return consignment, nil
}

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure())

	defer conn.Close()

	client := pbf.NewShippingServiceClient(conn)

	if err != nil {
		log.Fatalf("Connection error: %v", err)
	}

	file := defaultFilename

	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := ParseFile(file)

	if err != nil {
		log.Fatalf("Parse file error: %v", err)
	}

	r, err := client.CreateConsignment(context.Background(), consignment)

	if err != nil {
		log.Fatalf("Error create consingment: %v", err)
	}

	log.Printf("Created: %v", r.Created)

	getAll, err := client.GetConsignments(context.Background(), &pbf.GetRequest{})

	if err != nil {
		log.Fatalf("Error get consingment: %v", err)
	}

	for _, v := range getAll.Consignments {
		log.Println(v)
	}
}
