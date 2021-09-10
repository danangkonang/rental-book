package seeder

import (
	"fmt"
	"os"
	"time"

	"github.com/danangkonang/rental-book/db/migration"
)

func (s *Seeder) Transactions() {
	start := time.Now()
	query := `
		INSERT INTO
			transactions (user_id, book_id, start_rental, created_at, updated_at)
		VALUES
			($1, $2, $3, $4, $5),
			($6, $7, $8, $9, $10),
			($11, $12, $13, $14, $15),
			($16, $17, $18, $19, $20),
			($21, $22, $23, $24, $25)
	`
	stmt, _ := migration.Connection().Db.Prepare(query)
	_, err := stmt.Exec(
		1, 1, time.Now(), time.Now(), time.Now(),
		2, 2, time.Now().AddDate(0, 0, -2), time.Now(), time.Now(),
		3, 3, time.Now().AddDate(0, 0, -3), time.Now(), time.Now(),
		4, 4, time.Now().AddDate(0, 0, -5), time.Now(), time.Now(),
		5, 5, time.Now().AddDate(0, 0, -10), time.Now(), time.Now(),
	)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	duration := time.Since(start)
	fmt.Println("insert table 2021_09_10_094428_seeder_transactions.go", string(migration.Green), "success", string(migration.Reset), "in", fmt.Sprintf("%.2f", duration.Seconds()), "second")
}
