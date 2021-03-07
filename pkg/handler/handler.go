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
	//gin.SetMode(gin.ReleaseMode)
	router.Use(cors)
	api := router.Group("/api")
	{
		domains := api.Group("/domains")
		{
			domains.POST("/:id", h.setFlagDomain)
			domains.POST("/", h.getAllDomains)

			offer := domains.Group(":id/offer")
			{
				offer.POST("/set_name", h.setName)
				offer.GET("/get_name", h.getName)
				offer.GET("/", h.getOffers)
			}

			prices := domains.Group(":id/prices")
			{
				prices.GET("/", h.getAllPrices)
			}
		}
	}

	return router
}

func cors(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET")
}