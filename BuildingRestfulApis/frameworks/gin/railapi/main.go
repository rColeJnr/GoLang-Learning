package main

import (
	"database/sql"
	"gin/railapi/data"
	"gin/railapi/dbutils"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	var err error
	dbutils.DB, err = sql.Open("sqlite3", "./railapi.db")
	if err != nil {
		log.Println("Driver creation failed")
	}
	dbutils.Initialize(dbutils.DB)
	r := gin.Default()
	// add routes
	r.GET("/stations/:stationID", data.GetStation)
	r.POST("stations", data.CreateStation)
	r.DELETE("stations/:stationID", data.RemoveStation)
	log.Fatal(r.Run(":2908"))
}
