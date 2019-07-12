package utils

import "log"

func FailError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s : %s", msg, err)
	}
}
