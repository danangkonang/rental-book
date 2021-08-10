package service

import (
	"github.com/danangkonang/rental-book/entity"
	"github.com/danangkonang/rental-book/repository"
)

type StudentService interface {
	CreateStudent(s *entity.Student) error
	FindStudents() ([]entity.Student, error)
	UpdateStudent(s *entity.Student) error
	DeleteStudent(s *entity.Student) error
}

type studentService struct {
	repository repository.StudentsRepository
}

func NewServiceStudent(repository repository.StudentsRepository) StudentService {
	return &studentService{
		repository: repository,
	}
}

func (u *studentService) CreateStudent(s *entity.Student) error {
	return u.repository.CreateStudent(s)
}

func (u *studentService) FindStudents() ([]entity.Student, error) {
	return u.repository.FindStudents()
}

func (u *studentService) UpdateStudent(s *entity.Student) error {
	return u.repository.UpdateStudent(s)
}

func (u *studentService) DeleteStudent(s *entity.Student) error {
	return u.repository.DeleteStudent(s)
}
