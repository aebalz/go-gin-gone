package main

import (
	"fmt"

	"github.com/aebalz/go-gin-gone/configs"
	"github.com/aebalz/go-gin-gone/pkg/database"
	"github.com/aebalz/go-gin-gone/pkg/server"
	"github.com/aebalz/go-gin-gone/routes"
	"github.com/aebalz/go-gin-gone/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	configs.LoadConfig()
	isDebug := utils.StringToBool(configs.AppConfig.AppDebug)

	// Set up Database
	db := database.NewMySQLDatabase()
	if err := db.Connect(configs.AppConfig, isDebug); err != nil {
		panic("Failed to connect to MySQL database: " + err.Error())
	}
	defer db.Close()

	// Set up the Gin router
	r := gin.Default()
	routes.InitialRoutesApp(r, db.GetDB())

	// // Run the server with graceful shutdown
	appPort := fmt.Sprintf(":%s", configs.AppConfig.AppPort)
	server.RunServer(r, appPort)
}
