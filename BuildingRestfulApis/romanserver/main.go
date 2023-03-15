package main

import (
	"fmt"
	"html"
	"net/http"
	"romanserver/data"
	"strconv"
	"strings"
	"time"
)

func main() {
	// http package has methods for dealing with requests
	http.HandleFunc("/", romanHandler)

	// create a server and run it on 1205 port
	port := ":1205"
	s := &http.Server{
		Addr:           port,
		ReadTimeout:    10 * time.Second, // after 10 secs return 408 reqeust timeout
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
func romanHandler(rw http.ResponseWriter, r *http.Request) {
	urlPathElements := strings.Split(r.URL.Path, "/") //r.URL.Path is the path of the HTTP request

	// if request is GET with correct syntax
	if urlPathElements[1] == "roman_number" {
		// Atoi converts an alphanumeric string to an int
		number, _ := strconv.Atoi(strings.TrimSpace(urlPathElements[2]))
		if number == 0 || number > 10 {
			// if resource is not in the list, send not found status
			rw.WriteHeader(http.StatusNotFound)
			rw.Write([]byte("404 - Not Found"))
		} else {
			fmt.Fprintf(rw, "%q", html.EscapeString(data.Numerals[number])) // Escapes special chars to valid html chars.
		}
	} else {
		// For al other reqeust, tell taht client sent a bad reqeust
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("400 - Bad request"))
	}
}
