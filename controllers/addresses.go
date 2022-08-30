package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/lengebretsen/go-practice/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type addUpdateAddressBody struct {
	UserId uuid.UUID `json:"userId" binding:"required"`
	Street string    `json:"street"`
	City   string    `json:"city"`
	State  string    `json:"state"`
	Zip    string    `json:"zip"`
	Type   string    `json:"type"`
}

// FetchAddresses retrieves a list of all addresses in the system
// @Summary retrieve a list of all addresses in the system
// @Tags addresses
// @ID fetch-all-addrs
// @Produce json
// @Success 200 {object} []models.Address
// @Router /addresses [get]
func (h handler) FetchAddresses(c *gin.Context) {
	addrs, err := h.addresses.FetchAddresses()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, ApiError{Message: "Error fetching address records", Error: err})
		return
	}
	c.IndentedJSON(http.StatusOK, addrs)
}

// FetchAddress retrieves a single address by Id
// @Summary retrieve an address by Id
// @Tags addresses
// @ID fetch-addr
// @Produce json
// @Param id path string true "address ID"
// @Success 200 {object} models.Address
// @Failure 400 {object} ApiError
// @Failure 404 {object} ApiError
// @Router /addresses/{id} [get]
func (h handler) FetchAddress(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, ApiError{Message: fmt.Sprintf("Id [%s] is not a valid UUID", idParam), Error: err})
		return
	}

	addr, err := h.addresses.FetchOneAddress(id)
	if err != nil {
		if errors.Is(err, models.ErrModelNotFound) {
			c.IndentedJSON(http.StatusNotFound, ApiError{Message: fmt.Sprintf("No address exists with Id [%s]", idParam), Error: err})
			return
		} else {
			c.IndentedJSON(http.StatusInternalServerError, ApiError{Message: fmt.Sprintf("Error fetching address record with Id [%s]", id), Error: err})
			return
		}
	}
	c.IndentedJSON(http.StatusOK, addr)
}

// FetchAddressesForUser retrieves a list of addresses associated with a user
// @Summary retrieve a list of addresses by the user's Id
// @Tags users, addresses
// @ID fetch-addrs-for-user
// @Produce json
// @Param id path string true "user ID"
// @Success 200 {object} []models.Address
// @Failure 400 {object} ApiError
// @Failure 404 {object} ApiError
// @Router /users/{id}/addresses [get]
func (h handler) FetchAddressesForUser(c *gin.Context) {
	idParam := c.Param("id")
	userId, err := uuid.Parse(idParam)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, ApiError{Message: fmt.Sprintf("Id [%s] is not a valid UUID", idParam), Error: err})
		return
	}

	//lookup user to make sure they exist, and send back 404 if they do not
	_, err = h.users.SelectOneUser(userId)
	if err != nil {
		if errors.Is(err, models.ErrModelNotFound) {
			c.IndentedJSON(http.StatusNotFound, ApiError{Message: fmt.Sprintf("No user exists with Id [%s]", idParam), Error: err})
			return
		} else {
			c.IndentedJSON(http.StatusInternalServerError, ApiError{Message: fmt.Sprintf("Error fetching address records for user [%s]", idParam), Error: err})
			return
		}
	}

	addrs, err := h.addresses.FindAddressesByUserId(userId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, ApiError{Message: fmt.Sprintf("Error fetching address records for user [%s]", idParam), Error: err})
		return
	}
	c.IndentedJSON(http.StatusOK, addrs)
}

