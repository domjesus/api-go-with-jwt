package controllers

import (
	database "domjesus/go-with-docker/db"
	"domjesus/go-with-docker/errors"
	"domjesus/go-with-docker/models"
	"domjesus/go-with-docker/repositories"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
)

func Create(w http.ResponseWriter, r *http.Request) {
	connection, _ := database.ConectaComBancoDeDados()
	defer database.Closedatabase(connection)

	// resBody, err1 := ioutil.ReadAll(r.Body)
	// if err1 != nil {
	// 	fmt.Printf("client: could not read response body: %s\n", err1)
	// 	os.Exit(1)
	// }
	// fmt.Printf("client: response body: %s\n", resBody)

	var location models.Location

	err := json.NewDecoder(r.Body).Decode(&location)
	if err != nil {
		var err errors.Error
		err = errors.SetError(err, "Error in reading body")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}
	// fmt.Print(location)

	connection.Create(&location)

	w.Header().Set("Content-Type", "aplication/json")

	json.NewEncoder(w).Encode(location)
	return

}

func GetAll(w http.ResponseWriter, r *http.Request) {

	locations := repositories.AllLocations()
	locations_json, _ := json.Marshal(locations)
	w.Write(locations_json)

}

func GetLocationById(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	location := repositories.GetLocationById(id)

	if location.ID == 0 {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	locations_json, _ := json.Marshal(location)

	w.Write(locations_json)

}

func GetMyLocations(w http.ResponseWriter, r *http.Request) {
	// var userId float64

	token, _ := jwt.Parse(r.Header["Authorization"][0], func(token *jwt.Token) (interface{}, error) {
		// if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		// return nil, fmt.Errorf("There was an error in parsing")
		// }
		// fmt.Print("Id do user:s", token)
		// return token, nil

		return nil, nil
	})

	claims, _ := token.Claims.(jwt.MapClaims)
	// var userId uint
	// fmt.Print("Claims: ", claims)

	// fmt.Printf("Type of %T\n", claims["id"])

	// if ok && token.Valid {
	userId := claims["id"]

	// }

	fmt.Print("Id do user:", userId)

	my_locations := repositories.GetMyLocations(userId.(float64))
	// print("Locations: ", my_locations)

	locations_json, _ := json.Marshal(my_locations)

	w.Write(locations_json)

}
