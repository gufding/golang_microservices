package app

import (
	"net/http"

	"github.com/gufding/golang_microservices/mvc/controllers"
)

func StartApp() {
	// Specify that the get users is the controller to be used
	http.HandleFunc("/users", controllers.GetUser)

	// Launch app using 8080 port
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
