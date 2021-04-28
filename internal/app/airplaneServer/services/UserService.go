package services

import (
	"airplaneServer/internal/app/airplaneServer/database/redis"
	"airplaneServer/internal/app/airplaneServer/repositories"
	"airplaneServer/pkg/hash"
	// "airplaneServer/pkg/logger"
	"errors"
	"github.com/jinzhu/gorm"
	"strconv"
	"time"
)

func NewUserService(UserRepo *repositories.UserRepo, ScoreRepo *repositories.ScoreRepo) *UserService {
	return &UserService{
		UserRepo:  UserRepo,
		ScoreRepo: ScoreRepo,
	}
}

type UserService struct {
	UserRepo  repositories.IUserRepo
	ScoreRepo repositories.IScoreRepo
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

func (u *UserService) WriteScore(token string, score int64) error {
	// check token
	client := redis.GetRedisClient()
	userIDString, err := client.Get(redis.Ctx, token).Result()
	if err != nil && err.Error() != "redis: nil" {
		return err
	}
	if userIDString == "" {
		return errors.New("token error")
	}

	userID, err2 := strconv.ParseInt(userIDString, 10, 64)
	if err2 != nil {
		return err2
	}

	// write score
	err3 := u.ScoreRepo.Create(userID, score)
	if err3 != nil {
		return err3
	}

	return nil

}
