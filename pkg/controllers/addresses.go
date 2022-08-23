package controllers

import (
	"engebretsen/simple_web_svc/models"
	"engebretsen/simple_web_svc/pkg/common"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AddUpdateAddressBody struct {
	UserId uuid.UUID `json:"userId"`
	Street string    `json:"street"`
	City   string    `json:"city"`
	State  string    `json:"state"`
	Zip    string    `json:"zip"`
	Type   string    `json:"type"`
}

func (h handler) FetchAddresses(c *gin.Context) {
	addrs, err := h.addresses.FetchAddresses()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, common.ApiError{Message: "Error fetching address records", Error: err})
		return
	}
	c.IndentedJSON(http.StatusOK, addrs)
}

func (h handler) FetchAddress(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, common.ApiError{Message: fmt.Sprintf("Id [%s] is not a valid UUID", idParam), Error: err})
		return
	}

	addr, err := h.addresses.FetchOneAddress(id)
	if err != nil {
		switch e := err.(type) {
		case *models.ErrModelNotFound:
			c.IndentedJSON(http.StatusNotFound, common.ApiError{Message: fmt.Sprintf("No address exists with Id [%s]", idParam), Error: e})
			return
		default:
			c.IndentedJSON(http.StatusInternalServerError, common.ApiError{Message: fmt.Sprintf("Error fetching address record with Id [%s]", id), Error: err})
			return
		}
	}
	c.IndentedJSON(http.StatusOK, addr)
}

func (h handler) FetchAddressesForUser(c *gin.Context) {
	idParam := c.Param("id")
	userId, err := uuid.Parse(idParam)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, common.ApiError{Message: fmt.Sprintf("Id [%s] is not a valid UUID", idParam), Error: err})
		return
	}

	addrs, err := h.addresses.FindAddressesByUserId(userId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, common.ApiError{Message: fmt.Sprintf("Error fetching address records for user [%s]", idParam), Error: err})
		return
	}
	c.IndentedJSON(http.StatusOK, addrs)
}

// AddAddress stores a new address for a user
func (h handler) AddAddress(c *gin.Context) {
	var reqBody AddUpdateAddressBody

	if err := c.BindJSON(&reqBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, common.ApiError{Message: "Invalid request body.", Error: err})
		return
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
		c.IndentedJSON(http.StatusInternalServerError, common.ApiError{Message: "Error creating new address", Error: err})
		return
	}

	c.IndentedJSON(http.StatusCreated, newAddr)
}

func (h handler) UpdateAddress(c *gin.Context) {
	var reqBody AddUpdateAddressBody
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

	updatedAddr, err := h.addresses.UpdateAddress(
		models.Address{Id: id, UserId: reqBody.UserId, Street: reqBody.Street, City: reqBody.City, State: reqBody.State, Zip: reqBody.Zip, Type: reqBody.Type},
	)
	if err != nil {
		switch e := err.(type) {
		case *models.ErrModelNotFound:
			c.IndentedJSON(http.StatusNotFound, common.ApiError{Message: fmt.Sprintf("No address exists with Id [%s]", idParam), Error: e})
			return
		default:
			c.IndentedJSON(http.StatusInternalServerError, common.ApiError{Message: fmt.Sprintf("Error updating address record with Id [%s]", id), Error: err})
			return
		}
	}

	c.IndentedJSON(http.StatusOK, updatedAddr)
}

func (h handler) DeleteAddress(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, common.ApiError{Message: fmt.Sprintf("Id [%s] is not a valid UUID", idParam), Error: err})
		return
	}

	err = h.addresses.DeleteAddress(id)
	if err != nil {
		switch e := err.(type) {
		case *models.ErrModelNotFound:
			c.IndentedJSON(http.StatusNotFound, common.ApiError{Message: fmt.Sprintf("No address exists with Id [%s]", idParam), Error: e})
			return
		default:
			c.IndentedJSON(http.StatusInternalServerError, common.ApiError{Message: fmt.Sprintf("Error deleting address record with Id [%s]", id), Error: err})
			return
		}
	}
	c.Status(http.StatusNoContent)
}
