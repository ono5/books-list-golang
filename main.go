package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"database/sql"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	models "github.com/ono5/books-list-golang/model"
	"github.com/ono5/books-list-golang/utils"
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
	postgresURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("USER"),
		os.Getenv("PASSWORD"),
		os.Getenv("HOST"),
		os.Getenv("PORT"),
		os.Getenv("DBNAME"))

	var err error
	db, err = sql.Open("postgres", postgresURL)
	utils.LogFatal(err, "SQL Open")
	err = db.Ping()
	utils.LogFatal(err, "Pint To DB")

	router := mux.NewRouter()

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", addBook).Methods("POST")
	router.HandleFunc("/books", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")

	fmt.Println("Server is running at port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	books = []models.Book{}

	rows, err := db.Query("select * from books")
	utils.LogFatal(err, "getBooks: SELECT")

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		utils.LogFatal(err, "getBooks: Scan")
		books = append(books, book)
	}

	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	params := mux.Vars(r)

	rows := db.QueryRow("select * from books where id=$1", params["id"])
	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	utils.LogFatal(err, "getBook: Scan")
	json.NewEncoder(w).Encode(book)
}

func addBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	var bookID int

	json.NewDecoder(r.Body).Decode(&book)
	err := db.QueryRow("insert into books (title, author, year) values($1, $2, $3) RETURNING id;",
		book.Title, book.Author, book.Year,
	).Scan(&bookID)

	utils.LogFatal(err, "addBook: QueryRow")

	json.NewEncoder(w).Encode(bookID)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)

	result, err := db.Exec("update books set title=$1, author=$2, year=$3 where id=$4 RETURNING id",
		&book.Title,
		&book.Author,
		&book.Year,
		&book.ID,
	)
	utils.LogFatal(err, "updateBook: Exec")

	rowsUpdated, err := result.RowsAffected()
	utils.LogFatal(err, "RowsAffected")

	json.NewEncoder(w).Encode(rowsUpdated)
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	result, err := db.Exec("delete from books where id = $1", params["id"])
	utils.LogFatal(err, "removeBook: Exec")

	rowsDeleted, err := result.RowsAffected()
	utils.LogFatal(err, "removeBook: RowsAffected")

	json.NewEncoder(w).Encode(rowsDeleted)
}
