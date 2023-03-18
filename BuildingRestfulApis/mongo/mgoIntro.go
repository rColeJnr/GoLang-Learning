package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Movie holds a movie data
type Movie struct {
	Name      string   `bson:"name"`
	Year      string   `bson:"year"`
	Directors []string `bson:"directors"`
	Writers   []string `bson:"writers"`
	BoxOffice `bson:"boxOffice"`
}

// BoxOffice is nested in Movie
type BoxOffice struct {
	Budget uint64 `bson:"budget"`
	Gross  uint64 `bson:"gross"`
}

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	fmt.Println("ClientOptopm TYPE:", reflect.TypeOf(clientOptions), "\n")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("Mongo.connect() ERROR: ", err)
		os.Exit(1)
	}
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	c := client.Database("appdb").Collection("movies")

	// Create a movie
	darkNight := &Movie{
		Name:      "The Dark Knight",
		Year:      "2008",
		Directors: []string{"Christopher Nolan"},
		Writers:   []string{"Jonathan Nolan", "Christopher Nolan"},
		BoxOffice: BoxOffice{
			Budget: 185000000,
			Gross:  533316061,
		},
	}

	// Insert into MongoDB
	result, err := c.InsertOne(ctx, darkNight)
	if err != nil {
		log.Fatal(err)
	}

	// Now query the movie back
	// movie := Movie{}
	// bson.M is used for nested fields
	// err = c.Find(ctx, bson.M{"boxOffice.budget": bson.M{"$gt": 150000000}}).One(&movie)
	movie := c.FindOne(ctx, bson.M{"boxoffice.budget": bson.M{"$gt": 12555}})

	fmt.Println("InsertOne() api result type: ", &result)
	// fmt.Println("InsertOne() api result type: ", movie)
	fmt.Println("Movie:", &movie)

}
