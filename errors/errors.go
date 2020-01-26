package errors

import (
	"log"
)

// HandleError handles any error in this project
func HandleError(msg string, err error) {
	if err != nil {
		log.Fatalln(msg + ":" + err.Error())
	}
}
