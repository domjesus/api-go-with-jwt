package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	database "domjesus/go-with-docker/db"
	"domjesus/go-with-docker/routes"

	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", "URL-XPTO",
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Logger working fine ")

	// r := mux.NewRouter()
	// r.HandleFunc("/", HomeHandler)
	//RELATION BELONGS TO
	// var book []models.Book
	// var book models.Book
	// book.ID = 3
	// book.AuthorID = 3
	// var author models.Author
	// book := models.Book{
	// 	ISBN:  "987654321",
	// 	Title: "Titulo do livro 4",
	// 	Author: models.Author{
	// 		Name:   "Nome do autor do livro 9",
	// 		Age:    "59",
	// 		Gender: "Male",
	// 	},
	// 	AuthorID: 4,
	// }

	connection, _ := database.ConectaComBancoDeDados(sugar)
	defer database.Closedatabase(connection)

	// connection.Preload("Author").Find(&book) //GET THE MODEL AND RELATION
	// connection.Where("author_id = ?", 4).Preload("Author").Find(&book) //GET THE MODEL AND RELATION BY THE FK
	// connection.Create(&book)
	// connection.Save(&book)

	// fmt.Println("Livro: ", book)
	// return

	routes.CreateRouter()
	routes.InitializeRoute(sugar)

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

// func HomeHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Rota Home")
// 	w.Write([]byte(os.Getenv("JWT_SECRET")))
// }

func ServerStart() {
	var port string
	godotenv.Load()

	port = os.Getenv("PORT")

	if port == "" {
		port = "80"
	}

	fmt.Println("Server started at http://localhost:" + port)

	err := http.ListenAndServe(":"+port, handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Access-Control-Allow-Origin", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(routes.Router))
	if err != nil {
		log.Fatal(err)
	}
}
