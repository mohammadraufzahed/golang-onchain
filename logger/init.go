package logger

import (
	"log"
	"os"
)

var Logger *log.Logger = log.Default()

func Init() {
	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile("./logs/logs", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	Logger.SetOutput(file)
}
