package migration

import (
	"fmt"
	"os"
)

func (m *Migration) Books() {
	query := `
		CREATE TABLE books(
			id serial PRIMARY KEY,
			book_name VARCHAR (225) NOT NULL,
			max_rental_days INTEGER NOT NULL,
			penalty_per_day INTEGER NOT NULL,
			is_available BOOLEAN DEFAULT TRUE,
			created_at TIMESTAMP NULL,
			updated_at TIMESTAMP NULL
		)
	`
	_, err := Connection().Db.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	fmt.Println(string(Green), "success", string(Reset), "create table 2021_09_10_094334_migration_books.go")
}
