package services

import (
	"github.com/gufding/golang_microservices/mvc/domain"
	"github.com/gufding/golang_microservices/mvc/utils"
)

type userService struct {}

var (
	UserService userService
)

func (u *userService) GetUser(UserId int64) (*domain.User, *utils.ApplicationError) {
	// Call the domain to get the user
	return domain.UserDao.GetUser(UserId)
}
