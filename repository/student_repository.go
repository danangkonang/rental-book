package repository

import (
	"database/sql"
	"errors"

	"github.com/danangkonang/rental-book/config"
	"github.com/danangkonang/rental-book/entity"
)

type StudentsRepository interface {
	CreateStudent(s *entity.Student) error
	FindStudents() ([]entity.Student, error)
	UpdateStudent(s *entity.Student) error
	DeleteStudent(s *entity.Student) error
}

func NewRepositoryStudent(Con *config.DB) StudentsRepository {
	return &conStudentRepo{
		Psql: Con.Postgresql,
	}
}

type conStudentRepo struct {
	Psql *sql.DB
}

func (c *conStudentRepo) CreateStudent(s *entity.Student) error {
	query := `
		INSERT INTO
			students (student_name, created_at, updated_at)
		VALUES
			($1, $2, $3)
	`
	_, err := c.Psql.Exec(query, s.StudentName, s.CreatedAt, s.UpdatedAt)
	return err
}

func (c *conStudentRepo) FindStudents() ([]entity.Student, error) {
	query := `
		SELECT
			*
		FROM
			students
	`
	row, err := c.Psql.Query(query)
	if err != nil {
		return nil, errors.New("internal server error")
	}
	var students []entity.Student
	for row.Next() {
		var student entity.Student
		err := row.Scan(&student.Id, &student.StudentName, &student.CreatedAt, &student.UpdatedAt)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	defer row.Close()
	if len(students) == 0 {
		return make([]entity.Student, 0), nil
	}
	return students, nil
}

func (c *conStudentRepo) UpdateStudent(s *entity.Student) error {
	var query string = `
		UPDATE
			students
		SET
			student_name = $1, updated_at = $2
		WHERE
			id = $3
	`
	_, err := c.Psql.Exec(query, s.StudentName, s.UpdatedAt, s.Id)
	return err
}

func (c *conStudentRepo) DeleteStudent(s *entity.Student) error {
	var query string = `
		DELETE
		FROM
			students
		WHERE
			id = $1
	`
	_, err := c.Psql.Exec(query, s.Id)
	return err
}
