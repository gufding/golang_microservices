package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gufding/golang_microservices/mvc/services"
	"github.com/gufding/golang_microservices/mvc/utils"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {
	// Check that the user ID is a number
	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "User ID must be a number.\n",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}

	// Get the user using the passed user ID
	user, apiErr := services.GetUser(userId)
	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		resp.WriteHeader(apiErr.StatusCode)
		resp.Write(jsonValue)
		return
	}

	// Return user to client
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}
