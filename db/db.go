package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados(l *zap.SugaredLogger) (*gorm.DB, error) {
	// fmt.Println("Starting connect to DB")

	// stringDeConexao := os.Getenv("DATABASE_URL")
	godotenv.Load()
	// if err != nil {
	// 	// l.Error("Error loading .env file")
	// 	fmt.Println("Error loading .env file")
	// }

	// sslmode := " sslmode=disable"

	// if env == "local" {
	// sslmode = "sslmode=disable"
	// } else {
	// sslmode = "sslmode=require"
	// }

	// var stringDeConexao string
	// fmt.Println("Ambiente: ", env)

	// stringDeConexao = "host=" + os.Getenv("DATABASE_HOST") + " user=" + os.Getenv("DATABASE_USER") + " password=" + os.Getenv("DATABASE_PASSWORD") + " dbname=" + os.Getenv("DATABASE_NAME") + " port=" + os.Getenv("DATABASE_PORT") + sslmode

	stringDeConexao := os.Getenv("DATABASE_USER") + ":" + os.Getenv("DATABASE_PASSWORD") + "@tcp(" + os.Getenv("DATABASE_HOST") + ":3306)/" + os.Getenv("DATABASE_NAME") + "?charset=utf8mb4&parseTime=True&loc=Local"

	// fmt.Println("String de conexao: ", stringDeConexao)

	DB, err = gorm.Open(mysql.Open(stringDeConexao), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// fmt.Println("DB connected")

	if l != nil {
		l.Info("DB connected")
	}

	// DB.AutoMigrate(&models.Book{})
	// DB.AutoMigrate(&models.User{})
	// DB.AutoMigrate(&models.Location{})
	// DB.AutoMigrate(&models.Trash{})
	// , &models.Book{})
	// DB.Migrator().CreateTable(&models.Author{})
	// DB.Migrator().CreateTable(&models.Book{})

	// l.Info("DB connected")

	return DB, nil
}

func Closedatabase(connection *gorm.DB) {
	// logger, _ := zap.NewProduction()
	// defer logger.Sync() // flushes buffer, if any
	// sugar := logger.Sugar()

	conn, _ := connection.DB()
	fmt.Println("Closing connection...")
	conn.Close()
}
