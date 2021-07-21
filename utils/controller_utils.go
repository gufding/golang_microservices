package utils

import "github.com/gin-gonic/gin"

func Respond(c *gin.Context, status int, body interface{}) {
  // If the response is requested in XML then respond with XML
  if c.GetHeader("Accept") == "application/xml" {
    c.XML(status, body)
    return
  }
  
  // Otherwise respond with JSON
  c.JSON(status, body)
}

func RespondError (c *gin.Context, err *ApplicationError) {
  // If the response is requested in XML then respond with XML
  if c.GetHeader("Accept") == "application/xml" {
    c.XML(err.StatusCode, err)
    return
  }
  
  // Otherwise respond with JSON
  c.JSON(err.StatusCode, err)
}