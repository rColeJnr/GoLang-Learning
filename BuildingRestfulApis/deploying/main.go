// This file is a basic Go server to illustrate the proxy server's functioning.
// then, we add a configuration to Nginx to proxy port 8000
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

type Book struct {
	ID            int
	ISBN          string
	Author        string
	PublishedYear string
}

func main() {
	// file open for reading, writing and appending
	f, err := os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("error openining file: %v", err)

	}
	defer f.Close()
	log.SetOutput(f)
	// Function handler for handling requests
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%q", r.UserAgent())
		// Fill the book details
		book := Book{
			ID:            123,
			ISBN:          "0-201-8484-2",
			Author:        "Ricardo Something",
			PublishedYear: "2022",
		}
		// convert struct to json using marshal
		jsonData, _ := json.Marshal(book)
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	})
	s := &http.Server{
		Addr:           ":1205",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal("listrning on port", s.ListenAndServe())
}
