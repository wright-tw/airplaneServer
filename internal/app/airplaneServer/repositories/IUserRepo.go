package repositories

import (
	"airplaneServer/internal/app/airplaneServer/models"
)

type IUserRepo interface {
	Find(string) (models.User, error)
	Create(string, string) error
}
