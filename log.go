package main

import (
	"io"
	"log"
	"os"
)

func initLogger() {
	logFileName := "app.log"

	logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("unable to create logger file, Error -> %v\n", err.Error())
	}

	logWriter := io.MultiWriter(os.Stdout, logFile)

	log.SetOutput(logWriter)

}
