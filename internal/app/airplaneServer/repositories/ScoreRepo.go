package repositories

import (
	"airplaneServer/internal/app/airplaneServer/database/mysql"
	"airplaneServer/internal/app/airplaneServer/models"
)

func NewScoreRepo(db *mysql.DB) *ScoreRepo {
	return &ScoreRepo{DB: db}
}

type ScoreRepo struct {
	DB mysql.IDB
}

func (u *ScoreRepo) Create(userID int64, score int64) error {
	var scoreRecord models.Score
	scoreRecord.UserID = userID
	scoreRecord.Score = score
	if error := u.DB.GetConnect().Create(&scoreRecord).Error; error != nil {
		return error
	}
	return nil
}
