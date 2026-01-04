package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // El guion bajo es porque solo necesitamos el driver
)

func GetConnection() (*sql.DB, error) {
	// Estos datos deben coincidir con tu docker-compose.yml
	connStr := "host=localhost port=5432 user=user_logistics password=password123 dbname=last_mile_db sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Verificar si la conexión es real
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("✅ Conexión a la base de datos exitosa!")
	return db, nil
}
