package api

import (
	"log"
	"mercafacil-challenge/api/controllers"
	"mercafacil-challenge/api/seeder"
	"os"

	"github.com/joho/godotenv"
)

var server = controllers.Server{}

func Run() {
	var err error
	err = godotenv.Load()
	if err == nil {
		log.Println("Getting the env values")
	} else {
		log.Fatalf("Error getting env, not comming through %v", err)
	}

	server.InitializeDB(os.Getenv("DB_DRIVER_MYSQL"), os.Getenv("DB_USER_MYSQL"), os.Getenv("DB_PASSWORD_MYSQL"), os.Getenv("DB_PORT_MYSQL"), os.Getenv("DB_HOST_MYSQL"), os.Getenv("DB_NAME_MYSQL"))
	server.InitializeDB(os.Getenv("DB_DRIVER_POSTGRESQL"), os.Getenv("DB_USER_POSTGRESQL"), os.Getenv("DB_PASSWORD_POSTGRESQL"), os.Getenv("DB_PORT_POSTGRESQL"), os.Getenv("DB_HOST_POSTGRESQL"), os.Getenv("DB_NAME_POSTGRESQL"))

	seeder.Load(server.DBMySQL, os.Getenv("API_SECRET_MACAPA"))
	seeder.Load(server.DBPostgreSQL, os.Getenv("API_SECRET_VAREJAO"))

	server.InitializeRoutes()
	server.Run(":8080")
}
