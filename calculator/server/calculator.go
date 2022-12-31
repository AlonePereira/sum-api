package service

import (
	"context"
	"log"

	pb "github.com/AlonePereira/sum-api/greet/proto"
)

type server struct {
	pb.CalculatorServiceServer
}

func (s *server) Calculate(ctx context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	log.Printf("Request %v", in)
	return &pb.CalculatorResponse{}, nil
}
