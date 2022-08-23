package main

import (
	"engebretsen/simple_web_svc/models"
	"engebretsen/simple_web_svc/pkg/common/db"
	"engebretsen/simple_web_svc/pkg/controllers"
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

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

	controllers.RegisterRoutes(router, models.UserModel{DB: database})
	router.Run("localhost: 8080")
}
