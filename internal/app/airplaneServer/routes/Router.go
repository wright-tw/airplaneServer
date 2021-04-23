package routes

import (
	"airplaneServer/internal/app/airplaneServer/controllers"
	"airplaneServer/internal/app/airplaneServer/middlewares"
	"github.com/gin-gonic/gin"
)

func NewRouter(
	homeController *controllers.HomeController,
) *Router {
	return &Router{
		HomeController: homeController,
	}
}

type Router struct {
	HomeController *controllers.HomeController
}

func (r *Router) Setting(server *gin.Engine) {
	// ping
	server.GET("ping", r.HomeController.Ping)
	server.POST("reg-or-login", r.HomeController.RegOrLogin)
	server.POST("score", middlewares.Check, r.HomeController.RegOrLogin)
}
