package config

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDB() {
	// Railway menyediakan env DATABASE_URL secara otomatis
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		host := os.Getenv("PGHOST")
		port := os.Getenv("PGPORT")
		user := os.Getenv("PGUSER")
		pass := os.Getenv("PGPASSWORD")
		dbname := os.Getenv("PGDATABASE")

		if host != "" && dbname != "" {
			dbURL = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, pass, host, port, dbname)
		} else {
			dbURL = "postgres://postgres:password@localhost:5432/TokoBuku"
		}
	}

	var err error
	DB, err = pgxpool.New(context.Background(), dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Gagal koneksi ke database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Koneksi Database Berhasil!")
}