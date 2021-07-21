package app

import (
  "github.com/gufding/golang_microservices/mvc/controllers"
)

func mapURLs() {
  // Specify that the get users is the controller to be used
  router.GET("/users/:user_id", controllers.GetUser)
}
