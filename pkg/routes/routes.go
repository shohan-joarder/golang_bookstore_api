package routes

import (
	"github.com/gorilla/mux"
	"github.com/shohan-joarder/go_bookstore/pkg/controllers"
)


var RegisterBookStore = func (router *mux.Router)  {
	router.HandleFunc("/books/",controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/",controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}",controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}",controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}",controllers.DeleteBook).Methods("DELETE")
}