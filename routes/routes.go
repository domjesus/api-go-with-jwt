package routes

import (
	"domjesus/go-with-docker/controllers"
	login "domjesus/go-with-docker/controllers"
	"domjesus/go-with-docker/middlewares"
	"domjesus/go-with-docker/views"
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
	Router.HandleFunc("/admin", middlewares.IsAuthorized(views.AdminIndex)).Methods("GET")
	Router.HandleFunc("/user", middlewares.IsAuthorized(views.UserIndex)).Methods("GET")
	Router.HandleFunc("/book", controllers.BookCreate).Methods("POST")
	Router.HandleFunc("/books", middlewares.IsAuthorized(controllers.ListAllBooks)).Methods("GET")
	Router.Path("/books/{id}").HandlerFunc(middlewares.IsAuthorized(controllers.GetBookById)).Methods("GET")
	Router.HandleFunc("/", Index).Methods("GET")
	Router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
	})
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HOME PUBLIC INDEX PAGE"))
}
