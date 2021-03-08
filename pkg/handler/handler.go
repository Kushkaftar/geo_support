package handler

import (
	"github.com/Kushkaftar/geo_support/pkg/service"
	"github.com/gin-gonic/gin"

	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	"github.com/swaggo/gin-swagger/swaggerFiles"

	// swagger embed files

	_ "github.com/Kushkaftar/geo_support/docs"
)

// Handler ...
type Handler struct {
	services *service.Service
}

// NewHandler ...
func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

// InitRoutes ...
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
		}
		prices := api.Group("/prices")
		{
			prices.GET("/", h.getAllPrices)
			prices.POST("/set_price", h.setPrices)
			prices.POST("/update_price", h.updatePrices)
		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}

func cors(c *gin.Context) {
	c.Header("Content-Type", "application/json")
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST, GET")
}
