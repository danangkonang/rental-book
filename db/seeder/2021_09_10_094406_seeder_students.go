package seeder

import (
	"fmt"
	"os"
	"time"

	"github.com/danangkonang/rental-book/db/migration"
	"syreclabs.com/go/faker"
)

func (s *Seeder) Students() {
	start := time.Now()
	query := `
		INSERT INTO
			students (student_name, created_at, updated_at)
		VALUES
			($1, $2, $3)
	`
	for i := 0; i < 10; i++ {
		stmt, _ := migration.Connection().Db.Prepare(query)
		_, err := stmt.Exec(
			faker.Internet().UserName(), time.Now(), time.Now(),
		)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}
	}

	duration := time.Since(start)
	fmt.Println("insert table 2021_09_10_094406_seeder_students.go", string(migration.Green), "success", string(migration.Reset), "in", fmt.Sprintf("%.2f", duration.Seconds()), "second")
}
