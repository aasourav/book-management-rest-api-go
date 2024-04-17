package routes

import (
	"gorilla-test/controllers"

	"github.com/gorilla/mux"
)

func BookRoutes(router *mux.Router) {
	router.HandleFunc("/getBooks", controllers.GetBooks)
	router.HandleFunc("/getBook/{bookId}", controllers.GetBook)
	router.HandleFunc("/deleteBook/{bookId}", controllers.DeleteBook)
	router.HandleFunc("/addBook", controllers.AddBook)
}
