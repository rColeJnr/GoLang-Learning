package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func FileServer() {
	router := httprouter.New()
	// Mapping the methods is possible with HttpRouter
	router.ServeFiles("D:\\Bandicam\\*filepath", http.Dir("C:\\Users\\ricar\\Documents"))
	log.Fatal(http.ListenAndServe(":8000", router))
}
