package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Article handler
func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	// mux.Vars returns all path params as a map
	vars := mux.Vars(r)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category is: %v\n", vars["category"])
	fmt.Fprintf(w, "ID is: %v\n", vars["id"])
}

// Query-bsed matching
func QueryHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Got parameter id:%s!\n", queryParams["id"])
	fmt.Fprintf(w, "Got parameter category:%s!\n", queryParams["category"])
}
