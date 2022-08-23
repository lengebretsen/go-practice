package controllers

import (
	"engebretsen/simple_web_svc/models"

	"github.com/gin-gonic/gin"
)

type handler struct {
	users     models.UserRepository
	addresses models.AddressRepository
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

	addressRoutes := r.Group("/addresses")
	addressRoutes.POST("/", h.AddAddress)
	addressRoutes.GET("/", h.FetchAddresses)
}
