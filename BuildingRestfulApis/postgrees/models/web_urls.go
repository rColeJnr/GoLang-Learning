package models

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func InitDB() (*sql.DB, error) {
	var err error
	// Database type - postgress | username | password | server ip address | database name | Security mode
	db, err := sql.Open("postgres", "postgres://rick:password@localhost/mydb?sslmode=disable")
	if err != nil {
		log.Println("error here", err)
		return nil, err
	} else {
		// Create model for our URL service
		stmt, err := db.Prepare("CREATE TABLE WEB_URL(ID SERIAL PRIMARY KEY, URL TEXT NOT NULL );")
		if err != nil {
			log.Println(err, "prepare")
			return nil, err
		}
		res, err := stmt.Exec()
		log.Println(res)
		if err != nil {
			log.Println(err, "exec err")
			return nil, err
		}
		return db, nil
	}
}
