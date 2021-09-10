package migration

import (
	"fmt"
	"os"
)

func (m *Migration) Students() {
	query := `
		CREATE TABLE students(
			id serial PRIMARY KEY,
			student_name VARCHAR (225) NOT NULL,
			created_at TIMESTAMP NULL,
			updated_at TIMESTAMP NULL
		)
	`
	_, err := Connection().Db.Exec(query)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	fmt.Println(string(Green), "success", string(Reset), "create table 2021_09_10_094308_migration_students.go")
}
