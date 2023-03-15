package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

// CustomServeMus is a struct which can be a multiplexer
type CustomServeMux struct{}

// this is the function handler to be overridden
func (p *CustomServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		giveRandom(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func giveRandom(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Your random number is: %f", rand.Float32())
}
