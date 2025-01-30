package main

import (
	"log"
	"os"

	"github.com/Magowtham/telephone_recharge_machine_http_server/infrastructure"
	"github.com/joho/godotenv"
)

func init() {
	mode := os.Getenv("SERVER_MODE")

	if mode == "dev" {
		if err := godotenv.Load(); err != nil {
			log.Fatalf("error loading the env file, Error -> %v\n", err.Error())
			return
		} else {
			log.Println(".env file loaded successfully")
			return
		}
	}

	if mode != "prod" {
		log.Fatalln("please set the SERVER_MODE env variable to 'prod' or 'dev'")
	}
}

func main() {

	//intializing the logger file
	initLogger()

	infrastructure.Run()
}
