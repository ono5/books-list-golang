package main

import (
	"fmt"
	"log"
	"net/http"

	"database/sql"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	controllers "github.com/ono5/books-list-golang/controllers"
	"github.com/ono5/books-list-golang/driver"
	models "github.com/ono5/books-list-golang/models"
	"github.com/subosito/gotenv"
)

var (
	books []models.Book
	db    *sql.DB
)

func init() {
	gotenv.Load()
}

func main() {
	db, err := driver.ConnectDB()

	if err != nil {
		log.Fatal(err)
	}
	controller := controllers.Controller{}

	router := mux.NewRouter()

	router.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", controller.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books", controller.AddBook(db)).Methods("POST")
	router.HandleFunc("/books", controller.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}", controller.RemoveBook(db)).Methods("DELETE")

	fmt.Println("Server is running at port 8000")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"*"}))(router)))
}
