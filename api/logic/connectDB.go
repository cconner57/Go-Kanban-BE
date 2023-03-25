package connectDB

import (

	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	_ "github.com/lib/pq"
)

func ConnectDB() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dotenvHost := os.Getenv("DB_HOST")
	dotenvUser := os.Getenv("DB_USER")
	dotenvPassword := os.Getenv("DB_PASSWORD")
	dotenvDBName := os.Getenv("DB_NAME")

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require", dotenvHost, 5432, dotenvUser, dotenvPassword, dotenvDBName)

    db, err := sql.Open("postgres", psqlconn)
    if err != nil {
        panic(err)
    }
    defer db.Close()

    err = db.Ping()
	if err != nil {
        panic(err)
    }

    fmt.Println("Connected!")
}