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

func main() {
	// Any struct that has severHttp func can be a multiplexer
	newMux := http.NewServeMux()

	newMux.HandleFunc("/randomFloat", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, rand.Float64())
	})

	newMux.HandleFunc("/randomInt", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, rand.Intn(100))
	})
	// mux := &CustomServeMux{}
	http.ListenAndServe(":8000", newMux)
}
