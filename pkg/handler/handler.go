package handler

import (
	"github.com/Kushkaftar/geo_support/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		domains := api.Group("/domains")
		{
			domains.POST("/:id", h.setFlagDomain)
			domains.GET("/", h.getAllDomains)

			prices := domains.Group(":id/prices")
			{
				prices.GET("/", h.getAllPrices)
			}
		}
	}

	return router
}