package users

import (
	"engebretsen/simple_web_svc/models"
	"engebretsen/simple_web_svc/pkg/common"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UpdateUserBody struct {
	FirstName string
	LastName  string
}

func (h handler) UpdateUser(c *gin.Context) {
	var reqBody UpdateUserBody
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, common.ApiError{Message: fmt.Sprintf("Id [%s] is not a valid UUID", idParam), Error: err})
	}

	if err := c.BindJSON(&reqBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, common.ApiError{Message: "Invalid request body.", Error: err})
	}

	updatedUser, err := h.users.UpdateUser(models.User{Id: id, FirstName: reqBody.FirstName, LastName: reqBody.LastName})
	if err != nil {
		switch e := err.(type) {
		case *models.ErrUserNotFound:
			c.IndentedJSON(http.StatusNotFound, common.ApiError{Message: fmt.Sprintf("No user exists with Id [%s]", idParam), Error: e})
		default:
			c.IndentedJSON(http.StatusInternalServerError, common.ApiError{Message: fmt.Sprintf("Error updating user record with Id [%s]", id), Error: err})
		}
	}

	c.IndentedJSON(http.StatusOK, updatedUser)
}
