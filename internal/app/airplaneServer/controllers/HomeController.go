package controllers

import (
	"airplaneServer/internal/app/airplaneServer/constants"
	"airplaneServer/internal/app/airplaneServer/services"
	"airplaneServer/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
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

	token, error := h.UserService.RegOrLogin(username, password)
	if error == nil {
		c.JSON(http.StatusOK, h.MakeResponse(constants.ResCodeSuccess, "RegDone",
			map[string]interface{}{
				"token": token,
			}))
	} else {
		c.JSON(http.StatusOK, h.MakeResponse(constants.ResCodeOther, error.Error(), h.EmptyData()))
	}

}
