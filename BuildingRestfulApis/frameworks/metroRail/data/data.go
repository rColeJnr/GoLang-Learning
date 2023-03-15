package data

import (
	"encoding/json"
	"log"
	"metrorail/dbutils"
	"net/http"
	"time"

	"github.com/emicklei/go-restful"
)

type TrainStruct struct {
	ID              int
	DriverName      string
	OperatingStatus bool
}

// StationResource holds information about locations
type StationStruct struct {
	ID          int
	Name        string
	OpeningTime time.Time
	ClosingTime time.Time
}

// ScheduleResource links both trains and stations
type ScheduleStruct struct {
	ID          int
	TrainID     int
	StationID   int
	ArrivalTime time.Time
}

// All the above structs are exac representations of the db tables

func (t *TrainStruct) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.
		Path("/v1/trains").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/{train-id}").To(t.getTrain))
	ws.Route(ws.POST("").To(t.createTrain))
	ws.Route(ws.DELETE("/{train-id}").To(t.removeTrain))
	container.Add(ws)
}

// GET
func (t TrainStruct) getTrain(r *restful.Request, re *restful.Response) {
	id := r.PathParameter("train-id")
	err := dbutils.DB.QueryRow("select ID, DRIVER_NAME, OPERATING_STATUS FROM train where id=?", id).Scan(&t.ID, &t.DriverName, &t.OperatingStatus)
	if err != nil {
		log.Println(err)
		re.AddHeader("Content-Type", "text/plain")
		re.WriteErrorString(http.StatusNotFound, "Train could not be found.")
	} else {
		re.WriteEntity(t)
	}
}

// POST
func (t TrainStruct) createTrain(r *restful.Request, re *restful.Response) {
	log.Println(r.Request.Body)
	decoder := json.NewDecoder(r.Request.Body)
	var b TrainStruct
	err := decoder.Decode(&b)
	log.Println(b.DriverName, b.OperatingStatus)
	if err != nil {
		log.Println(err)
		re.AddHeader("Content-Type", "text/plain")
		re.WriteErrorString(http.StatusBadRequest, "Train could not be created.")
	}
	statement, _ := dbutils.DB.Prepare("insert into train (DRIVER_NAME, OPERATING_STATUS) Values (?, ?)")
	result, err := statement.Exec(b.DriverName, b.OperatingStatus)
	if err == nil {
		newID, _ := result.LastInsertId()
		b.ID = int(newID)
		re.WriteHeaderAndEntity(http.StatusCreated, b)
	} else {
		re.AddHeader("Content-type", "text/plain")
		re.WriteErrorString(http.StatusInternalServerError, err.Error())
	}
}

// Delete train
func (t TrainStruct) removeTrain(r *restful.Request, re *restful.Response) {
	id := r.PathParameter("train-id")
	statement, _ := dbutils.DB.Prepare("delete from train where id=?")
	_, err := statement.Exec(id)
	if err == nil {
		re.WriteHeader(http.StatusOK)
	} else {
		re.AddHeader("Content-type", "text/plain")
		re.WriteErrorString(http.StatusInternalServerError, err.Error())
	}
}