// AddAddress stores a new address
// @Summary store a new address
// @Tags addresses
// @ID add-addr
// @Produce json
// @Param data body addUpdateAddressBody true "new address data"
// @Success 200 {object} []models.Address
// @Failure 400 {object} ApiError
// @Failure 404 {object} ApiError
// @Router /addresses [post]
func (h handler) AddAddress(c *gin.Context) {
	var reqBody addUpdateAddressBody

	err := c.BindJSON(&reqBody)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, ApiError{Message: "Invalid request body.", Error: err})
		return
	}

	//lookup user to make sure they exist, and send back 404 if they do not
	_, err = h.users.SelectOneUser(reqBody.UserId)
	if err != nil {
		if errors.Is(err, models.ErrModelNotFound) {
			c.IndentedJSON(http.StatusNotFound, ApiError{Message: fmt.Sprintf("No user exists with Id [%s]", reqBody.UserId), Error: err})
			return
		} else {
			c.IndentedJSON(http.StatusInternalServerError, ApiError{Message: "Error creating new address", Error: err})
			return
		}
	}

	newAddr, err := h.addresses.InsertAddress(models.Address{
		Id:     uuid.New(),
		UserId: reqBody.UserId,
		Street: reqBody.Street,
		City:   reqBody.City,
		State:  reqBody.State,
		Zip:    reqBody.Zip,
		Type:   reqBody.Type,
	})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, ApiError{Message: "Error creating new address", Error: err})
		return
	}

	c.IndentedJSON(http.StatusCreated, newAddr)
}

// AddAddress updates an existing address
// @Summary update an existing address by Id
// @Tags addresses
// @ID update-addr
// @Produce json
// @Param id path string true "address ID"
// @Param data body addUpdateAddressBody true "updated address data"
// @Success 200 {object} []models.Address
// @Failure 400 {object} ApiError
// @Failure 404 {object} ApiError
// @Router /addresses/{id} [put]
func (h handler) UpdateAddress(c *gin.Context) {
	var reqBody addUpdateAddressBody
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, ApiError{Message: fmt.Sprintf("Id [%s] is not a valid UUID", idParam), Error: err})
		return
	}

	err = c.BindJSON(&reqBody)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, ApiError{Message: "Invalid request body.", Error: err})
		return
	}

	//lookup user to make sure they exist, and send back 404 if they do not
	_, err = h.users.SelectOneUser(reqBody.UserId)
	if err != nil {
		if errors.Is(err, models.ErrModelNotFound) {
			c.IndentedJSON(http.StatusNotFound, ApiError{Message: fmt.Sprintf("No user exists with Id [%s]", reqBody.UserId), Error: err})
			return
		} else {
			c.IndentedJSON(http.StatusInternalServerError, ApiError{Message: fmt.Sprintf("Error updating address record with Id [%s]", id), Error: err})
			return
		}
	}

	updatedAddr, err := h.addresses.UpdateAddress(
		models.Address{Id: id, UserId: reqBody.UserId, Street: reqBody.Street, City: reqBody.City, State: reqBody.State, Zip: reqBody.Zip, Type: reqBody.Type},
	)
	if err != nil {
		if errors.Is(err, models.ErrModelNotFound) {
			c.IndentedJSON(http.StatusNotFound, ApiError{Message: fmt.Sprintf("No address exists with Id [%s]", idParam), Error: err})
			return
		} else {
			c.IndentedJSON(http.StatusInternalServerError, ApiError{Message: fmt.Sprintf("Error updating address record with Id [%s]", id), Error: err})
			return
		}
	}

	c.IndentedJSON(http.StatusOK, updatedAddr)
}

// DeleteAddress deletes an existing address
// @Summary remove an existing address by Id
// @Tags addresses
// @ID delete-addr
// @Produce json
// @Param id path string true "address ID"
// @Success 204
// @Failure 400 {object} ApiError
// @Failure 404 {object} ApiError
// @Router /addresses/{id} [delete]
func (h handler) DeleteAddress(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, ApiError{Message: fmt.Sprintf("Id [%s] is not a valid UUID", idParam), Error: err})
		return
	}

	err = h.addresses.DeleteAddress(id)
	if err != nil {
		if errors.Is(err, models.ErrModelNotFound) {
			c.IndentedJSON(http.StatusNotFound, ApiError{Message: fmt.Sprintf("No address exists with Id [%s]", idParam), Error: err})
			return
		} else {
			c.IndentedJSON(http.StatusInternalServerError, ApiError{Message: fmt.Sprintf("Error deleting address record with Id [%s]", id), Error: err})
			return
		}
	}
	c.Status(http.StatusNoContent)
}
