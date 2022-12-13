package controllers

import (
	"github.com/golang-jwt/jwt"
)

func logedUserIsAdmin(tokenString string) bool {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return nil, nil
	})

	claims, _ := token.Claims.(jwt.MapClaims)

	role := claims["role"].(string)

	// fmt.Println("userId: ", userId, "role: ", role)

	return role == "admin"

}
