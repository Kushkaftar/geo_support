package handler

import (
	"net/http"

	"github.com/Kushkaftar/geo_support/modelsstruct"
	"github.com/gin-gonic/gin"
)

// @Summary Get all prices
// @Description Get all prices
// @ID get-all-prices
// @Accept  json
// @Produce  json
// @Success 200 {object} modelsstruct.Prices
// @Router /api/prices/ [get]
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

// @Summary Set price
// @Description set price
// @ID set-price
// @Accept  json
// @Produce  json
// @Param input body modelsstruct.Price true "required"
// @Success 200 {json} {"set_price": "ok"}
// @Router /api/prices/set_price [post]
func (h *Handler) setPrices(c *gin.Context) {
	var p modelsstruct.Price

	if err := c.BindJSON(&p); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body "+err.Error())
		return
	}

	check, err := h.services.SetPrice(p)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "some error "+err.Error())
		return
	}

	if check != 1 {
		newErrorResponse(c, http.StatusBadRequest, "err "+err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"set_price": "ok",
	})
}

// @Summary Update price
// @Description Update price
// @ID update-price
// @Accept  json
// @Produce  json
// @Param input body modelsstruct.Price true "required"
// @Success 200 {json} {"update_price": "ok"}
// @Router /api/prices/update_price [post]
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
