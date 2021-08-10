package repository

import (
	"database/sql"
	"errors"

	"github.com/danangkonang/rental-book/config"
	"github.com/danangkonang/rental-book/entity"
	"github.com/danangkonang/rental-book/model"
)

type BooksRepository interface {
	CreateBook(s *entity.Book) error
	FindBooks() ([]entity.Book, error)
	DetailBook(id int) (*entity.Book, error)
	UpdateBook(s *entity.Book) error
	DeleteBook(s *entity.Book) error
	IsBookAvailable(b *model.TransactionRequest) error
}

func NewRepositoryBook(Con *config.DB) BooksRepository {
	return &conBookRepo{
		Psql: Con.Postgresql,
	}
}

type conBookRepo struct {
	Psql *sql.DB
}

func (c *conBookRepo) CreateBook(b *entity.Book) error {
	query := `
		INSERT INTO
			books (book_name, max_rental_days, penalty_per_day, is_available, created_at, updated_at)
		VALUES
			($1, $2, $3, $4, $5, $6)
	`
	_, err := c.Psql.Exec(query, b.BookName, b.MaxRentalDays, b.PenaltyPerDay, b.IsAvailable, b.CreatedAt, b.UpdatedAt)
	return err
}

func (c *conBookRepo) FindBooks() ([]entity.Book, error) {
	query := `
		SELECT
			*
		FROM
			books
	`
	row, err := c.Psql.Query(query)
	if err != nil {
		return nil, errors.New("internal server error")
	}
	var books []entity.Book
	for row.Next() {
		var book entity.Book
		err := row.Scan(&book.Id, &book.BookName, &book.MaxRentalDays, &book.PenaltyPerDay, &book.IsAvailable, &book.CreatedAt, &book.UpdatedAt)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	defer row.Close()
	if len(books) == 0 {
		return make([]entity.Book, 0), nil
	}
	return books, nil
}

func (c *conBookRepo) DetailBook(id int) (*entity.Book, error) {
	var book entity.Book
	query := `
		SELECT
			id, max_rental_days, penalty_per_day
		FROM
			books
		WHERE
			id = $1
	`
	row := c.Psql.QueryRow(query, id)
	err := row.Scan(&book.Id, &book.MaxRentalDays, &book.PenaltyPerDay)
	return &book, err
}

func (c *conBookRepo) UpdateBook(b *entity.Book) error {
	var query string = `
		UPDATE
			books
		SET
			book_name = $1, max_rental_days = $2, penalty_per_day = $3, is_available = $4, updated_at = $5
		WHERE
			id = $6
	`
	_, err := c.Psql.Exec(query, b.BookName, b.MaxRentalDays, b.PenaltyPerDay, b.IsAvailable, b.UpdatedAt, b.Id)
	return err
}

func (c *conBookRepo) DeleteBook(b *entity.Book) error {
	var query string = `
		DELETE
		FROM
			books
		WHERE
			id = $1
	`
	_, err := c.Psql.Exec(query, b.Id)
	return err
}

func (c *conBookRepo) IsBookAvailable(b *model.TransactionRequest) error {
	var transaction model.TransactionRespon
	var query string = `
		SELECT
			book_name
		FROM
			books
		WHERE
			id = $1 
		AND
			is_available = $2
	`
	result := c.Psql.QueryRow(query, b.BookId, true)
	// fmt.Println(result)
	row := result.Scan(&transaction.BookName)
	// fmt.Println(transaction)
	return row
}
