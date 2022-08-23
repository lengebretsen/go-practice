package controllers

import (
	"engebretsen/simple_web_svc/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type handler struct {
	users interface {
		SelectAllUsers() ([]models.User, error)
		SelectOneUser(id uuid.UUID) (models.User, error)
		InsertUser(usr models.User) (models.User, error)
		UpdateUser(usr models.User) (models.User, error)
		DeleteUser(id uuid.UUID) error
	}
}

func RegisterRoutes(r *gin.Engine, users models.UserModel) {
	h := &handler{
		users: users,
	}

	routes := r.Group("/users")
	routes.POST("/", h.AddUser)
	routes.GET("/", h.FetchUsers)
	routes.GET("/:id", h.FetchUser)
	routes.PUT("/:id", h.UpdateUser)
	routes.DELETE("/:id", h.DeleteUser)
}
