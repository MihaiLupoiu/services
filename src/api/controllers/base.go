package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MihaiLupoiu/services/src/api/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres driver
)

// Service global context
type Service struct {
	DB     *gorm.DB
	Router *mux.Router
}

// Initialize Server Database
func (service *Service) Initialize(DBType, DBUser, DBPassword, DBHost, DBName string, DBPort int) {
	var err error

	if DBType == "postgres" {
		DBConfig := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable password=%s", DBHost, DBPort, DBUser, DBName, DBPassword)
		service.DB, err = gorm.Open(DBType, DBConfig)
		if err != nil {
			log.Fatal("Cannot connect to postgres database. Error:", err)
		}
	} else {
		log.Fatal("Unsupported database type:", DBType)
	}

	// Migrate the schema
	service.DB.Debug().AutoMigrate(&models.User{})

	// Create Router and initialize callbacks
	service.Router = mux.NewRouter()
	service.initializeRoutes()
}

// Run http server
func (service *Service) Run(addr string) {
	log.Println("Listening to address", addr)
	log.Fatal(http.ListenAndServe(addr, service.Router))
}
