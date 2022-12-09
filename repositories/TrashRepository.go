package repositories

import (
	database "domjesus/go-with-docker/db"
	"domjesus/go-with-docker/models"
)

func AllTrashes() []models.Trash {

	var trashes []models.Trash

	connection, _ := database.ConectaComBancoDeDados()
	defer database.Closedatabase(connection)
	// connection.Where("id > ? ", 2).Find(&books)
	connection.Find(&trashes)
	return trashes
}
