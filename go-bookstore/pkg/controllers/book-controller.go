package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/yansb/go-bookstore/pkg/models"
	"github.com/yansb/go-bookstore/pkg/utils"
)

var NewBook models.Book

func GetBooks(response http.ResponseWriter, request *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	response.Header().Set("Content-Type", "pkglication/json")
	response.WriteHeader(http.StatusOK)
	response.Write(res)
}

func GetBookById(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	response.Header().Set("Content-Type", "pkglication/json")
	response.WriteHeader(http.StatusOK)
	response.Write(res)

}

func CreateBook(response http.ResponseWriter, request *http.Request) {
	NewBook := &models.Book{}
	utils.ParseBody(request, NewBook)
	b := NewBook.CreateBook()
	res, _ := json.Marshal(b)
	response.Header().Set("Content-Type", "pkglication/json")
	response.WriteHeader(http.StatusCreated)
	response.Write(res)
}

func DeleteBook(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails := models.DeleteBook(ID)
	res, _ := json.Marshal(bookDetails)
	response.Header().Set("Content-Type", "pkglication/json")
	response.WriteHeader(http.StatusOK)
	response.Write(res)
}

func UpdateBook(response http.ResponseWriter, request *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(request, updateBook)
	vars := mux.Vars(request)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	booksDetails, db := models.GetBookById(ID)
	if updateBook.Name != "" {
		booksDetails.Name = updateBook.Name
	}
	if updateBook.Author != "" {
		booksDetails.Author = updateBook.Author
	}
	if updateBook.Publication != "" {
		booksDetails.Publication = updateBook.Publication
	}
	db.Save(&booksDetails)
	res, _ := json.Marshal(booksDetails)
	response.Header().Set("Content-Type", "pkglication/json")
	response.WriteHeader(http.StatusOK)
	response.Write(res)
}
