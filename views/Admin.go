package views

import (
	"net/http"
)

func AdminIndex(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("Content of header: ", r.Header)

	if r.Header.Get("Role") != "admin" {
		w.Write([]byte("Not authorized. Not a admin."))
		return
	}
	w.Write([]byte("Welcome, Admin."))
}
