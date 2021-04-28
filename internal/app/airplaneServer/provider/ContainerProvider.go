package provider

import (
	"airplaneServer/internal/app/airplaneServer/controllers"
	"airplaneServer/internal/app/airplaneServer/database/mysql"
	"airplaneServer/internal/app/airplaneServer/models"
	"airplaneServer/internal/app/airplaneServer/repositories"
	"airplaneServer/internal/app/airplaneServer/routes"
	"airplaneServer/internal/app/airplaneServer/services"
	"go.uber.org/dig"
)

type ContainerProvider struct{}

var serviceList = []interface{}{

	// 控制器
	controllers.NewHomeController,

	// 服務
	services.NewUserService,

	// 資料組合
	repositories.NewUserRepo,
	repositories.NewScoreRepo,

	// 模型
	models.NewUser,
	models.NewScore,

	// 底層基本
	routes.NewRouter,
	mysql.NewDB,
}

func (c *ContainerProvider) GetInjectedRouter() *routes.Router {
	container := dig.New()

	for _, function := range serviceList {
		if provideErr := container.Provide(function); provideErr != nil {
			panic(provideErr)
		}
	}

	// 自動注入所有控制器
	router := &routes.Router{}
	invokeErr := container.Invoke(func(readyRouter *routes.Router) {
		router = readyRouter
	})
	if invokeErr != nil {
		panic(invokeErr)
	}

	return router
}
