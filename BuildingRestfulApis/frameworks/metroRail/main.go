package main

import (
	"database/sql"
	"log"
	"metrorail/data"
	"metrorail/dbutils"
	"net/http"

	"github.com/emicklei/go-restful"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Connect to Database
	var err error
	dbutils.DB, err = sql.Open("sqlite3", ".railapi.db")
	if err != nil {
		log.Println("Driver creation failed!")
	}

	// Create tables
	dbutils.Initialize(dbutils.DB)
	wsContainer := restful.NewContainer()
	wsContainer.Router(restful.CurlyRouter{})
	t := data.TrainStruct{}
	t.Register(wsContainer)
	log.Printf("Start listening on :1205")
	server := &http.Server{
		Addr:    ":1205",
		Handler: wsContainer,
	}
	log.Fatal(server.ListenAndServe())
}
