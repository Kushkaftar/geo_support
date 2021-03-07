package handler

import (
	"net/http"

	"github.com/Kushkaftar/geo_support/modelsstruct"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllPrices(c *gin.Context) {
	var ps modelsstruct.Prices

	prices, err := h.services.GetAllPrices()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "что-то пошло не так..."+err.Error())
		return
	}

	ps.Prices = prices
	c.JSON(http.StatusOK, ps)
}

func (h *Handler) setPrices(c *gin.Context) {

}
