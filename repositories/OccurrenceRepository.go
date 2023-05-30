package repositories

import (
	database "domjesus/go-with-docker/db"
	"domjesus/go-with-docker/models"
)

func AllOccurrences() []models.Occurrence {

	var occurrences []models.Occurrence

	connection, _ := database.ConectaComBancoDeDados(nil)
	defer database.Closedatabase(connection)
	// connection.Where("id > ? ", 2).Find(&books)
	connection.Find(&occurrences)
	return occurrences
}
