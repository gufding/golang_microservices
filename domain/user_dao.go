package domain

import (
	"log"
	"net/http"
	"strconv"
	
	"github.com/gufding/golang_microservices/mvc/utils"
)

var (
	users = map[int64]*User{
		123: {Id: 123, FirstName: "first", LastName: "last", Email: "email@gmail.com"},
	}

	UserDao usersDaoInterface
)

func init () {
	UserDao = &userDao{}
}

type usersDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

type userDao struct{}

func (u *userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {
	log.Println("We're accessing the database.")

	user := users[userId]

	// User not found
	if user == nil {
		return nil, &utils.ApplicationError{
			Message:    "User " + strconv.FormatInt(userId, 10) + " not found.",
			StatusCode: http.StatusNotFound,
			Code:       "not_found",
		}
	}

	// User found
	return user, nil
}
