package controllers

import (
	database "domjesus/go-with-docker/db"
	"domjesus/go-with-docker/errors"
	"domjesus/go-with-docker/models"
	"domjesus/go-with-docker/repositories"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateOccurrence(w http.ResponseWriter, r *http.Request) {
	connection, _ := database.ConectaComBancoDeDados(nil)
	defer database.Closedatabase(connection)

	var occurrence models.Occurrence

	err := json.NewDecoder(r.Body).Decode(&occurrence)
	if err != nil {
		var err errors.Error
		err = errors.SetError(err, "Error in reading body")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	connection.Create(&occurrence)

	w.Header().Set("Content-Type", "aplication/json")

	json.NewEncoder(w).Encode(occurrence)
	return

}

type occurrence struct {
	Name string `validate:"nonzero" json:"name"`
	Type string `validate:"nonzero" json:"type"`
	Icon string `validate:"nonzero" json:"icon"`
}

func GetAllOccurrences(w http.ResponseWriter, r *http.Request) {

	// isAdmin := logedUserIsAdmin(r.Header["Authorization"][0])

	// if isAdmin {
	occurrences := repositories.AllOccurrences()

	occurrences_json, err := json.Marshal(occurrences)

	if err != nil {
		fmt.Println("Erro recuperando 'occurrences'.")
	}

	w.Write(occurrences_json)
	return

}
