package routes

import (
	"domjesus/go-with-docker/controllers"
	location "domjesus/go-with-docker/controllers"
	login "domjesus/go-with-docker/controllers"
	occurrence "domjesus/go-with-docker/controllers"
	trash "domjesus/go-with-docker/controllers"
	"domjesus/go-with-docker/middlewares"
	"domjesus/go-with-docker/views"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

var (
	Router    *mux.Router
	secretkey string = "secretkeyjwt"
)

func CreateRouter() {
	Router = mux.NewRouter()
}

func InitializeRoute(l *zap.SugaredLogger) {

	Router.HandleFunc("/signup", login.SignUp).Methods("POST")
	Router.HandleFunc("/signin", login.SignIn).Methods("POST")

	Router.HandleFunc("/location", middlewares.IsAuthorized(location.Create)).Methods("POST")
	Router.HandleFunc("/locations", middlewares.IsAuthorized(location.GetAll)).Methods("GET")
	Router.HandleFunc("/occurrences", middlewares.IsAuthorized(occurrence.GetAllOccurrences)).Methods("GET")
	Router.Path("/location/{id}").HandlerFunc(middlewares.IsAuthorized(location.GetLocationById)).Methods("GET")
	Router.HandleFunc("/my_locations", middlewares.IsAuthorized(location.GetMyLocations)).Methods("GET")

	Router.HandleFunc("/trash", middlewares.IsAuthorized(trash.TrashCreate)).Methods("POST")
	Router.HandleFunc("/trash", trash.TrashGetAll).Methods("GET")

	Router.HandleFunc("/admin", middlewares.IsAuthorized(views.AdminIndex)).Methods("GET")
	Router.HandleFunc("/user", middlewares.IsAuthorized(views.UserIndex)).Methods("GET")
	Router.HandleFunc("/book", controllers.BookCreate).Methods("POST")
	Router.HandleFunc("/books_mobile", controllers.ListAllBooks).Methods("GET")
	Router.HandleFunc("/books", middlewares.IsAuthorized(controllers.ListAllBooks)).Methods("GET")
	Router.Path("/books/{id}").HandlerFunc(middlewares.IsAuthorized(controllers.GetBookById)).Methods("GET")
	Router.HandleFunc("/", Index).Methods("GET")
	Router.HandleFunc("/token", controllers.GetTokenInfo).Methods("GET")
	Router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
	})

	l.Info("Routes done")
}

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("HOME PUBLIC INDEX PAGE"))
}
