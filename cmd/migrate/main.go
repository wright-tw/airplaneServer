package main

import (
	"airplaneServer/configs"
	"airplaneServer/internal/app/airplaneServer/database/mysql"
	"airplaneServer/internal/app/airplaneServer/models"
)

func main() {
	configs.Setting()
	db := mysql.NewDB()
	db.Connect.AutoMigrate(
		// 把需要建立表的模型全部丟進來
		&models.User{},
	)
}
