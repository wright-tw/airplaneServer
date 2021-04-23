package models

type User struct {
	ID       int64  `json:"id"  gorm:"autoIncrement:true;primaryKey"`
	Username string `json:"username" gorm:"type:varchar(32);unique"`
	Password string `json:"password" gorm:"type:varchar(100);DEFAULT:null"`
}

func NewUser() *User {
	return &User{}
}

func (u *User) GetName(name string) string {
	return u.Username
}

func (u *User) TableName() string {
	return "users"
}
