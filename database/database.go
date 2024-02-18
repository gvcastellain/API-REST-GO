package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func StartDB() {
	var err error
	db, err = sql.Open("mysql", "root:123456@/")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS API")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("USE API")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS usuarios (
		id INTEGER AUTO_INCREMENT,
		nome VARCHAR(80),
		PRIMARY KEY (id)
	)`)
	if err != nil {
		panic(err)
	}
}

func GetDB() *sql.DB {
	return db
}
