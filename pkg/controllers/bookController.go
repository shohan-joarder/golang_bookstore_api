package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/shohan-joarder/go_bookstore/pkg/models"
	"github.com/shohan-joarder/go_bookstore/pkg/utils"
)

var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request)  {
	newBook := models.GetAllBooks()
	res, _ := json.Marshal(newBook)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request){

	vars := mux.Vars(r)
	bookId := vars["bookId"]

	ID, err := strconv.ParseInt(bookId,0,0)
	if err !=nil {
		fmt.Println("Error while parsing")
	}
	bookDetails, _ :=models.GetBookById(ID)

	res, _ :=json.Marshal(bookDetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func CreateBook(w http.ResponseWriter, r *http.Request )  {
	fmt.Println(r);
	createBook := &models.Book{}
	utils.ParseBody(r,createBook)
	b := createBook.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	
	Id, err := strconv.ParseInt(bookId,0,0)

	if err !=nil{
		fmt.Println("Error while parsing")
	}

	book := models.DeleteBookById(Id)

	res, _ :=json.Marshal(book)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateBook(w http.ResponseWriter, r *http.Request)  {
	updateBook := &models.Book{}
	utils.ParseBody(r, updateBook)

	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId,0,0)

	if err !=nil{
		fmt.Println("Error while parsing")
	}

	 bookDetails, db := models.GetBookById(Id)

	 if updateBook.Name != ""{
		bookDetails.Name = updateBook.Name
	 }

	 if updateBook.Author != ""{
		bookDetails.Author = updateBook.Author
	 }

	 if updateBook.Publication != ""{
		bookDetails.Publication = updateBook.Publication
	 }

	 db.Save(&bookDetails)

	 res, _:= json.Marshal(bookDetails)
	 w.Header().Set("Content-Type","pkglication/json")
	 w.WriteHeader(http.StatusOK)
	 w.Write(res)

}