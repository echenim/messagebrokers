package utils

import (
	"fmt"
	"log"
)

// //ToJSON function to convert string to json
// func ToJSON(m struct) []byte {

// 	b, err := json.Marshal(m)
// 	if err != nil {
// 		log.Fatalf("%s : %s", err, "Fail to convert to json")
// 	}

// 	return b
// }

//ErrorHandler function for capturing error
func ErrorHandler(err error, errorMessage string) {
	if err != nil {
		log.Fatalf("%s : %s", err, errorMessage)
		fmt.Printf("%v : %v", errorMessage, err)
	}
}
