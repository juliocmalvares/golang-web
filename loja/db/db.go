package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectaBanco() *sql.DB {
	connStr := "user=postgres dbname=alura_loja password=password host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err.Error())
	}
	return db
}
