package controllers

import (
	"engebretsen/simple_web_svc/models"
	"engebretsen/simple_web_svc/pkg/common"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AddUpdateUserBody struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

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

func (h handler) AddUser(c *gin.Context) {
	var reqBody AddUpdateUserBody

	if err := c.BindJSON(&reqBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, common.ApiError{Message: "Invalid request body.", Error: err})
	}
	newUser, err := h.users.InsertUser(models.User{Id: uuid.New(), FirstName: reqBody.FirstName, LastName: reqBody.LastName})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, common.ApiError{Message: "Error creating new user", Error: err})
	}

	c.IndentedJSON(http.StatusCreated, newUser)
}

func (h handler) UpdateUser(c *gin.Context) {
	var reqBody AddUpdateUserBody
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

func (h handler) DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, common.ApiError{Message: fmt.Sprintf("Id [%s] is not a valid UUID", idParam), Error: err})
	}
	err = h.users.DeleteUser(id)
	if err != nil {
		switch e := err.(type) {
		case *models.ErrUserNotFound:
			c.IndentedJSON(http.StatusNotFound, common.ApiError{Message: fmt.Sprintf("No user exists with Id [%s]", idParam), Error: e})
		default:
			c.IndentedJSON(http.StatusInternalServerError, common.ApiError{Message: fmt.Sprintf("Error deleting user record with Id [%s]", id), Error: err})
		}
	}
	c.Status(http.StatusNoContent)
}
