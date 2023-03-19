package main

import (
	"log"
	"rcole/postgres/models"
)

func main() {
	db, err := models.InitDB()

	if err != nil {
		log.Println(db, "this") // printing the db connection, which will be an address.
	}
	log.Println(err)

}
