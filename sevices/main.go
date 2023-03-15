package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"product-api/handlers"
	"time"

	"github.com/gorilla/mux"
	"github.com/nicholasjackson/env"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":1205", "Bind address for the server")

func main() {

	env.Parse()

	l := log.New(os.Stdout, "products-api", log.LstdFlags)

	// create handler
	ph := handlers.NewProducts(l)

	// Create a new server and register handlers
	sm := mux.NewRouter() // using the gorrila package

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}", ph.UpdateProducts)
	putRouter.Use(ph.MiddlewareValidateProduct)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.AddProduct)
	postRouter.Use(ph.MiddlewareValidateProduct)

	// create a new server
	s := &http.Server{
		Addr:         *bindAddress,      // Configure the bind address
		Handler:      sm,                // set the default handler
		IdleTimeout:  120 * time.Second, // max time for connections using TCP keep-alive
		ReadTimeout:  1 * time.Second,   // max time to read request from the client
		WriteTimeout: 1 * time.Second,   // max time to write response to the client
		ErrorLog:     l,                 // set the logger for the server
	}

	// start the server
	go func() {
		l.Printf("Starting server on port %v\n", *bindAddress)

		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)      // create os.Signal chan
	signal.Notify(sigChan, os.Interrupt) // notify our channel whenever system receives an interrupt signal
	signal.Notify(sigChan, os.Kill)      // notify for kill commands `ctlr+c`

	sig := <-sigChan // receiving block for our chan
	l.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second) // finish all processes in the next 30secs (db connections, large uploads, finish them close then stdown)
	s.Shutdown(tc)                                                     // shutdown our server

}
