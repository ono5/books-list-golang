package utils

import (
	"fmt"
	"log"
)

func LogFatal(err error, comment string) {
	if err != nil {
		fmt.Println("#############################################")
		log.Fatal(comment)
		log.Fatal(err)
		fmt.Println("#############################################")
	}
}
