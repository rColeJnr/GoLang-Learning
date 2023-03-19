package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	pb "serverpush/datafiles"
	"time"
)

const (
	port      = ":1205"
	noOfSteps = 3
)

type server struct {
}

func (s *server) MakeTransaction(in *pb.TransactionRequest, stream pb.MoneyTransaction_MakeTransactionServer) error {
	log.Printf("Got request for money transfer....\n")
	log.Printf("Amount: $%f, From A/c:%s, To A/c:%s\n", in.Amount, in.From, in.To)
	// SEnd streams here
	for i := 0; i < noOfSteps; i++ {
		// Simulating io using sleep (db operations here)
		//...
		time.Sleep(2 * time.Second)
		// onc task is done, send successful message to client
		if err := stream.Send(&pb.TransactionResponse{Status: "good", Step: int32(i), Description: fmt.Sprintf("Description of step %d,", int32(i))}); err != nil {
			log.Fatalf("%v.Send(%v) = %v", stream, "status", err)
		}
	}
	log.Printf("Successfully transfered amount $%v from %v to %v", in.Amount, in.From, in.To)
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// creare a new grpc server
	s := grpc.NewServer()
	// REgister with proto service
	pb.RegisterMoneyTransactionServer(s, &server{})
	// Register reflection service on gRpc server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
