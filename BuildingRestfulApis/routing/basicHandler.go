package main

import (
	"io"
	"log"
	"net/http"
)

// hello world, the web server
func MyServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello World!\n") // wirte hello World to the response
}
func main() {
	http.HandleFunc("/hello", MyServer)          // creates /hello route. handleFunc maps an URL to a func
	log.Fatal(http.ListenAndServe(":8000", nil)) // start the sever on given port and return error if somethings goes wrong
}

// Whenver the request comes on the /hello route, the handler function is execd
// A ResponseWrite interface is used by an HTTP handler to construct an HTTP response
// http.Request is an object that deals with all the porperties and methods of an HTTP request
