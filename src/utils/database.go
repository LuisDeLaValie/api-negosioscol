package utils

import (
	"fmt"
	"log"

	"database/sql"

	//Postgres Driver imported
	_ "github.com/lib/pq"
)

// ConnectDB connect to Postgres DB
func ConnectDB() *sql.DB {
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
		err = Error500(fmt.Sprintf("Error in connect the DB %v", err))
		log.Fatalf(err.Error())
		return nil
	}
	if err := DB.Ping(); err != nil {
		err = Error500(fmt.Sprintf("Error in make ping the DB " + err.Error()))
		log.Fatalf(err.Error())

		return nil
	}

	log.Println("DB connected")
	return DB
}
