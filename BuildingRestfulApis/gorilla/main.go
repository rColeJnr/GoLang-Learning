package main

import (
	"gorilla/router"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new router
	r := mux.NewRouter()
	// attach a path with handler
	r.HandleFunc("/articles/{category}/{id:[0-9]+}", router.ArticleHandler)
	srv := &http.Server{
		Handler:      r,
		Addr:         ":1205",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
