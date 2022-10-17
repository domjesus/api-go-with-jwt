package routes

import (
	login "domjesus/go-with-docker/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	Router    *mux.Router
	secretkey string = "secretkeyjwt"
)

func CreateRouter() {
	Router = mux.NewRouter()
}

func InitializeRoute() {

	Router.HandleFunc("/signup", login.SignUp).Methods("POST")
	Router.HandleFunc("/signin", login.SignIn).Methods("POST")
	// router.HandleFunc("/admin", IsAuthorized(AdminIndex)).Methods("GET")
	// router.HandleFunc("/user", IsAuthorized(UserIndex)).Methods("GET")
	Router.HandleFunc("/", Index).Methods("GET")
	Router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
	})
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HOME PUBLIC INDEX PAGE"))
}
