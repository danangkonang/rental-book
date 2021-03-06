package migration

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
)

type Migration struct{}

type DB struct {
	Db *sql.DB
}

var (
	Green = "\033[32m"
	Reset = "\033[0m"
)

func Connection() *DB {
	var connection string
	psql := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	mysql := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	switch os.Getenv("DB_DRIVER") {
	case "postgres":
		connection = psql
	case "mysql":
		connection = mysql
	default:
		connection = psql
	}
	db, err := sql.Open(os.Getenv("DB_DRIVER"), connection)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(100)
	db.SetConnMaxLifetime(5 * time.Minute)
	return &DB{Db: db}
}
