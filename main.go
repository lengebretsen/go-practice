package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/lengebretsen/go-practice/conf"
	"github.com/lengebretsen/go-practice/controllers"
	"github.com/lengebretsen/go-practice/db"
	"github.com/lengebretsen/go-practice/models"

	_ "github.com/lengebretsen/go-practice/docs"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
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
	conf.LoadConfig() //Load viper config

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
	router.Run(fmt.Sprintf("%s:%s", viper.Get("server.host"), viper.Get("server.port")))
}
