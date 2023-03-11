/*
The net/http package provides all the features we need to write Http clients and servers. It gives
us the capability to send reqeust to other servers communicating using HTTP protocol
as well as the ability to run a HTTP server that can route request to separate GO files.
*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloWorldResponse struct {
	// change the output field to be `message`
	Message string `json:"message"`
	// do not output this field
	Author string `json:"-"`
	// do not output if value isEmpty
	Date string `json:",omitempty"`
	// convert output to a string and rename "id"
	Id int `json:"id, string"`
}

func main() {
	port := 1205

	/*The first thing we are doing is calling the HandleFunc method on the http package. The HandleFunc method creates a Handler
	type on the DefaultServeMux handler, mapping the path passed in the first parameter to the function in the second parameter:*/
	// func HandleFunc(patter string, handler func(ResponseWriter, *Request))
	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Server starting on port %v\n", port)
	// here we start the server. pssing the network address: 1205, this means we would like to bind the server
	// to all available IP addresses on port 8080

	// ListenAndServer blocks if the server starts correctly we will never exit on a succeessful start
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))

}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{Message: "HelloWorld"}
	encoder := json.NewEncoder(w)
	encoder.Encode(&response)
	data, err := json.Marshal(response)
	if err != nil {
		panic("Ooops")
	}

	fmt.Fprint(w, string(data))
}

/* `panic` causes normal exec to stop and all deferred function call in the Go routine are exec,
the program will then crash with a log message.*/
