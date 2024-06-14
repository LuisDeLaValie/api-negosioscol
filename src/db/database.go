package db

import (
	"fmt"
	"log"
	"os"

	"database/sql"

	//Postgres Driver imported

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// ConnectDB connect to Postgres DB
func ConnectDB() (*sql.DB, error) {

	// Obtener la URL de conexión de PostgreSQL de las variables de entorno
	postgresURL := os.Getenv("ConnectPosgreSQL")
	if postgresURL == "" {
		// Cargar las variables de entorno del archivo .env
		err := godotenv.Load()
		if err != nil {
			fmt.Printf("Error al cargar el archivo .env: %v", err)
		}

		postgresURL = os.Getenv("ConnectPosgreSQL")
		if postgresURL == "" {

			log.Fatal("La variable ConnectPosgreSQL no está definida en el archivo .env")
		}
		fmt.Println(postgresURL)
	}

	fmt.Println(postgresURL)
	//Connect to DB
	var DB *sql.DB
	DB, err := sql.Open("postgres", postgresURL)

	if err != nil {
		log.Fatalf(err.Error())
		return nil, fmt.Errorf("Error in connect the DB %v", err)
	}
	if err := DB.Ping(); err != nil {
		log.Fatalf(err.Error())

		return nil, fmt.Errorf("Error in make ping the DB: %v ", err)
	}

	log.Println("DB connected")
	return DB, nil
}
