package main

import (
	"io"
	"net/http"
)

// hello world, the web server
func MyServer(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello World!\n") // wirte hello World to the response
}

// Whenver the request comes on the /hello route, the handler function is execd
// A ResponseWrite interface is used by an HTTP handler to construct an HTTP response
// http.Request is an object that deals with all the porperties and methods of an HTTP request
