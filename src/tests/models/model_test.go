package models

import (
	"flag"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/MihaiLupoiu/services/src/api/controllers"
	"github.com/MihaiLupoiu/services/src/api/models"
	"github.com/jinzhu/gorm"
)

var service = controllers.Service{}
var userInstance = models.User{}

func TestMain(m *testing.M) {
	configPath := flag.String("config", "../../../configs/local.json", "JSON config file")
	flag.Parse()

	config := models.Config{}
	config.Load(*configPath)

	Database(config.Database.Type, config.Database.User, config.Database.Password, config.Database.Host, config.Database.DatabaseName, config.Service.APISecret, config.Database.Port)

	os.Exit(m.Run())
}

func Database(DBType, DBUser, DBPassword, DBHost, DBName, APISecret string, DBPort int) {
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
}

func refreshUserTable() error {
	err := service.DB.DropTableIfExists(&models.User{}).Error
	if err != nil {
		return err
	}
	err = service.DB.AutoMigrate(&models.User{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed table")
	return nil
}

func seedOneUser() (models.User, error) {

	refreshUserTable()

	user := models.User{
		FirstName:   "Juan",
		LastName:    "Antonio",
		Email:       "ja@gmail.com",
		Password:    "password",
		PhoneNumber: 686554322,
		Country:     "Spain",
		PostalCode:  12005,
	}

	err := service.DB.Model(&models.User{}).Create(&user).Error
	if err != nil {
		log.Fatalf("cannot seed users table: %v", err)
	}
	return user, nil
}

func seedUsers() error {

	users := []models.User{
		models.User{
			FirstName: "Steven",
			LastName:  "Victor",
			Email:     "steven@gmail.com",
			Password:  "password",
		},
		models.User{
			FirstName: "Kenny",
			LastName:  "Morris",
			Email:     "kenny@gmail.com",
			Password:  "password",
		},
	}

	for i, _ := range users {
		err := service.DB.Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			return err
		}
	}
	return nil
}
