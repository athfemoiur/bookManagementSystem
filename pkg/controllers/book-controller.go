package controllers

import (
	"encoding/json"
	"github.com/athfemoiur/bookManagementSystem/pkg/models"
	"github.com/athfemoiur/bookManagementSystem/pkg/utils"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		log.Fatal(err)
	}
	book, _ := models.GetBookById(id)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var newBook = &models.Book{}
	utils.ParseBody(r, newBook)
	b := newBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		log.Fatal(err)
	}
	models.DeleteBook(id)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusNoContent)

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var updatedBook = &models.Book{}
	utils.ParseBody(r, updatedBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		log.Fatal(err)
	}
	book, db := models.GetBookById(id)
	book.Name = updatedBook.Name
	book.Author = updatedBook.Author
	book.Publication = updatedBook.Publication
	db.Save(&book)
	res, _ := json.Marshal(book)
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}
