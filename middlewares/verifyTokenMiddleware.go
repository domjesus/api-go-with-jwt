package middlewares

import (
	"domjesus/go-with-docker/errors"
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
)

func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// fmt.Println("Header: ", r.Header)

		if r.Header["Authorization"] == nil {
			var err errors.Error
			err = errors.SetError(err, "No Token Found")
			// json.NewEncoder(w).Encode(err)
			http.Error(w, "No token Found", 401)
			return
		}

		var mySigningKey = []byte(os.Getenv("JWT_SECRET"))

		token, err := jwt.Parse(r.Header["Authorization"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error in parsing")
			}
			return mySigningKey, nil
		})

		// fmt.Println("Erro no token: ", err)

		if err != nil {
			var err errors.Error
			err = errors.SetError(err, "Your Token has been expired")
			// json.NewEncoder(w).Encode(err)
			http.Error(w, err.Message, 401)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// fmt.Println("Token validated: ", claims)
			if claims["role"] == "admin" {
				r.Header.Set("Role", "admin")
				handler.ServeHTTP(w, r)
				// fmt.Println("Token validated 'admin': ", claims, r.Header)
				return
				// || claims["role"] == ""
			} else if claims["role"] == "user" {

				// fmt.Println("Token validated 'user': ", claims, r.Header)
				r.Header.Set("Role", "user")
				handler.ServeHTTP(w, r)
				return
			}
		}

		var reserr errors.Error
		reserr = errors.SetError(reserr, "Not Authorized. Not Rules.")
		http.Error(w, "Forbiden", http.StatusForbidden)
		// json.NewEncoder(w).Encode(err)
	}
}
