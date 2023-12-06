package db

import (
	"fmt"
	"log"

	"database/sql"

	//Postgres Driver imported
	_ "github.com/lib/pq"
)

// ConnectDB connect to Postgres DB
func ConnectDB() (*sql.DB, error) {
	var (
		host     = "localhost"
		user     = "postgres"
		port     = 5432
		password = "postgres"
		name     = "postgres"
	)
	//Connect to DB
	var DB *sql.DB
	DB, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, name))

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
