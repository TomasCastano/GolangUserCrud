package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	// Cargar el archivo .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal(" Error cargando el archivo .env")
	}

	// Variables de entorno
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Construir la cadena de conexión
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	// Conectar a PostgreSQL
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(" Error conectando a la base de datos:", err)
	}
	// Probar conexión
	err = DB.Ping()

	if err != nil {
		log.Fatal(" Error en la conexión:", err)
	}

	log.Println(" Conectado a la base de datos")
}
