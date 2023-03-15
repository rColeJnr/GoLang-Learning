package data

import (
	"gin/railapi/dbutils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

type TrainStruct struct {
	ID              int    `json:"id"`
	DriverName      string `json:"driverName"`
	OperatingStatus bool   `json:"operatingStatus"`
}

// StationResource holds information about locations
type StationStruct struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	OpeningTime time.Time `json:"opening_time"`
	ClosingTime time.Time `json:"closing_time"`
}

// ScheduleResource links both trains and stations
type ScheduleStruct struct {
	ID          int       `json:"id"`
	TrainID     int       `json:"trainID"`
	StationID   int       `json:"stationID"`
	ArrivalTime time.Time `json:"arrivalTime"`
}

// All the above structs are exac representations of the db tables

func GetStation(c *gin.Context) {
	var station StationStruct
	id := c.Param("stationID")
	err := dbutils.DB.QueryRow("select ID, NAME, CAST(OPENING_TIME as CHAR), CAST(CLOSING_TIME as CHAR) from station where id=?", id).Scan(&station.ID, &station.Name, &station.OpeningTime, &station.ClosingTime)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(200, gin.H{"result": station})
	}
}

func CreateStation(c *gin.Context) {
	var station StationStruct
	// parse the body into our resource
	if err := c.BindJSON(&station); err == nil {
		// format time to Go time format
		statement, _ := dbutils.DB.Prepare("insert into station (NAME, OPENING_TIME, CLOSING_TIME) values (?, ?, ?)")
		result, _ := statement.Exec(station.Name, station.OpeningTime, station.ClosingTime)
		if err == nil {
			newId, _ := result.LastInsertId()
			station.ID = int(newId)
			c.JSON(http.StatusOK, gin.H{"result": station})

		} else {
			c.String(http.StatusInternalServerError, err.Error())
		}
	} else {
		c.String(http.StatusInternalServerError, err.Error())
	}
}

// RemoveStation handles the removing of resources
func RemoveStation(c *gin.Context) {
	id := c.Param("stationID")
	statement, _ := dbutils.DB.Prepare("delete from station where id=?")
	_, err := statement.Exec(id)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": err.Error()})
	} else {
		c.String(http.StatusOK, "")
	}
}