package utils

import (
	"fmt"
	"log"
)

//ErrorHandler function for capturing error
func ErrorHandler(err error, errorMessage string) {
	if err != nil {
		log.Fatalf("%s : %s", err, errorMessage)
		fmt.Printf("%v : %v", errorMessage, err)
	}
}
