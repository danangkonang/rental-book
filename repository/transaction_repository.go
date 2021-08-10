package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/danangkonang/rental-book/config"
	"github.com/danangkonang/rental-book/model"
)

type TransactionsRepository interface {
	CreateTansaction(s *model.TransactionRequest) error
	FindTransactions() ([]model.TransactionRespon, error)
	ReturnBook(t *model.TransactionRequest) error
	DetailTransaction(id int, returned bool) (*model.TransactionRespon, error)
}

func NewRepositoryTransaction(Con *config.DB) TransactionsRepository {
	return &conTransactionRepo{
		Psql: Con.Postgresql,
	}
}

type conTransactionRepo struct {
	Psql *sql.DB
}

func (c *conTransactionRepo) FindTransactions() ([]model.TransactionRespon, error) {
	query := `
		SELECT
			t.id,
			t.user_id,
			t.book_id,
			t.start_rental,
			t.finish_rental,
			t.have_penalty,
			t.returned,
			t.created_at,
			t.updated_at,
			b.book_name,
			s.student_name
		FROM
			transactions t
		JOIN
			books b
		ON
			b.id = t.book_id
		JOIN
			students s
		ON
			s.id = t.user_id
	`
	row, err := c.Psql.Query(query)
	if err != nil {
		return nil, errors.New("internal server error")
	}
	var transactions []model.TransactionRespon
	for row.Next() {
		var trans model.TransactionRespon
		err := row.Scan(&trans.Id, &trans.UserId, &trans.BookId, &trans.StrartRental, &trans.FinishRental, &trans.HavePenalty, &trans.Returned, &trans.CreatedAt, &trans.UpdatedAt, &trans.BookName, &trans.StudentName)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, trans)
	}
	defer row.Close()
	if len(transactions) == 0 {
		return make([]model.TransactionRespon, 0), nil
	}
	return transactions, nil
}

func (c *conTransactionRepo) CreateTansaction(t *model.TransactionRequest) error {
	ctx := context.Background()
	tx, err := c.Psql.BeginTx(ctx, nil)
	if err != nil {
		return errors.New("internal server error")
	}
	insertQuery := `INSERT INTO transactions (user_id, book_id) VALUES ($1, $2)`
	_, err = tx.ExecContext(ctx, insertQuery, t.UserId, t.BookId)
	if err != nil {
		tx.Rollback()
		return errors.New("internal server error")
	}
	updateQuery := `UPDATE books SET is_available = $1 WHERE id = $2`
	_, err = tx.ExecContext(ctx, updateQuery, t.IsAvailable, t.BookId)
	if err != nil {
		tx.Rollback()
		return errors.New("internal server error")
	}

	err = tx.Commit()
	if err != nil {
		return errors.New("internal server error")
	}
	return nil
}

func (c *conTransactionRepo) ReturnBook(t *model.TransactionRequest) error {
	ctx := context.Background()
	tx, err := c.Psql.BeginTx(ctx, nil)
	if err != nil {
		return errors.New("internal server error")
	}
	updateTransactionQuery := `
		UPDATE
			transactions
		SET
			finish_rental = $1,
			have_penalty = $2,
			returned = $3
		WHERE id = $4
	`
	_, err = tx.ExecContext(ctx, updateTransactionQuery, t.FinishRental, t.HavePenalty, t.Returned, t.Id)
	if err != nil {
		tx.Rollback()
		return errors.New("internal server error")
	}
	updateBookQuery := `UPDATE books SET is_available = $1 WHERE id = $2`
	_, err = tx.ExecContext(ctx, updateBookQuery, true, t.BookId)
	if err != nil {
		tx.Rollback()
		return errors.New("internal server error")
	}

	err = tx.Commit()
	if err != nil {
		return errors.New("internal server error")
	}
	return nil
}

func (c *conTransactionRepo) DetailTransaction(id int, returned bool) (*model.TransactionRespon, error) {
	var transaction model.TransactionRespon
	query := `
		SELECT
			t.id,
			t.user_id,
			t.book_id,
			t.start_rental,
			t.finish_rental,
			t.have_penalty,
			t.returned,
			t.created_at,
			t.updated_at,
			b.book_name,
			s.student_name
		FROM
			transactions t
		JOIN
			books b
		ON
			b.id = t.book_id
		JOIN
			students s
		ON
			s.id = t.user_id
		WHERE
			t.id = $1 AND t.returned = $2
	`
	row := c.Psql.QueryRow(query, id, returned)
	err := row.Scan(&transaction.Id, &transaction.UserId, &transaction.BookId, &transaction.StrartRental, &transaction.FinishRental, &transaction.HavePenalty, &transaction.Returned, &transaction.CreatedAt, &transaction.UpdatedAt, &transaction.BookName, &transaction.StudentName)
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}
