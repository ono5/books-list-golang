package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	models "github.com/ono5/books-list-golang/model"
)

func LogFatal(err error, comment string) {
	if err != nil {
		fmt.Println("#############################################")
		log.Fatal(comment)
		log.Fatal(err)
		fmt.Println("#############################################")
	}
}

func SendError(w http.ResponseWriter, status int, err models.Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(err)
}

func SendSuccess(w http.ResponseWriter, data any) {
	json.NewEncoder(w).Encode(data)
}
