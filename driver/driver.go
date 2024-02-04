package driver

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/ono5/books-list-golang/utils"
)

func ConnectDB() *sql.DB {
	postgresURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("USER"),
		os.Getenv("PASSWORD"),
		os.Getenv("HOST"),
		os.Getenv("PORT"),
		os.Getenv("DBNAME"))

	var err error
	db, err := sql.Open("postgres", postgresURL)
	utils.LogFatal(err, "SQL Open")
	err = db.Ping()
	utils.LogFatal(err, "Pint To DB")

	return db
}
