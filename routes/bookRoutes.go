package routes

import (
	"gorilla-test/controllers"
	"gorilla-test/middleware"

	"github.com/gorilla/mux"
)

func BookRoutes(router *mux.Router) {
	router.HandleFunc("/getBooks", middleware.ValidateUser(controllers.GetBooks))
	router.HandleFunc("/getBook/{bookId}", controllers.GetBook)
	router.HandleFunc("/deleteBook/{bookId}", controllers.DeleteBook)
	router.HandleFunc("/addBook", controllers.AddBook)
}
