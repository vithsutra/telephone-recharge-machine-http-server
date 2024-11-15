package main

import (
	"log"

	"github.com/Magowtham/telephone_recharge_machine_http_server/infrastructure"
	"github.com/joho/godotenv"
)

func main() {

	//intializing the logger file
	initLogger()

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading the env file, Error -> %v\n", err.Error())
	}

	infrastructure.Run()
}
