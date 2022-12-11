package controllers

import (
	database "domjesus/go-with-docker/db"
	"domjesus/go-with-docker/errors"
	"domjesus/go-with-docker/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/validator.v2"
)

type Authentication struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	Id          uint   `json:"id"`
	Role        string `json:"role"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	TokenString string `json:"token"`
}

func SignIn(w http.ResponseWriter, r *http.Request) {

	connection, _ := database.ConectaComBancoDeDados()
	// defer Closedatabase(connection)
	defer database.Closedatabase(connection)

	var authdetails Authentication
	err := json.NewDecoder(r.Body).Decode(&authdetails)
	if err != nil {
		var err errors.Error
		err = errors.SetError(err, "Error in reading body")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	var authuser models.User
	connection.Where("email = ?", authdetails.Email).First(&authuser)
	if authuser.Email == "" {
		var err errors.Error
		err = errors.SetError(err, "Username or Password is incorrect")
		w.Header().Set("Content-Type", "application/json")
		// json.NewEncoder(w).Encode(err)
		http.Error(w, "Invalid credentials", 401)
		return
	}

	check := CheckPasswordHash(authdetails.Password, authuser.Password)

	if !check {
		var err errors.Error
		err = errors.SetError(err, "Username or Password is incorrect")
		w.Header().Set("Content-Type", "application/json")
		// json.NewEncoder(w).Encode(err)
		http.Error(w, "Invalid credentials", 401)
		return
	}

	validToken, err := GenerateJWT(authuser.ID, authuser.Email, authuser.Role, authuser.Name)
	if err != nil {
		var err errors.Error
		err = errors.SetError(err, "Failed to generate token")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	var token Token
	token.Id = authuser.ID
	token.Email = authuser.Email
	token.Name = authuser.Name
	token.Role = authuser.Role
	token.TokenString = validToken
	w.Header().Set("Content-Type", "application/json")
	// database.Closedatabase(connection)
	json.NewEncoder(w).Encode(token)
	return
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateJWT(id uint, email, role string, name string) (string, error) {
	secretkey := os.Getenv("JWT_SECRET")

	var mySigningKey = []byte(secretkey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["id"] = id
	claims["email"] = email
	claims["name"] = name
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	// fmt.Println("Token data: ", claims)

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	connection, _ := database.ConectaComBancoDeDados()

	// defer CloseDatabase(connection)
	defer database.Closedatabase(connection)

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		var err errors.Error
		err = errors.SetError(err, "Error in reading payload. Possible malformed request body")
		w.Header().Set("Content-Type", "application/json")
		// json.NewEncoder(w).Encode(err)

		http.Error(w, err.Message, 400)
		return
	}

	// if user.password != user.password_confirm {
	// 	var err errors.Error
	// 	err = errors.SetError(err, "Passwords not matches.")
	// 	w.Header().Set("Content-Type", "application/json")

	// 	http.Error(w, "Error in reading payload.", 403)
	// 	return
	// }

	if err := models.ValidaUser(&user); err != nil {
		http.Error(w, err.Error(), 422)
		return
	}

	if err := models.ValidaEmail(user.Email); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	validator.SetValidationFunc("passwords_matches", models.ValidaPasswords)
	if err := models.ValidaPasswords(user.Password, user.Password_confirm); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var dbuser models.User
	connection.Where("email = ?", user.Email).First(&dbuser)

	//check email is alredy registered or not
	if dbuser.Email != "" {
		var err errors.Error
		err = errors.SetError(err, "Email already in use")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, "Email already in use", 400)

		// json.NewEncoder(w).Encode(err)
		return
	}

	user.Password, err = GeneratehashPassword(user.Password)
	if err != nil {
		fmt.Println("Error in password hashing.")
		os.Exit(0)
	}

	//insert user details indatabase
	connection.Create(&user)
	w.Header().Set("Content-Type", "aplication/json")
	user.Password_confirm = ""

	database.Closedatabase(connection)
	json.NewEncoder(w).Encode(user)
	return
}

func GeneratehashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
