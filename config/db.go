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
		// Fallback untuk lokal jika lupa set env
		dbURL = "postgres://postgres:password@localhost:5432/TokoBuku"
	}

	var err error
	DB, err = pgxpool.New(context.Background(), dbURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Gagal koneksi ke database: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Koneksi Database Berhasil!")
}