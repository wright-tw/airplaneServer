package controllers

import (
	"airplaneServer/internal/app/airplaneServer/constants"
	"airplaneServer/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewHomeController() *HomeController {
	return &HomeController{}
}

type HomeController struct {
	BaseController
}

func (h HomeController) Ping(c *gin.Context) {
	logger.Info("ping~~~", logger.DEFAULT)
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (h HomeController) RegOrLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	logger.Info(username, logger.DEFAULT)
	logger.Info(password, logger.DEFAULT)
	c.JSON(http.StatusOK, h.MakeResponse(constants.ResCodeSuccess, "RegDone", h.EmptyData()))
}
