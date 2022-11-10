package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func ConnectDatabase() (engine *gin.Engine, db *sql.DB) {
	// cargar el env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error: Loading .env")
	}

	// configurar db
	configDB := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: os.Getenv("DBNAME"),
	}

	// abrir la db
	db, err = sql.Open("mysql", configDB.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	// retornar el engine
	engine = gin.Default()

	return engine, db
}