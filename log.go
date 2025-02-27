package main

import (
	"io"
	"log"
	"os"
)

func initLogger() {
	logFilePath := "/var/logs/app.log" // Ensure it matches your Docker volume

	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("unable to create logger file, Error -> %v\n", err.Error())
	}

	logWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(logWriter)

	log.Println("Logger initialized successfully!")
}
