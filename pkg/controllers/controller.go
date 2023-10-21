package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/saeedmzr/go-simple_book_store/pkg/models"
	"github.com/saeedmzr/go-simple_book_store/pkg/utils"
	"net/http"
	"strconv"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	Books := models.GetAllBooks()
	res, _ := json.Marshal(Books)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(res)
	if err != nil {
		return
	}

}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	BookId := vars["bookId"]
	ID, err := strconv.ParseInt(BookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing.")
	}

	Book, _ := models.GetBookById(ID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(Book)
	_, err = w.Write(res)
	if err != nil {
		return
	}
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	BookModel := &models.Book{}
	utils.ParseBody(r, BookModel)
	b := BookModel.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(res)
	if err != nil {
		return
	}
}
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	BookId := vars["bookId"]
	ID, err := strconv.ParseInt(BookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing.")
	}

	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(res)
	if err != nil {
		return
	}
}
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var UpdateBook = &models.Book{}
	utils.ParseBody(r, UpdateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("err while parsing.")
	}
	Book, db := models.GetBookById(ID)
	if UpdateBook.Name != "" {
		Book.Name = UpdateBook.Name
	}
	if UpdateBook.Name != "" {
		Book.Author = UpdateBook.Author
	}
	if UpdateBook.Publication != "" {
		Book.Publication = UpdateBook.Publication
	}
	db.Save(&Book)
	res, _ := json.Marshal(Book)
	w.Header().Set("Content-Type", "Application/json")
	_, err = w.Write(res)
	if err != nil {
		return
	}
}
