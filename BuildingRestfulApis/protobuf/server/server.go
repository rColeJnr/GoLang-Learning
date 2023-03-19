package main

import (
	"context" // for RPC context
	"log"
	"net"
	pb "protobuf/datafiles"
)
import "google.golang.org/grpc"
import "google.golang.org/grpc/reflection"

const (
	port = ":1205"
)

// server is used to create MOneytransactionserver
type server struct {
}

// the in variable has the RPC request details. Is is basically a struct that maps to the TransactionRequest defined tin the proto file
func (s *server) MakeTransaction(ctx context.Context, in *pb.TransactionRequest) (*pb.TransactionResponse, error) {
	log.Println("Got request for money transfer...")
	log.Printf("Amount: %f, from A/c:%s, to A/c:%s\n", in.Amount, in.From, in.To)
	// Do db logic here....
	return &pb.TransactionResponse{Confirmation: true}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMoneyTransactionServer(s, &server{})
	// REgister reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}