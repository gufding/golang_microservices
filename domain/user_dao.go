package domain

import (
	"net/http"

	"github.com/gufding/golang_microservices/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "first", LastName: "last", Email: "email@gmail.com"},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	user := users[userId]

	// User not found
	if user == nil {
		return nil, &utils.ApplicationError{
			Message:    "User not found.\n",
			StatusCode: http.StatusNotFound,
			Code:       "not_found\n",
		}
	}

	// User found
	return user, nil
}
