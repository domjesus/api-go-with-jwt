package database

import (
	"domjesus/go-with-docker/models"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() (*gorm.DB, error) {

	// stringDeConexao := os.Getenv("DATABASE_URL")
	err := godotenv.Load()
	if err != nil {
		// l.Error("Error loading .env file")
		fmt.Println("ERror loading .env file")
	}

	sslmode := " sslmode=disable"

	// env := os.Getenv("ENV")

	// if env == "local" {
	// sslmode = "sslmode=disable"
	// } else {
	// sslmode = "sslmode=require"
	// }

	var stringDeConexao string
	// fmt.Println("Ambiente: ", env)

	stringDeConexao = "host=" + os.Getenv("DATABASE_HOST") + " user=" + os.Getenv("DATABASE_USER") + " password=" + os.Getenv("DATABASE_PASSWORD") + " dbname=" + os.Getenv("DATABASE_NAME") + " port=" + os.Getenv("DATABASE_PORT") + sslmode

	// fmt.Println("STring de conexao: ", stringDeConexao)

	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		return nil, err
	}

	DB.AutoMigrate(&models.Users)

	fmt.Println("DB connected")
	// l.Info("DB connected")

	return DB, nil
}

// func Closedatabase(connection *gorm.DB) {
// logger, _ := zap.NewProduction()
// defer logger.Sync() // flushes buffer, if any
// sugar := logger.Sugar()

// sqldb := *gorm.DB
// sqldb.Close()
// }
