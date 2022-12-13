package repositories

import (
	database "domjesus/go-with-docker/db"
	"domjesus/go-with-docker/models"
)

func AllLocations() []models.Location {

	var locations []models.Location

	connection, _ := database.ConectaComBancoDeDados(nil)
	defer database.Closedatabase(connection)
	// connection.Where("id > ? ", 2).Find(&books)
	connection.Preload("User").Preload("Trash").Find(&locations)
	return locations
}

func GetLocationById(id int) models.Location {
	var location models.Location

	connection, _ := database.ConectaComBancoDeDados(nil)
	defer database.Closedatabase(connection)
	connection.Omit("Password", "Password_confirm").Preload("User").Preload("Trash").First(&location, id)
	return location
}

func GetMyLocations(userId float64) []models.Location {

	var locations []models.Location

	connection, _ := database.ConectaComBancoDeDados(nil)
	defer database.Closedatabase(connection)
	connection.Preload("Trash").Where("user_id = ?", userId).Find(&locations)
	return locations
}
