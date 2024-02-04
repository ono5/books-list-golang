package utils

import (
	"encoding/json"
	"net/http"

	models "github.com/ono5/books-list-golang/models"
)

func SendError(w http.ResponseWriter, status int, err models.Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(err)
}

func SendSuccess(w http.ResponseWriter, data any) {
	json.NewEncoder(w).Encode(data)
}
