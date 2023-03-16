package data

import (
	"encoding/json"
	"github.com/emicklei/go-restful"
	"log"
	"metrorail/dbutils"
	"net/http"
)

// TrainStruct holds information about trains
type TrainStruct struct {
	ID              int
	DriverName      string
	OperatingStatus bool
}

// StationStruct holds information about stations
type StationStruct struct {
	ID          int
	Name        string
	OpeningTime string
	ClosingTime string
}

// ScheduleStruct links both trains and train schedules
type ScheduleStruct struct {
	ID          int
	TrainID     int
	StationID   int
	ArrivalTime string
}

// All the above structs are exac representations of the db tables

func (t *TrainStruct) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.
		Path("/v1/trains"). // this consumes and produces values can be set for each route
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/{train-id}").To(t.getTrain))
	ws.Route(ws.POST("").To(t.createTrain))
	ws.Route(ws.DELETE("/{train-id}").To(t.removeTrain))
	container.Add(ws)
}

func (s *StationStruct) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.
		Path("/v1/stations").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/{station-id}").To(s.getStation))
	ws.Route(ws.POST("").To(s.createStation))
	ws.Route(ws.DELETE("/{station-id}").To(s.removeStation))
	ws.Route(ws.PATCH("/{station-id}").To(s.updateStation))
	//ws.Route(ws.PUT("/{station-id}").To(s.replaceStation))
	container.Add(ws)
}

func (s *ScheduleStruct) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.
		Path("/v1/schedule").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/{schedule-id}").To(s.getSchedule))
	ws.Route(ws.POST("").To(s.createSchedule))
	ws.Route(ws.DELETE("/{schedule-id}").To(s.removeSchedule))
	//ws.Route(ws.PATCH("/{schedule-id}").To(s.updateSchedule))
	//ws.Route(ws.PUT("/{schedule-id}").To(s.replaceSchedule))
	container.Add(ws)
}

// UPDATE methods
func (u *StationStruct) updateStation(r *restful.Request, re *restful.Response) {
	body := r.Request.Body
	decoder := json.NewDecoder(body)
	var b StationStruct
	err := decoder.Decode(&b)
	if err != nil {
		log.Println(err)
		re.AddHeader("Ricardo", "Junior")
		_ = re.WriteErrorString(http.StatusBadRequest, "Couldn't update station")
	}

	//st, _ := dbutils.DB.Prepare(
	//	)
	//st.Exec(b.Name, b.OpeningTime, b.ClosingTime, b.ID)
	result, err := dbutils.DB.Exec("update station set DRIVER_NAME=?, OPENING_TIME=?, CLOSING_TIME=? where ID=?", b.Name, b.OpeningTime, b.ClosingTime, b.ID)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := result.RowsAffected()
	if rows != 1 {
		log.Fatalf("Expeted to affect 1 row, affected #%d", rows)
	}
	newID, _ := result.LastInsertId()
	b.ID = int(newID)
	_ = re.WriteHeaderAndEntity(http.StatusOK, b)
}

// GET methods

func (s *ScheduleStruct) getSchedule(r *restful.Request, re *restful.Response) {
	id := r.PathParameter("schedule-id")
	err := dbutils.DB.QueryRow("select ID, TRAIN_ID, STATION_ID, ARRIVAL_TIME FROM schedule where id=?", id).
		Scan(&s.ID, &s.TrainID, &s.StationID, &s.ArrivalTime)
	if err != nil {
		log.Println(err)
		re.AddHeader("Content-Type", "text/plain")
		_ = re.WriteErrorString(http.StatusNotFound, "Schedule could not be found")
	} else {
		err := re.WriteEntity(s)
		if err != nil {
			log.Println("Failed to write entity, schedule")
		}
	}
}

func (s *StationStruct) getStation(r *restful.Request, re *restful.Response) {
	id := r.PathParameter("station-id")
	err := dbutils.DB.QueryRow("select ID, DRIVER_NAME, OPENING_TIME, CLOSING_TIME FROM station where id=?", id).
		Scan(&s.ID, &s.Name, &s.OpeningTime, &s.ClosingTime)
	if err != nil {
		log.Println(err)
		re.AddHeader("Content-Type", "text/plain")
		_ = re.WriteErrorString(http.StatusNotFound, "Station could not be found")
	} else {
		err := re.WriteEntity(s)
		if err != nil {
			log.Println("Failed to write entity, station")
		}
	}
}

func (t *TrainStruct) getTrain(r *restful.Request, re *restful.Response) {
	id := r.PathParameter("train-id")
	err := dbutils.DB.QueryRow("select ID, DRIVER_NAME, OPERATING_STATUS FROM train where id=?", id).
		Scan(&t.ID, &t.DriverName, &t.OperatingStatus)
	if err != nil {
		log.Println(err)
		re.AddHeader("Content-Type", "text/plain")
		_ = re.WriteErrorString(http.StatusNotFound, "Train could not be found.")
	} else {
		err := re.WriteEntity(t)
		if err != nil {
			log.Fatal("Failed to write entity, train")
		}
	}
}

