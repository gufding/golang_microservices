package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	
	"github.com/gufding/golang_microservices/mvc/services"
	"github.com/gufding/golang_microservices/mvc/utils"
)

func GetUser(c *gin.Context) {
	// Check that the user ID is a number
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		apiErr := &utils.ApplicationError{
			Message:    "User ID must be a number.\n",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		utils.RespondError(c, apiErr)
		return
	}

	// Get the user using the passed user ID
	user, apiErr := services.UserService.GetUser(userId)
	if apiErr != nil {
		utils.RespondError(c, apiErr)
		return
	}

	// Return user to client
	utils.Respond(c, http.StatusOK, user)
}
