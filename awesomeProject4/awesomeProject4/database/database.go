package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func ConnectToDB() *sql.DB {
	conexao := "user=postgres dbname=postgres password=minhasenha host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		log.Fatalf("Error opening the database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	log.Println("Successfully connected to the database")
	return db
}
