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

func TrashCreate(w http.ResponseWriter, r *http.Request) {
	connection, _ := database.ConectaComBancoDeDados(nil)
	defer database.Closedatabase(connection)

	var trash models.Trash

	err := json.NewDecoder(r.Body).Decode(&trash)
	if err != nil {
		var err errors.Error
		err = errors.SetError(err, "Error in reading body")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	fmt.Print(trash)

	connection.Create(&trash)

	w.Header().Set("Content-Type", "aplication/json")

	json.NewEncoder(w).Encode(trash)
	return

}

func TrashGetAll(w http.ResponseWriter, r *http.Request) {

	trashes := repositories.AllTrashes()
	trashes_json, _ := json.Marshal(trashes)
	w.Write(trashes_json)

}
