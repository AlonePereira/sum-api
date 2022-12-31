package main

import (
	"context"
	"log"

	pb "github.com/AlonePereira/sum-api/calculator/proto"
)

func doCalculate(client pb.CalculatorServiceClient) {
	resp, err := client.Calculate(context.Background(), &pb.CalculatorRequest{
		FirstNumber:  10,
		SecondNumber: 15,
	})

	if err != nil {
		log.Fatalf("Erro Calculate %v\n", err)
	}

	log.Printf("Caculate for %v", resp.Result)
}
