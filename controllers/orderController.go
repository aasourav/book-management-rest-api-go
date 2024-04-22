package controllers

import (
	"encoding/json"
	"gorilla-test/data"
	datatypes "gorilla-test/dataTypes"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func OrderBook(w http.ResponseWriter, r *http.Request) {
	username := r.Header.Get("username")

	vars := mux.Vars(r)
	bookIdStr := vars["bookId"]

	// if err {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte("book id not proided in url"))
	// 	return
	// }

	bookId, err := strconv.Atoi(bookIdStr)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request"))
		return
	}

	_, exist := data.Books[bookId]
	if !exist {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Book not in the store"))
		return
	}

	data.Orders[username] = append(data.Orders[username], bookId)
	w.Write([]byte("book ordered Successfully"))

}

func ListAllOrderedBooks(w http.ResponseWriter, r *http.Request) {
	username := r.Header.Get("username")

	var listOfBooks datatypes.Books

	orders := data.Orders[username]

	for _, bookId := range orders {
		book := datatypes.Book{
			Id:   int(bookId),
			Name: data.Books[int(bookId)],
		}
		listOfBooks.AddBooksToList(book)
	}

	marsheledData, err := json.Marshal(listOfBooks)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	w.Write(marsheledData)
}
