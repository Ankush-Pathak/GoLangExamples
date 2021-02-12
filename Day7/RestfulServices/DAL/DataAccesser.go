package DAL

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Book struct {
	ID     int //First letter capital so that member is also accessible publicly
	Name   string
	ISBN   string
	Author string `json:"BookAuthor"` //When encoding need to use BookAuthor instead of Author
	Pages  int
}

func Connect() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB connected...")
	return client
}

func GetBookByIsbn(isbn string) Book {
	var book Book
	session := Connect()

	collection := session.Database("webinardb").Collection("books")

	filter := bson.M{"isbn": isbn}
	error := collection.FindOne(context.TODO(), filter).Decode(&book)
	if error != nil {
		fmt.Println(error)
	}
	return book

}

func InsertBook(book Book) {
	session := Connect()

	collection := session.Database("webinardb").Collection("books")

	result, error := collection.InsertOne(context.TODO(), book)
	fmt.Println("Insertion result:", *result)
	if error != nil {
		fmt.Println("iInserting error:", error)
	}
}

func DeleteBookByIsbn(isbn string) {
	session := Connect()

	collection := session.Database("webinardb").Collection("books")

	filter := bson.M{"isbn": isbn}

	res, error := collection.DeleteOne(context.TODO(), filter)

	fmt.Println("Deletion result:", *res)
	if error != nil {
		fmt.Println("Deletion error:", error)
	}
}

func GetBooks() []Book {
	var books []Book
	//Static data
	/*books = append(books, Book{1, "Go", "123", "G", 100})
	books = append(books, Book{2, "C", "456", "C", 101})
	books = append(books, Book{3, "Python", "789", "P", 102})
	books = append(books, Book{4, "Java", "101112", "J", 103})*/

	session := Connect()

	collection := session.Database("webinardb").Collection("books")
	cursor, error := collection.Find(context.TODO(), bson.M{})
	defer cursor.Close(context.TODO())
	if error != nil {
		fmt.Println("Error when fetching data:", error)
	}

	for cursor.Next(context.TODO()) {
		//Convert MongoDB record to book object
		var book Book
		error := cursor.Decode(&book)
		if error != nil {
			fmt.Println("Error when decoding Mongo data:", error)
		}
		books = append(books, book)
	}

	return books
}
