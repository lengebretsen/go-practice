package main

import (
	"engebretsen/simple_web_svc/models"
	"engebretsen/simple_web_svc/pkg/common/db"
	"engebretsen/simple_web_svc/pkg/controllers"
	"log"
	"os"
	"strconv"

	_ "engebretsen/simple_web_svc/docs"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Go + Gin API Practice
// @version 1.0
// @description A small go/gin web app providing a simple REST API for managing users and addresses

// @host localhost:8080
// @BasePath /
// @query.collection.format multi

func main() {
	database, err := db.Init()
	if err != nil {
		log.Fatal("Failed to initialize database")
	}
	defer database.Close()

	//Write PID file for make down target
	pid := os.Getpid()
	err = os.WriteFile("./GINSVR.pid", []byte(strconv.Itoa(pid)), 0644)
	if err != nil {
		log.Fatal("Failed to write PID file.")
	}

	router := gin.Default()

	// docs route
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	controllers.RegisterRoutes(router, models.UserModel{DB: database}, models.AddressModel{DB: database})
	router.Run("localhost: 8080")
}
