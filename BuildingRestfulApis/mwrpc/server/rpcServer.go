package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Args struct {
}
type TimeServer int64

/*
GiveServerTime takes the Args object as the first argument and a reply pointer
object
It sets the reply pointer object but does not return anything except an error
The Args struct here has no fields because this server is not expecting the client to
send any arguments
*/
func (t *TimeServer) GiveServerTime(args *Args, reply *int64) error {
	// fill reply pointer to send the data back
	*reply = time.Now().Unix()
	return nil
}

func main() {
	// create a new rpc server
	timeServer := new(TimeServer)
	// Register server
	rpc.Register(timeServer) //reg server to rpc
	rpc.HandleHTTP()         // register an HTTP handler for RPC messages to `default server`
	// Listen for reqeusts on port
	l, e := net.Listen("tcp", ":1205") // start tcp server taht listens on port 1205
	if e != nil {
		log.Fatal("listenn error:", e)
	}
	http.Serve(l, nil)
}
