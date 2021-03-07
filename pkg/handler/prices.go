package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllPrices(c *gin.Context) {
	prices, err := h.services.GetAllPrices()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "что-то пошло не так..."+err.Error())
		return
	}

	c.JSON(http.StatusOK, prices)
}

func (h *Handler) setPrices(c *gin.Context) {

}
