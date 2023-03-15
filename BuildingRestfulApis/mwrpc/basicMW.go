package main

import (
	"fmt"
	"net/http"
)

func Middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Executing middleware before request phase!")
		// pass control back to the handler
		handler.ServeHTTP(w, r)
		fmt.Println("Executing middleware after response phase!")
	})
}

func MainLogic(w http.ResponseWriter, r *http.Request) {
	// Business logic goes here
	fmt.Println("Executing mainHandler...")
	w.Write([]byte("ok"))
}

// HandlerFunc returns a HTTP Handler
// func main() {
// 	mainLogcHandler := http.HandlerFunc(MainLogic) // create a handler func by passing main handlerfunc
// 	http.Handle("/", Middleware(mainLogcHandler)) // create middleware takes handler returns handler
// 	http.ListenAndServe(":1205", nil)
// }
