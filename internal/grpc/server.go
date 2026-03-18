package grpc

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func StartGRPC() {

	lis, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()

	log.Println("gRPC server running")

	s.Serve(lis)
}