package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"domjesus/go-with-docker/routes"

	"github.com/gorilla/handlers"
)

func main() {
	// r := mux.NewRouter()
	// r.HandleFunc("/", HomeHandler)
	routes.CreateRouter()
	routes.InitializeRoute()
	// connection, _ := database.ConectaComBancoDeDados()

	ServerStart()
	// r.HandleFunc("/home", verifyJWT(handlePage))

	// logger, _ := zap.NewProduction()
	// defer logger.Sync() // flushes buffer, if any
	// sugar := logger.Sugar()

	// if err := db.ConectaComBancoDeDados(sugar); err != nil {
	// fmt.Printf("Error: %s", err)
	// }

	// port := ":8000"
	// fmt.Println("Server run at port ", port)
	// http.ListenAndServe(port, r)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Rota Home")
	w.Write([]byte(os.Getenv("JWT_SECRET")))
}

func ServerStart() {
	fmt.Println("Server started at http://localhost:8080")

	err := http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Access-Control-Allow-Origin", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(routes.Router))
	if err != nil {
		log.Fatal(err)
	}
}
