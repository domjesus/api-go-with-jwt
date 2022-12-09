package repositories

import (
	database "domjesus/go-with-docker/db"
	"domjesus/go-with-docker/models"
)

func AllBooks() []models.Book {

	var books []models.Book

	connection, _ := database.ConectaComBancoDeDados()
	defer database.Closedatabase(connection)
	// connection.Where("id > ? ", 2).Find(&books)
	connection.Find(&books)
	return books
}

func GetBookById(id int) models.Book {
	var book models.Book

	connection, _ := database.ConectaComBancoDeDados()
	defer database.Closedatabase(connection)
	connection.First(&book, id)
	return book
}
