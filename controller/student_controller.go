package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/danangkonang/rental-book/entity"
	"github.com/danangkonang/rental-book/helper"
	"github.com/danangkonang/rental-book/service"
)

type studentController struct {
	service service.StudentService
}

func NewControllerStudent(studentService service.StudentService) *studentController {
	return &studentController{
		service: studentService,
	}
}

func (c *studentController) CreateStudent(w http.ResponseWriter, r *http.Request) {
	var student entity.Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		helper.MakeRespon(w, 400, err.Error(), nil)
		return
	}
	defer r.Body.Close()

	student.CreatedAt = time.Now()
	student.UpdatedAt = time.Now()

	if err := c.service.CreateStudent(&student); err != nil {
		helper.MakeRespon(w, 400, err.Error(), nil)
		return
	}
	helper.MakeRespon(w, 200, "success", nil)
}

func (c *studentController) FindStudents(w http.ResponseWriter, r *http.Request) {
	res, err := c.service.FindStudents()
	if err != nil {
		helper.MakeRespon(w, 400, err.Error(), nil)
		return
	}
	helper.MakeRespon(w, 200, "success", res)
}

func (c *studentController) UpdateStudent(w http.ResponseWriter, r *http.Request) {
	var student entity.Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		helper.MakeRespon(w, 400, err.Error(), nil)
		return
	}
	defer r.Body.Close()

	student.UpdatedAt = time.Now()

	if err := c.service.UpdateStudent(&student); err != nil {
		helper.MakeRespon(w, 400, err.Error(), nil)
		return
	}
	helper.MakeRespon(w, 200, "success", nil)
}

func (c *studentController) DeleteStudent(w http.ResponseWriter, r *http.Request) {
	var student entity.Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		helper.MakeRespon(w, 400, err.Error(), nil)
		return
	}
	defer r.Body.Close()

	if err := c.service.DeleteStudent(&student); err != nil {
		helper.MakeRespon(w, 400, err.Error(), nil)
		return
	}
	helper.MakeRespon(w, 200, "success", nil)
}
