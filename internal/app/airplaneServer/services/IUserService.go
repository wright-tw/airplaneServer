package services

type IUserService interface {
	RegOrLogin(username string, password string) (string, error)
}
