package repositories

import (
	"airplaneServer/internal/app/airplaneServer/database/mysql"
	"airplaneServer/internal/app/airplaneServer/models"
)

func NewUserRepo(db *mysql.DB) *UserRepo {
	return &UserRepo{DB: db}
}

type UserRepo struct {
	DB mysql.IDB
}

func (u *UserRepo) Create(username string, password string) error {
	var user models.User
	user.Username = username
	user.Password = password
	if error := u.DB.GetConnect().Create(&user).Error; error != nil {
		return error
	}
	return nil
}

func (u *UserRepo) Find(username string) (models.User, error) {
	var User models.User
	result := u.DB.GetConnect().
		Where("username = ?", username).
		First(&User)
	err := result.Error
	if err != nil {
		return User, err
	}
	return User, nil
}
