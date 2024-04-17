package main

import (
	"gorilla-test/routes"
	"net/http"

	"github.com/gorilla/mux"
)

/*
user management system
login
signup
order
see the list of books he ordered
see the list of books he order3ed
owner should able to add new books
owner should able to deletre books
*/
func main() {
	router := mux.NewRouter()
	routes.BookRoutes(router)
	routes.OrderRoutes(router)
	routes.UserRoutes(router)

	http.ListenAndServe(":8000", router)
}

// package main

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// func main() {
// 	fmt.Println("Hello world")
// 	router := mux.NewRouter()

// 	router.Use(globalMiddleware)

// 	router.HandleFunc("/home", sampleMiddleware(Home))
// 	http.ListenAndServe(":8000", router)
// }

// func Home(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Home"))
// }

// func sampleMiddleware(h http.HandlerFunc) http.HandlerFunc {

// 	return func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("this is middleware")
// 		h.ServeHTTP(w, r)
// 	}
// }

// func globalMiddleware(h http.Handler) http.Handler {
// 	fmt.Println("this is global middleware")
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		h.ServeHTTP(w, r)
// 	})
// }
