package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Kushkaftar/geo_support/modelsstruct"
	"github.com/gin-gonic/gin"
)

// @Summary Set flag to domain
// @Description set_flag is out parameters: 0 - new(set automatic), 1 - active, 2 - ignore
// @ID set-flag-domain
// @Accept  json
// @Produce  json
// @Param id query int true "id domain"
// @Param input body modelsstruct.Flag true "set flag"
// @Success 200
// @Router /api/domains/{id} [post]
// setFlagDomain ...
func (h *Handler) setFlagDomain(c *gin.Context) {
	var f modelsstruct.Flag

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

	err = h.services.PreSetFlag(f.SetFlag, id)

	// TODO: delete fmt
	fmt.Println("f.SetFlag", f.SetFlag)
	fmt.Println("id", id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "set flag is fail "+err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"setFlagDomain": "ok",
	})
}

// @Summary Get all domains
// @Description Show all folders in dirrectory
// @ID get-all-domains
// @Accept  json
// @Produce  json
// @Success 200 {object} modelsstruct.Domains
// @Router /api/domains/ [post]
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
