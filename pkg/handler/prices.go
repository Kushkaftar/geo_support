package handler

import (
	"net/http"

	"github.com/Kushkaftar/geo_support/modelsStruct"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllPrices(c *gin.Context) {
	var ps modelsStruct.Prices
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
