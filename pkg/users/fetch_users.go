package users

import (
	"engebretsen/simple_web_svc/models"
	"engebretsen/simple_web_svc/pkg/common"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h handler) FetchUsers(c *gin.Context) {
	users, err := h.users.SelectAllUsers()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, common.ApiError{Message: "Error fetching user records", Error: err})
	}
	c.IndentedJSON(http.StatusOK, users)
}

func (h handler) FetchUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, common.ApiError{Message: fmt.Sprintf("Id [%s] is not a valid UUID", idParam), Error: err})
	}

	user, err := h.users.SelectOneUser(id)
	if err != nil {
		switch e := err.(type) {
		case *models.ErrUserNotFound:
			c.IndentedJSON(http.StatusNotFound, common.ApiError{Message: fmt.Sprintf("No user exists with Id [%s]", idParam), Error: e})
		default:
			c.IndentedJSON(http.StatusInternalServerError, common.ApiError{Message: fmt.Sprintf("Error fetching user record with Id [%s]", id), Error: err})
		}
	}
	c.IndentedJSON(http.StatusOK, user)
}
