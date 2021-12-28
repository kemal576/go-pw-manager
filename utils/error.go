package utils

import "log"

func ErrorCheck(err error) {
	if err != nil {
		log.Panic(err)
	}
}
