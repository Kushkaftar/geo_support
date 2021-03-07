package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Kushkaftar/geo_support/modelsstruct"
	"github.com/gin-gonic/gin"
)

func (h *Handler) setName(c *gin.Context) {
	var n modelsstruct.Name

	domainId := c.Param("id")

	id, err := strconv.Atoi(domainId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid param id"+err.Error())
		return
	}
	n.Id = id
	fmt.Println(n)

	if err = c.BindJSON(&n); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body"+err.Error())
		return
	}
	fmt.Println(n)

	// временная заглушка
	c.JSON(http.StatusOK, map[string]interface{}{
		"setNameOffer": "Ok",
	})
}

func (h *Handler) getName(c *gin.Context) {

	// временная заглушка
	c.JSON(http.StatusOK, map[string]interface{}{
		"title": "test_offer",
	})
}

func (h *Handler) getOffers(c *gin.Context) {

	// временная заглушка
	c.JSON(http.StatusOK, map[string]interface{}{
		"title": "test_offer",
	})
}
