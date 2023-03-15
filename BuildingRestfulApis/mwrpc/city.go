package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type city struct {
	Name string
	Area uint64
}

func mainLogic(w http.ResponseWriter, r *http.Request) {
	// check if the method is POST
	if r.Method == http.MethodPost {
		var tempCity city
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&tempCity)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()
		// Your resource creation logic goes here
		fmt.Printf("Got &s city with area of &d kilometers!\n", tempCity.Name, tempCity.Area)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("201 - Created"))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 - Method not allowed"))
	}
}

// middleware to check contetn type as json
func filerContentType(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Currently in the check content type middleware")
		// Filtering requests by MIME type
		if r.Header.Get("Content-type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte("415 - Unsupported Media Type. Please send JSON\n"))
			return
		}
		handler.ServeHTTP(w, r)
	})
}

// add server timestamp for response cookie
func setServerTimeCookie(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
		// sEtting cookie to each and every response
		cookie := http.Cookie{Name: "Server-time(UTC)", Value: strconv.FormatInt(time.Now().Unix(), 10)}
		http.SetCookie(w, &cookie)
		log.Println("Currentlyin the set server tiem middleware")
	})
}
func main() {
	mainLogicHandler := http.HandlerFunc(mainLogic)
	http.Handle("/city", filerContentType(setServerTimeCookie(mainLogicHandler)))
	http.ListenAndServe(":1205", nil)
}
