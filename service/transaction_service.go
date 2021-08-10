package service

import (
	"database/sql"
	"errors"
	"math"
	"time"

	"github.com/danangkonang/rental-book/model"
	"github.com/danangkonang/rental-book/repository"
)

type TransactionsService interface {
	FindTransactions() ([]model.TransactionRespon, error)
	CreateTansaction(t *model.TransactionRequest) error
	ReturnBook(t *model.TransactionRequest) (*model.TransactionRespon, error)
}

type transactionService struct {
	repoTran repository.TransactionsRepository
	repoBook repository.BooksRepository
}

func NewServiceTransaction(repoTran repository.TransactionsRepository, repoBook repository.BooksRepository) TransactionsService {
	return &transactionService{
		repoTran: repoTran,
		repoBook: repoBook,
	}
}

func (u *transactionService) FindTransactions() ([]model.TransactionRespon, error) {
	return u.repoTran.FindTransactions()
}

func (u *transactionService) CreateTansaction(t *model.TransactionRequest) error {
	if err := u.repoBook.IsBookAvailable(t); err != nil {
		if err == sql.ErrNoRows {
			return errors.New("book is not available")
		}
		return errors.New("internal server error")
	}

	return u.repoTran.CreateTansaction(t)
}

func (u *transactionService) ReturnBook(t *model.TransactionRequest) (*model.TransactionRespon, error) {
	trRespon, err := u.repoTran.DetailTransaction(t.Id, false)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("wrong transaction id")
		}
		return nil, errors.New("internal server error")
	}

	bookRespon, err := u.repoBook.DetailBook(trRespon.BookId)
	if err != nil {
		return nil, errors.New("internal server error")
	}

	oneDay := 60 * 60 * 24
	timeReturned := time.Now()
	studentRental := timeReturned.Add(7 * time.Hour).Sub(trRespon.StrartRental).Seconds()
	limitRentalInSecond := oneDay * bookRespon.MaxRentalDays
	penalty := studentRental - float64(limitRentalInSecond)
	if penalty > 0 {
		t.HavePenalty = int(math.Ceil(penalty/float64(oneDay))) * bookRespon.PenaltyPerDay
	}
	t.FinishRental = timeReturned
	t.BookId = trRespon.Id

	if err := u.repoTran.ReturnBook(t); err != nil {
		return nil, err
	}

	detail, err := u.repoTran.DetailTransaction(t.Id, true)
	if err != nil {
		return nil, errors.New("internal server error")
	}
	return detail, nil
}
