package handler

import (
	"github.com/Bloodstein/dns-shop-subscriber-backend/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	subscribersHandler SubscribersHandler
}

type SubscribersHandler interface {
	GetAll(ctx *gin.Context)
	GetOne(ctx *gin.Context)
	Create(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Monitoring(ctx *gin.Context)
	PricesListening()
}

func (h *Handler) Routes() *gin.Engine {
	router := gin.New()

	firstApi := router.Group("api/v1")
	{
		subsApi := firstApi.Group("subscribes")
		{
			subsApi.GET("get-all", h.subscribersHandler.GetAll)
			subsApi.GET("get-one/:id", h.subscribersHandler.GetOne)
			subsApi.POST("create", h.subscribersHandler.Create)
			subsApi.POST("delete/:id", h.subscribersHandler.Delete)
		}
		firstApi.GET("monitoring/:id", h.subscribersHandler.Monitoring)
	}

	return router
}

func (h *Handler) RunPricesListening() {
	h.subscribersHandler.PricesListening()
}

func NewHandler(srv service.Service) Handler {
	return Handler{
		subscribersHandler: NewSubscribersHandler(srv),
	}
}
