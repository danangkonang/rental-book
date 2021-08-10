package migration

import (
	"fmt"
	"os"
)

func (m *Migration) Books() {
	_, err := Connection().Db.Exec(`
		CREATE TABLE books(
			id serial PRIMARY KEY,
			book_name VARCHAR (225) NOT NULL,
			max_rental_days INTEGER NOT NULL,
			penalty_per_day INTEGER NOT NULL,
			is_available BOOLEAN DEFAULT TRUE,
			created_at TIMESTAMP NULL,
			updated_at TIMESTAMP NULL
		);
	`)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	fmt.Println("success create table 2021_08_07_073516_migration_books.go")
}
