package config

import (
	"database/sql"
	"fmt"
	"scrach_api/helper"

	_ "github.com/lib/pq" // Ensure the postgres driver is imported
	"github.com/rs/zerolog/log"
)

const (
	host    = "localhost"
	port    = 5432
	user    = "macbookairm2"
	pass    = ""
	db_name = "book_db" // Ensure this is the correct database name
)

func ConnectDatabase() *sql.DB {
	// Define connection info
	sqlInfo := fmt.Sprintf(`host=%s port=%d user=%s password='' dbname=%s sslmode=disable`, host, port, user, db_name)

	// Open connection to the database
	db, err := sql.Open("postgres", sqlInfo)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to open database connection")
		helper.PanicIfError(err) // Log the error and stop further execution
	}

	// Ping the database to check the connection
	err = db.Ping()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to ping the database")
		helper.PanicIfError(err)
	}

	log.Info().Msg("Database Connected")
	return db
}
