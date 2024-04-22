package routes

import (
	"gorilla-test/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func UserRoutes(router *mux.Router) {
	router.HandleFunc("/login", controllers.Login).Methods(http.MethodPost)
	router.HandleFunc("/signup", controllers.SignUp).Methods(http.MethodPost)
}
