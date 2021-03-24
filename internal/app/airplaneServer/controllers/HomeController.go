package controllers

import (
	"airplaneServer/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	// "time"
)

func NewHomeController() *HomeController {
	return &HomeController{}
}

type HomeController struct {
	BaseController
}

func (controller HomeController) Ping(context *gin.Context) {
	// time.Sleep(2 * time.Second)
	logger.Info("ping~~~", logger.DEFAULT)
	context.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
