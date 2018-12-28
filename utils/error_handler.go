package utils

import "log"

// FailOnError logs the error
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", err, msg)
	}
}
