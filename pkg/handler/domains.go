package handler

import (
	"net/http"
	"strconv"

	"github.com/Kushkaftar/geo_support/modelsstruct"
	"github.com/gin-gonic/gin"
)

// setFlagDomain ...
func (h *Handler) setFlagDomain(c *gin.Context) {
	var f modelsstruct.Flag

	//c.Header("Access-Control-Allow-Origin", "*")
	//c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")

	domainID := c.Param("id")
	id, err := strconv.Atoi(domainID)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid param id"+err.Error())
		return
	}

	if err := c.BindJSON(&f); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body"+err.Error())
		return
	}

	if f.SetFlag < 0 || f.SetFlag > 2 {
		newErrorResponse(c, http.StatusBadRequest,
			"set_flag is out parameters: 0 - new(set automatic), 1 - active, 2 - ignore")
		return
	}

	err = h.services.SetFlag(f.SetFlag, id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "set flag is fail"+err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"setFlagDomain": "ok",
	})
}

// getAllDomains ...
func (h *Handler) getAllDomains(c *gin.Context) {
	var ds modelsstruct.Domains

	domains, err := h.services.Domains.GetAllDomains()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "что-то пошло не так..."+err.Error())
		return
	}

	ds.Domains = domains
	c.JSON(http.StatusOK, ds)
}
