package repositories

import "airplaneServer/internal/app/airplaneServer/models"

type IUserRepo interface {
	Get() ([]models.User, error)
	Create(map[string]string) error
	Update(int, map[string]string) error
	Delete(int) error
}
