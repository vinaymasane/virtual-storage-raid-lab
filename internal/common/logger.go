package common

import "log"

func Info(msg string) {
	log.Printf("[INFO ] %s", msg)
}

func Warn(msg string) {
	log.Printf("[WARN ] %s", msg)
}

func Error(msg string) {
	log.Printf("[ERROR] %s", msg)
}