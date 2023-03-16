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
	dbutils.DB, err = sql.Open("sqlite3", "./railapi.db")
	if err != nil {
		log.Println("Driver creation failed!")
	}

	// Create tables
	dbutils.Initialize(dbutils.DB)
	wsContainer := restful.NewContainer()
	wsContainer.Router(restful.CurlyRouter{})
	t := data.TrainStruct{}
	t.Register(wsContainer)
	s := data.StationStruct{}
	s.Register(wsContainer)
	sc := data.ScheduleStruct{}
	sc.Register(wsContainer)
	var port string = ":1106"
	log.Printf("Start listening on %s", port)
	server := &http.Server{
		Addr:    port,
		Handler: wsContainer,
	}
	log.Fatal(server.ListenAndServe())
}
