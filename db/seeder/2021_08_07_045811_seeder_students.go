package seeder

import (
	"fmt"
	"os"
	"time"

	"github.com/danangkonang/rental-book/db/migration"
	"syreclabs.com/go/faker"
)

func (s *Seeder) Students() {
	for i := 0; i < 10; i++ {
		query := `
			INSERT INTO
				students (student_name, created_at, updated_at)
			VALUES
				($1, $2, $3)
		`
		_, err := migration.Connection().Db.Exec(
			query,
			faker.Internet().UserName(), time.Now(), time.Now(),
		)

		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}
	}
	fmt.Println("success insert table 2021_08_07_045811_seeder_students.go")
}
