package views

import (
	"net/http"
)

func UserIndex(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("Content of header in User: ", r.Header)
	if r.Header.Get("Role") != "user" {
		w.Write([]byte("Not Authorized. Not a user."))
		return
	}
	w.Write([]byte("Welcome, User."))
}
