package database

import (
	"database/sql"
	"time"

	"github.com/naomigrain/httprouter-crud-notes/config"
)

func NewDB(config config.DBConfig) *sql.DB {
	connString := config.Username + "@tcp(" + config.Host + ":" + config.Port + ")/" + config.Name
	db, err := sql.Open(config.Driver, connString)

	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
