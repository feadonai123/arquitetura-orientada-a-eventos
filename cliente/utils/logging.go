package utils

import (
	"log"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func LogInfo(msg string, who string) {
	log.Printf("[INFO] [%s] %s", who, msg)
}

func LogError(err error, msg string, who string) {
	log.Printf("[ERROR] [%s] %s: %s", who, msg, err)
}