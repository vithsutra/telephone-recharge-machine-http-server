package infrastructure

import (
	"log"
	"net/http"
	"os"

	"github.com/Magowtham/telephone_recharge_machine_http_server/infrastructure/db"
	"github.com/Magowtham/telephone_recharge_machine_http_server/infrastructure/repository"
	"github.com/Magowtham/telephone_recharge_machine_http_server/presentation/route"
)

func Run() {
	serverAddress := os.Getenv("SERVER_ADDRESS")

	if serverAddress == "" {
		log.Fatalln("missing or empty env variable SERVER_ADDRESS")
	}

	databaseConn, err := db.Connect()

	log.Println("connected to database")

	if err != nil {
		log.Fatalf("error occurred while connecting to database, Error -> %v\n", err.Error())
	}

	postgresRepository := repository.NewPostgresRepository(databaseConn)

	//initializing process of the database

	if err := postgresRepository.Init(); err != nil {
		log.Fatalln(err)
	}

	log.Println("database initialized successfully")

	router := route.Router(postgresRepository)

	log.Printf("server is running on %s\n", serverAddress)
	http.ListenAndServe(serverAddress, router)
}
