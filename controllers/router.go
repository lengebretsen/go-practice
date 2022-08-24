package controllers

import (
	"github.com/lengebretsen/go-practice/models"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

type handler struct {
	users     models.UserRepository
	addresses models.AddressRepository
}

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	// docs route
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

// RegisterRoutes initializes the routes and sets up the handler's reference to the model(s) for database access
func RegisterRoutes(r *gin.Engine, users models.UserRepository, addresses models.AddressRepository) {
	h := &handler{
		users:     users,
		addresses: addresses,
	}

	userRoutes := r.Group("/users")
	userRoutes.POST("/", h.AddUser)
	userRoutes.GET("/", h.FetchUsers)
	userRoutes.GET("/:id", h.FetchUser)
	userRoutes.PUT("/:id", h.UpdateUser)
	userRoutes.DELETE("/:id", h.DeleteUser)
	userRoutes.GET("/:id/addresses", h.FetchAddressesForUser)

	addressRoutes := r.Group("/addresses")
	addressRoutes.POST("/", h.AddAddress)
	addressRoutes.GET("/", h.FetchAddresses)
	addressRoutes.GET("/:id", h.FetchAddress)
	addressRoutes.PUT("/:id", h.UpdateAddress)
	addressRoutes.DELETE("/:id", h.DeleteAddress)
}
