package main

import (
	"gorilla/router"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

/* muxRouter/*
func main() {
	// Create a new router
	r := mux.NewRouter()
	// attach a path with handler
	r.HandleFunc("/articles/{category}/{id:[0-9]+}", router.ArticleHandler)
	r.HandleFunc("/articles", router.QueryHandler)
	r.Queries("id", "category")
	srv := &http.Server{
		Handler:      r,
		Addr:         ":1205",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
*/

// qParams
func main() {
	// create a new router
	r := mux.NewRouter()

	// Attach a path with handler
	r.HandleFunc("/articles", router.QParamsHandler)
	r.Queries("id", "category")
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
