package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ono5/books-list-golang/models"
	bookRepository "github.com/ono5/books-list-golang/repository/book"
	"github.com/ono5/books-list-golang/utils"
)

type Controller struct{}

var books []models.Book

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		books = []models.Book{}
		var error models.Error
		bookRepo := bookRepository.BookRepository{}
		books, err := bookRepo.GetBooks(db, book, books)

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, books)
	}
}

func (c Controller) GetBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var error models.Error
		bookRepo := bookRepository.BookRepository{}

		params := mux.Vars(r)
		id := params["id"]

		book, err := bookRepo.GetBook(db, book, id)

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, book)
	}
}

func (c Controller) AddBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var error models.Error

		bookRepo := bookRepository.BookRepository{}
		json.NewDecoder(r.Body).Decode(&book)

		bookID, err := bookRepo.AddBook(db, book)

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, bookID)
	}
}

func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var error models.Error
		bookRepo := bookRepository.BookRepository{}

		json.NewDecoder(r.Body).Decode(&book)
		rowsUpdated, err := bookRepo.UpdateBook(db, book)

		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, rowsUpdated)
	}
}

func (c Controller) RemoveBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		bookID := params["id"]
		var error models.Error
		bookRepo := bookRepository.BookRepository{}

		rowsDeleted, err := bookRepo.RemoveBook(db, bookID)
		if err != nil {
			error.Message = "Server error"
			utils.SendError(w, http.StatusInternalServerError, error)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		utils.SendSuccess(w, rowsDeleted)
	}
}
