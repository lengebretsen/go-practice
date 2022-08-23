package users

import (
	"engebretsen/simple_web_svc/models"

	"github.com/gin-gonic/gin"
)

type handler struct {
	users models.UserModel
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
