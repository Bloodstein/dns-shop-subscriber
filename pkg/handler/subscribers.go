package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Bloodstein/dns-shop-subscriber-backend/domain"
	"github.com/Bloodstein/dns-shop-subscriber-backend/pkg/service"
	"github.com/gin-gonic/gin"
)

type Subscribers struct {
	services service.Service
}

func NewSubscribersHandler(srv service.Service) Subscribers {
	return Subscribers{
		services: srv,
	}
}

func (this Subscribers) GetAll(ctx *gin.Context) {
	all := this.services.SubscribersService.GetAll()
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"response": "ok",
		"records":  all,
	})
}

func (this Subscribers) GetOne(ctx *gin.Context) {
	rowId, err := this.getRowId(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"response": "error",
			"message":  fmt.Sprintf("incorrect row id. %s", err.Error()),
		})
		return
	}

	rec := this.services.SubscribersService.GetOne(rowId)

	if rec == nil {
		ctx.JSON(http.StatusNotFound, map[string]string{
			"response": "not found",
			"message":  "subscribe with that ID not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"response": "ok",
		"id":       rec,
	})
}

func (this Subscribers) Create(ctx *gin.Context) {
	var body domain.CreateNewRequest
	err := ctx.BindJSON(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"response": "error",
			"message":  fmt.Sprintf("parse body error. %s", err.Error()),
		})
		return
	}

	newRowId, err := this.services.SubscribersService.Create(&body)

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"response": "ok",
		"id":       newRowId,
	})
}

func (this Subscribers) Delete(ctx *gin.Context) {
	rowId, err := this.getRowId(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"response": "error",
			"message":  fmt.Sprintf("incorrect row id. %s", err.Error()),
		})
		return
	}

	err = this.services.SubscribersService.Delete(rowId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"response": "error",
			"message":  fmt.Sprintf("fail to delete. %s", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"response": "ok",
		"id":       rowId,
	})
}

func (this Subscribers) Monitoring(ctx *gin.Context) {
	rowId, err := this.getRowId(ctx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"response": "error",
			"message":  fmt.Sprintf("incorrect row id. %s", err.Error()),
		})
		return
	}

	err = this.services.SubscribersService.Monitoring(rowId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, map[string]string{
			"response": "error",
			"message":  fmt.Sprintf("error to load monitoring. %s", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"response": "ok",
		"result":   make([]*domain.Monitoring, 1),
	})
}

func (this Subscribers) PricesListening() {
	for {
		log.Println("prices listening...")
		time.Sleep(time.Second * 30)
	}
}

func (this Subscribers) getRowId(ctx *gin.Context) (int, error) {
	idParam := ctx.Param("id")
	return strconv.Atoi(idParam)
}
