package controllers

import (
	"engebretsen/simple_web_svc/models"
	"engebretsen/simple_web_svc/pkg/common"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type addUpdateUserBody struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// FetchUsers retrieves a list of all users in the system
// @Summary retrieve a list of all users in the system
// @Tags users
// @ID fetch-all-users
// @Produce json
// @Success 200 {object} []models.User
// @Router /users [get]
func (h handler) FetchUsers(c *gin.Context) {
	users, err := h.users.SelectAllUsers()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, common.ApiError{Message: "Error fetching user records", Error: err})
		return
	}
	c.IndentedJSON(http.StatusOK, users)
}

// FetchUser retrieves a single user by id
// @Summary retrieve a user by Id
// @Tags users
// @ID fetch-user
// @Produce json
// @Param id path string true "user ID"
// @Success 200 {object} models.User
// @Failure 400 {object} common.ApiError
// @Failure 404 {object} common.ApiError
// @Router /users/{id} [get]
func (h handler) FetchUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, common.ApiError{Message: fmt.Sprintf("Id [%s] is not a valid UUID", idParam), Error: err})
		return
	}

	user, err := h.users.SelectOneUser(id)
	if err != nil {
		switch e := err.(type) {
		case *models.ErrModelNotFound:
			c.IndentedJSON(http.StatusNotFound, common.ApiError{Message: fmt.Sprintf("No user exists with Id [%s]", idParam), Error: e})
			return
		default:
			c.IndentedJSON(http.StatusInternalServerError, common.ApiError{Message: fmt.Sprintf("Error fetching user record with Id [%s]", id), Error: err})
			return
		}
	}
	c.IndentedJSON(http.StatusOK, user)
}

// AddUser stores a new user
// @Summary add a new user
// @Tags users
// @ID add-user
// @Produce json
// @Param data body addUpdateUserBody true "new user data"
// @Success 200 {object} models.User
// @Failure 400 {object} common.ApiError
// @Router /users [post]
func (h handler) AddUser(c *gin.Context) {
	var reqBody addUpdateUserBody

	if err := c.BindJSON(&reqBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, common.ApiError{Message: "Invalid request body.", Error: err})
		return
	}
	newUser, err := h.users.InsertUser(models.User{Id: uuid.New(), FirstName: reqBody.FirstName, LastName: reqBody.LastName})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, common.ApiError{Message: "Error creating new user", Error: err})
		return
	}

	c.IndentedJSON(http.StatusCreated, newUser)
}

// UpdateUser modifies an existing user
// @Summary modify an existing user
// @Tags users
// @ID update-user
// @Produce json
// @Param id path string true "user ID"
// @Param data body addUpdateUserBody true "new user data"
// @Success 200 {object} models.User
// @Failure 400 {object} common.ApiError
// @Failure 404 {object} common.ApiError
// @Router /users/{id} [put]
func (h handler) UpdateUser(c *gin.Context) {
	var reqBody addUpdateUserBody
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, common.ApiError{Message: fmt.Sprintf("Id [%s] is not a valid UUID", idParam), Error: err})
		return
	}

	if err := c.BindJSON(&reqBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, common.ApiError{Message: "Invalid request body.", Error: err})
		return
	}

	updatedUser, err := h.users.UpdateUser(models.User{Id: id, FirstName: reqBody.FirstName, LastName: reqBody.LastName})
	if err != nil {
		switch e := err.(type) {
		case *models.ErrModelNotFound:
			c.IndentedJSON(http.StatusNotFound, common.ApiError{Message: fmt.Sprintf("No user exists with Id [%s]", idParam), Error: e})
			return
		default:
			c.IndentedJSON(http.StatusInternalServerError, common.ApiError{Message: fmt.Sprintf("Error updating user record with Id [%s]", id), Error: err})
			return
		}
	}

	c.IndentedJSON(http.StatusOK, updatedUser)
}

// DeleteUser deletes an existing user, including any addresses associated with the user
// @Summary delete a user by Id, including any addresses associated with the user
// @Tags users
// @ID delete-user
// @Param id path string true "user ID"
// @Success 204
// @Failure 400 {object} common.ApiError
// @Failure 404 {object} common.ApiError
// @Router /users/{id} [delete]
func (h handler) DeleteUser(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, common.ApiError{Message: fmt.Sprintf("Id [%s] is not a valid UUID", idParam), Error: err})
		return
	}
	err = h.users.DeleteUser(id)
	if err != nil {
		switch e := err.(type) {
		case *models.ErrModelNotFound:
			c.IndentedJSON(http.StatusNotFound, common.ApiError{Message: fmt.Sprintf("No user exists with Id [%s]", idParam), Error: e})
			return
		default:
			c.IndentedJSON(http.StatusInternalServerError, common.ApiError{Message: fmt.Sprintf("Error deleting user record with Id [%s]", id), Error: err})
			return
		}
	}
	c.Status(http.StatusNoContent)
}
