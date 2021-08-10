package migration

import (
	"fmt"
	"os"
)

func (m *Migration) Students() {
	_, err := Connection().Db.Exec(`
		CREATE TABLE students(
			id serial PRIMARY KEY,
			student_name VARCHAR (225) NOT NULL,
			created_at TIMESTAMP NULL,
			updated_at TIMESTAMP NULL
		);
	`)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	fmt.Println("success create table 2021_08_07_045755_migration_students.go")
}
