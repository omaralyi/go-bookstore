package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/omaralyi/go-bookstore/pkg/models"
	"github.com/omaralyi/go-bookstore/pkg/utils"
	"net/http"
	"strconv"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	response, _ := json.Marshal(newBooks)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookID, err := strconv.ParseInt(params["BookId"], 0, 0)
	if err != nil {
		fmt.Println("error while parsing input")
		panic(err)
	}
	book, _ := models.GetBookById(bookID)
	response, _ := json.Marshal(book)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
