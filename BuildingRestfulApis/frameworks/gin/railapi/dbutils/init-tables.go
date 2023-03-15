package dbutils

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func Initialize(dbDriver *sql.DB) {

	statement, err := dbDriver.Prepare(train)
	if err != nil {
		log.Println(err) // driverError
	}

	// Create tables
	_, statementErr := statement.Exec()
	if statementErr != nil {
		log.Println("Table already exists!")
	}

	statement, _ = dbDriver.Prepare(station)
	statement.Exec()

	statement, _ = dbDriver.Prepare(schedule)
	statement.Exec()

	log.Println("All tables created/initialized successfully!")

}
