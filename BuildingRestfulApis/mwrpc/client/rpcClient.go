package main

import (
	"log"
	"net/rpc"
)

type Args struct {
}

func main() {
	var reply int64

	args := Args{}

	client, err := rpc.DialHTTP("tcp", "localhost:1205") // dial to the RPC server
	if err != nil {
		log.Fatal("dialing:", err)
	}

	err = client.Call("TimeServer.GiveServerTime", args, &reply) // call the remote func with the name:function format with args and reply with pointer obj
	if err != nil {
		log.Fatal("arith error:", err)
	}
	log.Printf("replay: %d", reply) // print the data returned by the remote server
}
