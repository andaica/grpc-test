package main

import (
	"context"
	"log"
	"time"

	pb "grpc-test/interface"

	"google.golang.org/grpc"
)

const (
	host = "localhost"
	port = "8888"
)

func main() {

	conn, err := grpc.Dial(host+":"+port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("cannot connect: %v", err)
	}
	defer conn.Close()

	myClient := pb.NewUserServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := pb.AuthInfo{
		Email:    "annt@ows.vn",
		Password: "12345678",
	}

	res, err := myClient.CreateUser(ctx, &req)
	if err != nil {
		log.Fatalf("error call grpc: %v", err)
	}

	log.Printf("response: %v", res)

}
