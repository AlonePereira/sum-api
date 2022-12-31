package main

import (
	"log"
	"net"

	pb "github.com/AlonePereira/sum-api/calculator/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:5001"

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v\n", err)
	}

}
