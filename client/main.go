package main

import (
	"context"
	"log"

	"github.com/rodrigorrch/challengs_platform/framework/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)

	User(err, client)
}

func User(err error, client pb.UserServiceClient) {
	request := &pb.UserRequest{
		Name:     "Rodrigo",
		Email:    "rodrigo.rrch@gmail.com",
		Password: "123456",
	}

	res, err := client.CreateUser(context.Background(), request)
	if err != nil {
		log.Fatalf("Error during execution: %v", err)
	}

	log.Println(res.Token)
}
