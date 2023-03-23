package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"rcole/jsonstore/models"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"time"
)

// DB stores the database session info. needs to be init once
type DBClient struct {
	db *gorm.DB
}

// UserResponse is the response to be send back for user
type UserResponse struct {
	User models.User `json:"user"`
	Data interface{} `json:"data"`
}

// GetUserByFirstName fetches the original URL for the given encoded string
func (driver *DBClient) GetUserByFirstName(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	name := r.FormValue("first_name")
	// Handle response details
	var query = "select * from \"user\" where data->>'first_name'=?"
	driver.db.Raw(query, name).Scan(&users)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	respJson, _ := json.Marshal(users)
	w.Write(respJson)
}

// GetUser fetches the original URL for the given string
func (driver *DBClient) GetUser(w http.ResponseWriter, r *http.Request) {
	var user = models.User{}
	vars := mux.Vars(r)
	// Handle response details
	driver.db.First(&user, vars["id"]) // Fetch the first record from the database with given ID. feeds the data to the user struct
	var userData interface{}
	// Unmarshal JSON string to interface
	json.Unmarshal([]byte(user.Data), &userData)
	var response = UserResponse{User: user, Data: userData}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	respJson, _ := json.Marshal(response)
	w.Write(respJson)
}

// postUser adds URL TO db AND GIVES BACK SHORT STRING
func (driver *DBClient) PostUser(w http.ResponseWriter, r *http.Request) {
	var user = models.User{}
	postBody, _ := ioutil.ReadAll(r.Body)
	user.Data = string(postBody)
	driver.db.Save(&user)
	responseMap := map[string]interface{}{"id": user.ID}
	var err string = ""
	if err != "" {
		w.Write([]byte("yes"))
	} else {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(responseMap)
		w.Write(response)
	}
}

func main() {
	db, err := models.InitDB()
	if err != nil {
		panic(err)
	}
	dbclient := &DBClient{db: db}
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Create new router
	r := mux.NewRouter()
	r.HandleFunc("/v1/user/{id:[a-zA-Z0-9]*}", dbclient.GetUser).Methods(http.MethodGet)
	r.HandleFunc("/v1/user", dbclient.PostUser).Methods(http.MethodPost)
	r.HandleFunc("/v1/user", dbclient.GetUserByFirstName).Methods(http.MethodGet)
	srv := &http.Server{
		Handler:      r,
		Addr:         ":1205",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
