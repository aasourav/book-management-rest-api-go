package routes

import (
	"gorilla-test/controllers"

	"github.com/gorilla/mux"
)

func OrderRoutes(router *mux.Router) {
	router.HandleFunc("/orderBook/{bookId}", controllers.OrderBook)
	router.HandleFunc("/listAllOrderedBooks", controllers.ListAllOrderedBooks)
}
