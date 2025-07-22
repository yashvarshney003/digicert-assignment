package main

import (
	"public_library/config"
	"public_library/database"
	"public_library/models"
	"public_library/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	database.Connect()
	database.DB.AutoMigrate(
		&models.Book{},
	)

	r := gin.Default()
	routes.SetupRoutes(r)

	r.Run(":" + config.AppConfig.AppPort)
}
