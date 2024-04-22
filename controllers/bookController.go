package controllers

import (
	"encoding/json"
	"gorilla-test/data"
	datatypes "gorilla-test/dataTypes"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookIdStr := vars["bookId"]
	bookId, err := strconv.Atoi(bookIdStr)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Bad request"))
		return
	}

	requestedBook, bookExists := data.Books[bookId]
	if !bookExists {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("book not found"))
		return
	}

	jsonEncodedResponse, err := json.Marshal(requestedBook)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Bad request"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	w.Write(jsonEncodedResponse)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	listOfBooks := data.Books
	var booksInArray datatypes.Books

	for bookId, book := range listOfBooks {
		book := datatypes.Book{
			Id:   bookId,
			Name: book,
		}
		booksInArray.AddBooksToList(book)
	}

	jsonEncodedResponse, err := json.Marshal(booksInArray)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Bad request"))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	w.Write(jsonEncodedResponse)

}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookIdStr := vars["bookId"]
	bookId, err := strconv.Atoi(bookIdStr)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Bad request"))
		return
	}

	_, bookExists := data.Books[bookId]
	if !bookExists {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("book not found"))
		return
	}

	delete(data.Books, bookId)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("book deleted"))
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	var newBook datatypes.Book
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad request"))
	}

	_, bookExists := data.Books[newBook.Id]
	if !bookExists {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("book not found"))
		return
	}

	data.Books[len(data.Books)+1] = newBook.Name

	w.Write([]byte("book added successfully"))

}
