package handlers

import (
	"fmt"
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

// creates a new goodbye handller with the given logger
func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l: l}
}

// writing a method for the ServeHttp struct
func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	g.l.Println("Handle Goodby request")

	// write the response
	fmt.Fprintf(rw, "Goodbey")
}
