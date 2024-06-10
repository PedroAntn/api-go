package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func ConnectToDB() (*sql.DB, error) {
	conexao := "user=postgres dbname=bd_biblioteca password=root host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		return nil, err
	}
	return db, nil
}
