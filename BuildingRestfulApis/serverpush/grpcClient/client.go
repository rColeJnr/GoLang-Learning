package main

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	pb "serverpush/datafiles"
)

const (
	address = "localhost:1205"
)

// REceiveStream listens to the stream contents and use them
func ReceiveStream(client pb.MoneyTransactionClient, request *pb.TransactionRequest) {
	log.Println("Started listening to the server stream!")
	stream, err := client.MakeTransaction(context.Background(), request)
	if err != nil {
		log.Fatalf("%v.MakeTransaction(_) = _, %v\n", client, err)
	}

	// listen to the stream of messages
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			// if thre are no more messages
			break
		}
		if err != nil {
			log.Fatalf("%v.MakeTransaxtion(_) = _, %v\n", client, err)
		}

		log.Printf("Status: %v, OPeration: %v", response.Status, response.Description)
	}
}

func main() {
	// set up a connection to the server
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connecdt: %v\n", err)
	}
	defer conn.Close()
	client := pb.NewMoneyTransactionClient(conn)

	// Prepare data. GEt this from clients like fron-end or Android app
	from := "1234"
	to := "5667"
	amount := float32(1205.2000)

	// Contact the server and print out its response.
	ReceiveStream(client, &pb.TransactionRequest{From: from, To: to, Amount: amount})
}
