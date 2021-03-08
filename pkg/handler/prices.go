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

func (h *Handler) updatePrices(c *gin.Context) {
	var p modelsstruct.Price

	if err := c.BindJSON(&p); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body"+err.Error())
		return
	}

	check, err := h.services.UpdatePrice(p)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "some error"+err.Error())
		return
	}

	if check == 0 {
		newErrorResponse(c, http.StatusBadRequest, "price not found")
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"update_price": "ok",
	})
}