// POST methods
func (t *TrainStruct) createTrain(r *restful.Request, re *restful.Response) {
	log.Println(r.Request.Body)
	decoder := json.NewDecoder(r.Request.Body)
	var b TrainStruct
	err := decoder.Decode(&b)
	log.Println(b.DriverName, b.OperatingStatus)
	if err != nil {
		log.Println(err)
		re.AddHeader("Content-Type", "text/plain")
		err := re.WriteErrorString(http.StatusBadRequest, "Train could not be created.")
		if err != nil {
			log.Fatal("Failed to write error Status bad request:", err.Error())
		}
	}
	statement, _ := dbutils.DB.Prepare("insert into train (DRIVER_NAME, OPERATING_STATUS) Values (?, ?)")
	result, err := statement.Exec(b.DriverName, b.OperatingStatus)
	if err == nil {
		newID, _ := result.LastInsertId()
		b.ID = int(newID)
		err := re.WriteHeaderAndEntity(http.StatusCreated, b)
		if err != nil {
			log.Fatal("Failed to write Header: ", err.Error())
		}
	} else {
		re.AddHeader("Content-type", "text/plain")
		err := re.WriteErrorString(http.StatusInternalServerError, err.Error())
		if err != nil {
			log.Fatal("Failed to write error:", err.Error())
		}
	}
}

func (t *StationStruct) createStation(r *restful.Request, re *restful.Response) {
	log.Println(r.Request.Body)
	decoder := json.NewDecoder(r.Request.Body)
	var b StationStruct
	err := decoder.Decode(&b) // decode the body and update the value of b
	if err != nil {
		log.Println(err)
		re.AddHeader("Content-Type", "text/plain")
		err := re.WriteErrorString(http.StatusBadRequest, "Station could not be created.")
		if err != nil {
			log.Fatal("Failed to write error Status bad request:", err.Error())
		}
	}
	log.Println(b.Name, b.OpeningTime, b.ClosingTime)
	statement, _ := dbutils.DB.Prepare("insert into station (DRIVER_NAME, OPENING_TIME, CLOSING_TIME) Values (?, ?, ?)")
	result, err := statement.Exec(b.Name, b.OpeningTime, b.ClosingTime)
	if err == nil {
		newID, _ := result.LastInsertId()
		b.ID = int(newID)
		err := re.WriteHeaderAndEntity(http.StatusCreated, b)
		if err != nil {
			log.Fatal("Failed to write Header: ", err.Error())
		}
	} else {
		re.AddHeader("Content-type", "text/plain")
		err := re.WriteErrorString(http.StatusInternalServerError, err.Error())
		if err != nil {
			log.Fatal("Failed to write error:", err.Error())
		}
	}
}

func (t *ScheduleStruct) createSchedule(r *restful.Request, re *restful.Response) {
	log.Println(r.Request.Body)
	decoder := json.NewDecoder(r.Request.Body)
	var b ScheduleStruct
	err := decoder.Decode(&b)
	if err != nil {
		log.Println(err)
		re.AddHeader("Content-Type", "text/plain")
		err := re.WriteErrorString(http.StatusBadRequest, "Train could not be created.")
		if err != nil {
			log.Fatal("Failed to write error Status bad request:", err.Error())
		}
	}
	log.Println(b.TrainID, b.StationID, b.ArrivalTime)
	statement, _ := dbutils.DB.Prepare("insert into schedule (TRAIN_ID, STATION_ID, ARRIVAL_TIME) Values (?, ?, ?)")
	result, err := statement.Exec(b.TrainID, b.StationID, b.ArrivalTime)
	if err == nil {
		newID, _ := result.LastInsertId()
		b.ID = int(newID)
		err := re.WriteHeaderAndEntity(http.StatusCreated, b)
		if err != nil {
			log.Fatal("Failed to write Header: ", err.Error())
		}
	} else {
		re.AddHeader("Content-type", "text/plain")
		err := re.WriteErrorString(http.StatusInternalServerError, err.Error())
		if err != nil {
			log.Fatal("Failed to write error:", err.Error())
		}
	}
}

// Remove methods

func (t *TrainStruct) removeTrain(r *restful.Request, re *restful.Response) {
	id := r.PathParameter("train-id")
	statement, _ := dbutils.DB.Prepare("delete from train where id=?")
	_, err := statement.Exec(id)
	if err == nil {
		re.WriteHeader(http.StatusOK)
	} else {
		re.AddHeader("Content-type", "text/plain")
		err := re.WriteErrorString(http.StatusInternalServerError, err.Error())
		if err != nil {
			log.Fatal("Failed to write error:", err.Error())
		}
	}
}

func (t *StationStruct) removeStation(r *restful.Request, re *restful.Response) {
	id := r.PathParameter("station-id")
	statement, _ := dbutils.DB.Prepare("delete from station where id=?")
	result, err := statement.Exec(id)
	if err == nil {
		rows, _ := result.RowsAffected()
		log.Printf("You afected this many rows: %d", rows)
		re.WriteHeader(http.StatusOK)
	} else {
		re.AddHeader("Content-type", "text/plain")
		err := re.WriteErrorString(http.StatusInternalServerError, err.Error())
		if err != nil {
			log.Fatal("Failed to write error:", err.Error())
		}
	}
}

func (t *ScheduleStruct) removeSchedule(r *restful.Request, re *restful.Response) {
	id := r.PathParameter("schedule-id")
	statement, _ := dbutils.DB.Prepare("delete from schedule where id=?")
	result, err := statement.Exec(id)
	if err == nil {
		rows, _ := result.RowsAffected()
		log.Printf("You afected this many rows: %d", rows)
		re.WriteHeader(http.StatusOK)
	} else {
		re.AddHeader("Content-type", "text/plain")
		err := re.WriteErrorString(http.StatusInternalServerError, err.Error())
		if err != nil {
			log.Fatal("Failed to write error:", err.Error())
		}
	}
}
