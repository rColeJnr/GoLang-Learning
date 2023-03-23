package models

import "github.com/jinzhu/gorm"
import _ "github.com/lib/pq"

// GORM tables represented by structs
type User struct {
	gorm.Model // indicates that this is a gorm table
	Orders     []Order
	Data       string `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB" json:"-"`
}

type Order struct {
	gorm.Model
	User User
	Data string `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB"`
}

// GORM creates tables with plural names. this overrides the tablename method
func (User) TableName() string {
	return "user"
}

func (Order) TableName() string {
	return "order"
}

func InitDB() (*gorm.DB, error) {
	var err error
	db, err := gorm.Open("postgres", "postgres://rick:password@localhost/mydb?sslmode=disable")
	if err != nil {
		return nil, err
	} else {
		/*
			 * The below AutoMIgrate is equivalent to this
				if !db.HasTAble("table") {
					db.CreateTable(&User{})
				}
				same this for order
		*/
		db.AutoMigrate(&User{}, &Order{})
		return db, nil
	}
}

/*Remember, PostgreSQL stores its users in a table called user. If you want
to create a new user table, create it using "user" (double quotes). Even
while retrieving use double quotes. Otherwise, the DB will fetch internal
user details:
SELECT * FROM "user"; // Correct way
SELECT * FROM user; // Wrong way. It fetches database
users*/
