package service

import (
	"github.com/danangkonang/rental-book/entity"
	"github.com/danangkonang/rental-book/repository"
)

type BookService interface {
	CreateBook(s *entity.Book) error
	FindBooks() ([]entity.Book, error)
	UpdateBook(s *entity.Book) error
	DeleteBook(s *entity.Book) error
}

type bookService struct {
	repository repository.BooksRepository
}

func NewServiceBook(repository repository.BooksRepository) BookService {
	return &bookService{
		repository: repository,
	}
}

func (u *bookService) CreateBook(s *entity.Book) error {
	return u.repository.CreateBook(s)
}

func (u *bookService) FindBooks() ([]entity.Book, error) {
	return u.repository.FindBooks()
}

func (u *bookService) UpdateBook(s *entity.Book) error {
	return u.repository.UpdateBook(s)
}

func (u *bookService) DeleteBook(s *entity.Book) error {
	return u.repository.DeleteBook(s)
}
