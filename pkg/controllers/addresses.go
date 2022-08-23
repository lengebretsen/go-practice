package controllers

import (
	"engebretsen/simple_web_svc/models"
	"engebretsen/simple_web_svc/pkg/common"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AddAddressBody struct {
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
	}
	c.IndentedJSON(http.StatusOK, addrs)
}

// AddUser stores a new user
func (h handler) AddAddress(c *gin.Context) {
	var reqBody AddAddressBody

	if err := c.BindJSON(&reqBody); err != nil {
		c.IndentedJSON(http.StatusBadRequest, common.ApiError{Message: "Invalid request body.", Error: err})
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
	}

	c.IndentedJSON(http.StatusCreated, newAddr)
}
