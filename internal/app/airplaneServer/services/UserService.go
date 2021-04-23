package services

import (
	"airplaneServer/internal/app/airplaneServer/database/redis"
	"airplaneServer/internal/app/airplaneServer/repositories"
	"airplaneServer/pkg/hash"
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

func NewUserService(UserRepo *repositories.UserRepo) *UserService {
	return &UserService{UserRepo: UserRepo}
}

type UserService struct {
	UserRepo repositories.IUserRepo
}

func (u *UserService) RegOrLogin(username string, password string) (string, error) {
	user, err := u.UserRepo.Find(username)
	password = hash.Md5(password)
	now := time.Now()
	token := ""

	// DB error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", err
	}

	// no username
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// reg
		err := u.UserRepo.Create(username, password)
		if err != nil {
			return "", err
		}

		newuser, err := u.UserRepo.Find(username)
		user = newuser
		if err != nil {
			return "", err
		}
	}

	// check password
	if user.Password != password {
		err := errors.New("password error")
		return "", err
	}

	token = hash.Md5(user.Username + now.Format("20060102150405"))
	client := redis.GetRedisClient()
	client.Set(redis.Ctx, token, user.ID, 1800*time.Second)

	// create token
	return token, nil

}
