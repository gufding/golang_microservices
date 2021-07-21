package services

import (
	"net/http"

	"github.com/gufding/golang_microservices/mvc/domain"
	"github.com/gufding/golang_microservices/mvc/utils"
)

type itemService struct {
}

var (
	ItemService itemService
)

func (s *itemService) GetItem(itemId string) (*domain.Item, *utils.ApplicationError) {
	return nil, &utils.ApplicationError{
		Message:    "implement me",
		StatusCode: http.StatusInternalServerError,
		Code:       "test",
	}
}
