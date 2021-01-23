package handler

import (
	"fmt"
	"github.com/Kushkaftar/geo_support/modelsStruct"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) setFlagDomain(c *gin.Context) {
	var d modelsStruct.Domain
	var f modelsStruct.Flag

	domainId := c.Param("id")
	id, err := strconv.Atoi(domainId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid param id")
		return
	}
	d.Id = id

	if err := c.BindJSON(&f); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}
	d.Flag = f.SetFlag
	fmt.Println(d)

	c.JSON(http.StatusOK, map[string]interface{}{
		"setFlagDomain": "test",
	})
}

func (h *Handler) getAllDomains(c *gin.Context) {
	var ds modelsStruct.Domains
	fmt.Println("qqq")
	domains, err := h.services.Domains.GetAllDomains()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "что-то пошло не так...")
		return
	}
	fmt.Println(domains)
	ds.Domains = domains
	c.JSON(http.StatusOK, ds)
}
