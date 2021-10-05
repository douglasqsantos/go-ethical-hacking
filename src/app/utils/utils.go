package utils

import "log"

// HandleError function to handle errors
func HandleError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// PanicMSG function to handle panic messages
func PanicMSG(err error, msg string) {
	if err != nil {
		panic(msg)
	}
}
