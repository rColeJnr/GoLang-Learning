package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	pb "protobuf/datafiles"
)

const (
	address = "localhost:1205"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v\n", err)
	}

	defer conn.Close()
	c := pb.NewMoneyTransactionClient(conn)

	// Prepare data. Get this from clients like Frontend or APP
	from := "1234"
	to := "5678"
	amount := float32(1205.2000)

	// Contact the server and print out its response.
	r, err := c.MakeTransaction(context.Background(), &pb.TransactionRequest{From: from, To: to, Amount: amount})
	if err != nil {
		log.Fatalf("could not transact: %v\n", err)
	}
	log.Printf("Transaction confirmed: %5t\n", r.Confirmation)
}
