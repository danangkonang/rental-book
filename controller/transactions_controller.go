package controller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/danangkonang/rental-book/helper"
	"github.com/danangkonang/rental-book/model"
	"github.com/danangkonang/rental-book/service"
)

type transactionController struct {
	service service.TransactionsService
}

func NewControllerTransaction(transactionsService service.TransactionsService) *transactionController {
	return &transactionController{
		service: transactionsService,
	}
}

func (c *transactionController) FindTransactions(w http.ResponseWriter, r *http.Request) {
	res, err := c.service.FindTransactions()
	if err != nil {
		helper.MakeRespon(w, 400, err.Error(), nil)
		return
	}
	helper.MakeRespon(w, 200, "success", res)
}

func (c *transactionController) CreateTansaction(w http.ResponseWriter, r *http.Request) {
	var tansaction model.TransactionRequest

	if err := json.NewDecoder(r.Body).Decode(&tansaction); err != nil {
		helper.MakeRespon(w, 500, "internal server error", nil)
		return
	}
	defer r.Body.Close()

	tansaction.StrartRental = time.Now()
	tansaction.IsAvailable = false

	if err := c.service.CreateTansaction(&tansaction); err != nil {
		helper.MakeRespon(w, 400, err.Error(), nil)
		return
	}
	helper.MakeRespon(w, 200, "success", nil)
}

func (c *transactionController) ReturnBook(w http.ResponseWriter, r *http.Request) {
	var tansaction model.TransactionRequest

	if err := json.NewDecoder(r.Body).Decode(&tansaction); err != nil {
		helper.MakeRespon(w, 500, "internal server error", nil)
		return
	}
	defer r.Body.Close()

	tansaction.FinishRental = time.Now()
	tansaction.IsAvailable = true
	tansaction.Returned = true
	result, err := c.service.ReturnBook(&tansaction)
	if err != nil {
		helper.MakeRespon(w, 500, err.Error(), nil)
		return
	}
	helper.MakeRespon(w, 200, "success", result)
}
