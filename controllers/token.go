package controllers

import (
	"domjesus/go-with-docker/errors"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
)

func GetTokenInfo(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("Token: ", r.Header["Authorization"])

	var mySigningKey = []byte(os.Getenv("JWT_SECRET"))

	token, err := jwt.Parse(r.Header["Authorization"][0], func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error in parsing")
		}
		return mySigningKey, nil
	})

	if err != nil {
		var err errors.Error
		err = errors.SetError(err, "Your Token has been expired")
		// json.NewEncoder(w).Encode(err)
		http.Error(w, err.Message, 401)
		return
	}

	type DataUser struct {
		Name string
		Role string
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println("Token validated: ", claims["name"], claims["role"])
		// w.Write(claims["name"])
		name := fmt.Sprintf("%v", claims["name"])
		role := fmt.Sprintf("%v", claims["role"])

		umUser := DataUser{name, role}

		user_json, _ := json.Marshal(umUser)

		w.Write([]byte(user_json))

		// fmt.Println("Name: ", name.(string))
		// w.Write([]byte("Name: " + name + " - Role: " + role)) //ESCREVE NA SAÍDA
	}

}
