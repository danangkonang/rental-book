package seeder

import (
	"fmt"
	"os"
	"time"

	"github.com/danangkonang/rental-book/db/migration"
)

func (s *Seeder) Books() {
	start := time.Now()
	query := `
		INSERT INTO
			books (book_name, max_rental_days, penalty_per_day, is_available, created_at, updated_at)
		VALUES
			($1, $2, $3, $4, $5, $6),
			($7, $8, $9, $10, $11, $12),
			($13, $14, $15, $16, $17, $18),
			($19, $20, $21, $22, $23, $24),
			($25, $26, $27, $28, $29, $30),
			($31, $32, $33, $34, $35, $36)
	`
	stmt, _ := migration.Connection().Db.Prepare(query)
	_, err := stmt.Exec(
		"komik doraemon", 3, 5000, false, time.Now(), time.Now(),
		"kamus bahasa indonesia", 3, 5000, false, time.Now(), time.Now(),
		"pengusaha sukses", 3, 5000, false, time.Now(), time.Now(),
		"pintar matematika", 3, 5000, false, time.Now(), time.Now(),
		"tokoh dunia", 3, 5000, false, time.Now(), time.Now(),
		"sejarah indonesia", 3, 5000, true, time.Now(), time.Now(),
	)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	duration := time.Since(start)
	fmt.Println("insert table 2021_09_10_094502_seeder_books.go", string(migration.Green), "success", string(migration.Reset), "in", fmt.Sprintf("%.2f", duration.Seconds()), "second")
}
