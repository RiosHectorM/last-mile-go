package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func GetConnection() (*sql.DB, error) {
	// Cargamos el archivo .env
	godotenv.Load()

	// Leemos las variables
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, pass, dbName)

	return sql.Open("postgres", connStr)
}

func Migraciones(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS packages (
		id SERIAL PRIMARY KEY,
		tracking_code TEXT UNIQUE NOT NULL,
		receiver_name TEXT NOT NULL,
		destination TEXT NOT NULL,
		weight DOUBLE PRECISION,
		status TEXT,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
	_, err := db.Exec(query)
	return err
}
