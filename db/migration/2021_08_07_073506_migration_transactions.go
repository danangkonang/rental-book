package migration

import (
	"fmt"
	"os"
)

func (m *Migration) Transactions() {
	_, err := Connection().Db.Exec(`
		CREATE TABLE transactions(
			id serial PRIMARY KEY,
			user_id INTEGER NOT NULL,
			book_id INTEGER NOT NULL,
			start_rental TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
			finish_rental TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			have_penalty INTEGER DEFAULT 0,
			returned BOOLEAN DEFAULT FALSE,
			created_at TIMESTAMP NULL,
			updated_at TIMESTAMP NULL
		);
	`)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	fmt.Println("success create table 2021_08_07_073506_migration_transactions.go")
}
