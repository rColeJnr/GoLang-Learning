package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	jsonparse "encoding/json"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc/json"

	"github.com/gorilla/rpc"
)

type Args struct {
	Id string
}

type Book struct {
	Id     string `"json:string,omitempty"`
	Name   string `"json:string,omitempty"`
	Author string `"json:string,omitempty"`
}

type JSONServer struct {
}

func (j *JSONServer) GiveBookDetail(r *http.Request, args *Args, reply *Book) error {
	var books []Book

	//  read json file and load data
	raw, reader := ioutil.ReadFile("./books.json")
	if reader != nil {
		log.Println("error:", reader)
		os.Exit(1)
	}

	// unmarshal JSON raw data into books array
	marshaller := jsonparse.Unmarshal(raw, &books)
	if marshaller != nil {
		log.Println("error: ", marshaller)
		os.Exit(1)
	}

	for _, book := range books {
		if book.Id == args.Id {
			// if book found, reply book
			*reply = book
			break
		}
	}
	return nil

}

/*
GiveServerTime takes the Args object as the first argument and a reply pointer
object
It sets the reply pointer object but does not return anything except an error
The Args struct here has no fields because this server is not expecting the client to
send any arguments
// */
// func (t *TimeServer) GiveServerTime(args *Args, reply *int64) error {
// 	// fill reply pointer to send the data back
// 	*reply = time.Now().Unix()
// 	return nil
// }

// func main() {
// 	// create a new rpc server
// 	timeServer := new(TimeServer)
// 	// Register server
// 	rpc.Register(timeServer) //reg server to rpc
// 	rpc.HandleHTTP()         // register an HTTP handler for RPC messages to `default server`
// 	// Listen for reqeusts on port
// 	l, e := net.Listen("tcp", ":1205") // start tcp server taht listens on port 1205
// 	if e != nil {
// 		log.Fatal("listenn error:", e)
// 	}
// 	http.Serve(l, nil)
// }

func main() {
	// create new rpc server
	s := rpc.NewServer()
	// register the type of data requested as json
	s.RegisterCodec(json.NewCodec(), "application/json")
	// register the service
	s.RegisterService(new(JSONServer), "")
	r := mux.NewRouter()
	r.Handle("/rpc", s)
	http.ListenAndServe(":1205", r)
}
