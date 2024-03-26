package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/piyushkumar/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017"
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

func init() {
	fmt.Println("welcom to mongo db api database")

	//! client options
	clientOptions := options.Client().ApplyURI(connectionString)

	//! connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB Connected")

	collection = client.Database(dbName).Collection(colName)

}

//! helper function to add data to the db

func insertOneMovie(movie model.Netflix) {

	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movie one  added", inserted.InsertedID)
}

//! updated 1 record

func updateOneMovie(movieId string) {

	id, _ := primitive.ObjectIDFromHex(movieId)

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	updated, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movie one updated", updated.ModifiedCount)

}

//! delete 1 record

func deleteOneMovie(movieId string) {

	id, _ := primitive.ObjectIDFromHex(movieId) //! convert string to object id

	filter := bson.M{"_id": id} //! filter to delete

	deleted, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movie one deleted", deleted.DeletedCount)

}

//! delete all records

func deleteAllMovies() {

	filter := bson.D{{}} //! filter to delete all

	deleted, err := collection.DeleteMany(context.Background(), filter, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movie all deleted", deleted.DeletedCount)

}

//! get all records

func getAllMovies() []primitive.M {

	//var movies []model.Netflix

	cursor, err := collection.Find(context.Background(), bson.D{{}}) //! not passing any filter to get all records

	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for cursor.Next(context.Background()) {
		var movie bson.M
		err := cursor.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie)

	}

	defer cursor.Close(context.Background())

	return movies
}

//! actual controllers

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	movies := getAllMovies()

	json.NewEncoder(w).Encode(movies)

}

func CreateMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "POST")

	var movie model.Netflix

	_ = json.NewDecoder(r.Body).Decode(&movie)

	insertOneMovie(movie)

	json.NewEncoder(w).Encode(movie)

}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	updateOneMovie(params["id"])

	json.NewEncoder(w).Encode("Movie updated successfully")

}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	deleteOneMovie(params["id"])

	json.NewEncoder(w).Encode("Movie deleted successfully")

}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	deleteAllMovies()

	json.NewEncoder(w).Encode("All movies deleted successfully")

}
