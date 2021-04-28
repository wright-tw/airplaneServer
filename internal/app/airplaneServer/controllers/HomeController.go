package controllers

import (
	"airplaneServer/internal/app/airplaneServer/constants"
	"airplaneServer/internal/app/airplaneServer/services"
	"airplaneServer/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func NewHomeController(service *services.UserService) *HomeController {
	return &HomeController{
		UserService: service,
	}
}

type HomeController struct {
	BaseController
	UserService services.IUserService
}

func (h *HomeController) Ping(c *gin.Context) {
	logger.Info("ping~~~", logger.DEFAULT)
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (h *HomeController) RegOrLogin(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	logger.Info(username+" Login....", logger.DEFAULT)

	token, err := h.UserService.RegOrLogin(username, password)
	if err == nil {
		c.JSON(http.StatusOK, h.MakeResponse(constants.ResCodeSuccess, "success",
			map[string]interface{}{
				"token": token,
			}))
	} else {
		c.JSON(http.StatusOK, h.MakeResponse(constants.ResCodeOther, err.Error(), h.EmptyData()))
	}
}

func (h *HomeController) Score(c *gin.Context) {
	token := c.PostForm("token")
	score, err := strconv.ParseInt(c.PostForm("score"), 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, h.MakeResponse(constants.ResCodeOther, err.Error(), h.EmptyData()))
		return
	}

	err = h.UserService.WriteScore(token, score)
	if err == nil {
		c.JSON(http.StatusOK, h.MakeResponse(constants.ResCodeSuccess, "success", h.EmptyData()))
	} else {
		c.JSON(http.StatusOK, h.MakeResponse(constants.ResCodeOther, err.Error(), h.EmptyData()))
	}
}
