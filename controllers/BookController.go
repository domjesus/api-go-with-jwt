package controllers

import (
	database "domjesus/go-with-docker/db"
	"domjesus/go-with-docker/errors"
	"domjesus/go-with-docker/models"
	"domjesus/go-with-docker/repositories"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func BookCreate(w http.ResponseWriter, r *http.Request) {
	connection, _ := database.ConectaComBancoDeDados()

	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)

	if err != nil {
		var err errors.Error
		err = errors.SetError(err, "Error in reading payload.")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "Error in request", 400)
		return
	}

	if err := models.ValidaDadosDaRequest(&book); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	connection.Create(&book)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func ListAllBooks(w http.ResponseWriter, r *http.Request) {

	books := repositories.AllBooks()
	books_json, _ := json.Marshal(books)
	w.Write(books_json)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	// fmt.Println("ID: ", id)
	book := repositories.GetBookById(id)
	if book.ID == 0 {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	book_json, _ := json.Marshal(book)

	w.Write(book_json)
}
