package services

import (
	"fmt"
	"log"
	"mercafacil-challenge/api/models"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Server struct {
	DBPostgreSQL *gorm.DB
	DBMySQL      *gorm.DB
	Router       *mux.Router
}

func (server *Server) InitializeDB(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error
	if Dbdriver == "mysql" {
		log.Println("Initializing DB MySQL")
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
		server.DBMySQL, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			log.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("Error:", err)
		} else {
			log.Printf("Connected to the %s database", Dbdriver)
			server.DBMySQL.Debug().AutoMigrate(&models.User{})
		}
	}
	if Dbdriver == "postgres" {
		log.Println("Initializing DB PostgreSQL")
		DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
		server.DBPostgreSQL, err = gorm.Open(Dbdriver, DBURL)
		if err != nil {
			log.Printf("Cannot connect to %s database", Dbdriver)
			log.Fatal("Error:", err)
		} else {
			log.Printf("Connected to the %s database", Dbdriver)
			server.DBPostgreSQL.Debug().AutoMigrate(&models.User{})
		}
	}
}

func (server *Server) InitializeRoutes() {
	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
