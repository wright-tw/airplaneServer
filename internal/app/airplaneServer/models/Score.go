package models

type Score struct {
	ID     int64 `json:"id"  gorm:"autoIncrement:true;primaryKey"`
	UserID int64 `json:"user_id" gorm:"type:int;index"`
	Score  int64 `json:"score" gorm:"type:int"`
}

func NewScore() *Score {
	return &Score{}
}

func (u *Score) TableName() string {
	return "scores"
}
