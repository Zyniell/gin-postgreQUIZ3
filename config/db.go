package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	// Mengambil variabel environment dari Railway
	// Pastikan PGDATABASE di Railway diisi: TokoBuku
	host := os.Getenv("PGHOST")
	port := os.Getenv("PGPORT")
	user := os.Getenv("PGUSER")
	password := os.Getenv("PGPASSWORD")
	dbname := os.Getenv("PGDATABASE")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Printf("Error opening database: %v\n", err)
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		fmt.Printf("Error connecting to database: %v\n", err)
		panic(err)
	}

	fmt.Printf("Successfully connected to database: %s\n", dbname)
}