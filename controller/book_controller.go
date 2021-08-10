package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/danangkonang/rental-book/entity"
	"github.com/danangkonang/rental-book/helper"
	"github.com/danangkonang/rental-book/service"
)

type bookController struct {
	service service.BookService
}

func NewControllerBook(booktService service.BookService) *bookController {
	return &bookController{
		service: booktService,
	}
}

func (c *bookController) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book entity.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		helper.MakeRespon(w, 400, err.Error(), nil)
		return
	}
	defer r.Body.Close()

	book.CreatedAt = time.Now()
	book.UpdatedAt = time.Now()

	if err := c.service.CreateBook(&book); err != nil {
		helper.MakeRespon(w, 400, err.Error(), nil)
		return
	}
	helper.MakeRespon(w, 200, "success", nil)
}

func (c *bookController) FindBooks(w http.ResponseWriter, r *http.Request) {
	res, err := c.service.FindBooks()
	if err != nil {
		helper.MakeRespon(w, 400, err.Error(), nil)
		return
	}
	helper.MakeRespon(w, 200, "success", res)
}

func (c *bookController) UpdateBook(w http.ResponseWriter, r *http.Request) {
	var book entity.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		helper.MakeRespon(w, 400, err.Error(), nil)
		return
	}
	defer r.Body.Close()

	book.UpdatedAt = time.Now()

	if err := c.service.UpdateBook(&book); err != nil {
		helper.MakeRespon(w, 400, err.Error(), nil)
		return
	}
	helper.MakeRespon(w, 200, "success", nil)
}

func (c *bookController) DeleteBook(w http.ResponseWriter, r *http.Request) {
	var book entity.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		helper.MakeRespon(w, 400, err.Error(), nil)
		return
	}
	defer r.Body.Close()

	if err := c.service.DeleteBook(&book); err != nil {
		helper.MakeRespon(w, 400, err.Error(), nil)
		return
	}
	helper.MakeRespon(w, 200, "success", nil)
}
