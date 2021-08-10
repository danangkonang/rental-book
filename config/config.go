package config

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type DB struct {
	Postgresql *sql.DB
}

func Connection() *DB {
	connection := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
		"postgres-db",
		5432,
		"postgres",
		"postgres",
		"default",
	)
	db, err := sql.Open("postgres", connection)
	if err != nil {
		log.Panic(err.Error())
	}
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(100)
	db.SetConnMaxLifetime(5 * time.Minute)
	return &DB{Postgresql: db}
}
