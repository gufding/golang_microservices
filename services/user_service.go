package services

import (
	"github.com/gufding/golang_microservices/mvc/domain"
	"github.com/gufding/golang_microservices/mvc/utils"
)

func GetUser(UserId int64) (*domain.User, *utils.ApplicationError) {
	return domain.GetUser(UserId)
}
