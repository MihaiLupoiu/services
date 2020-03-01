package api

import (
	"flag"
	"fmt"

	"github.com/MihaiLupoiu/services/src/api/controllers"
	"github.com/MihaiLupoiu/services/src/api/models"
)

var userService = controllers.Service{}

// Run main service funtion
func Run() {
	configPath := flag.String("config", "../configs/local.json", "JSON config file")
	flag.Parse()

	config := models.Config{}
	config.Load(*configPath)

	userService.Initialize(config.Database.Type, config.Database.User, config.Database.Password, config.Database.Host, config.Database.DatabaseName, config.Database.Port)

	userService.Run(fmt.Sprintf(":%d", config.Service.Port))
}
