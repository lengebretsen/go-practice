package users

import (
	"engebretsen/simple_web_svc/models"
	"engebretsen/simple_web_svc/pkg/common"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AddUserBody struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (h handler) AddUser(c *gin.Context) {
	var reqBody AddUserBody

	if err := c.BindJSON(&reqBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, common.ApiError{Message: "Invalid request body.", Error: err})
	}
	newUser, err := h.users.InsertUser(models.User{Id: uuid.New(), FirstName: reqBody.FirstName, LastName: reqBody.LastName})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, common.ApiError{Message: "Error creating new user", Error: err})
	}

	c.IndentedJSON(http.StatusCreated, newUser)
}
