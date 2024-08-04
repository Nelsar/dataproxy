package db

import (
	"database/sql"
	"dataproxy/configuration"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func DbConnect() *sql.DB {
	c := configuration.Exec().Database
	db, err := sql.Open("mysql", c.Dsn)

	if err != nil {
		log.Fatalln(err)
	}
	return db
}
