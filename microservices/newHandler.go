package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type helloWorldRequest struct {
	Name string `json: "name"`
}

type helloWorldResponse struct {
	Message string `json: "message"`
}

func main() {
	port := 1206

	handler := newValidationHandler(newHelloWorldHandler())

	http.Handle("/helloworld", handler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

type validationHandler struct {
	next http.Handler
}

func newValidationHandler(next http.Handler) http.Handler {
	return validationHandler{next: next}
}

func (h validationHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	var request helloWorldRequest
	const jsonStream = `
	[
		{"Name": "Ed"}
	]
	`
	decoder := json.NewDecoder(strings.NewReader(jsonStream))

	err := decoder.Decode(&request)
	if err != nil {
		http.Error(rw, "Bad request", http.StatusBadRequest)
		return // stop handler chain
	}
	// call the next handler int he chain
	h.next.ServeHTTP(rw, r)
}

type helloWordHandler struct{}

func newHelloWorldHandler() http.Handler {
	return helloWordHandler{}
}

func (h helloWordHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{Message: "Hello"}

	encoder := json.NewEncoder(rw)
	encoder.Encode(response)
}
