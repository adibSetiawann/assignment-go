package service

import (
	"github.com/adibSetiawann/transaction-api-go/model"
)

type TransactionService interface {
	Create(transactionRequest model.CreateTransaction) (model.TransactionResponse, error)
	Update(invoice string, transactionRequest model.UpdateTransaction) (model.TransactionResponse, error)
	Remove(invoice string) error
	GetAllData() ([]model.TransactionResponse, error)
	GetByInvoice(invoice string) ([]model.TransactionResponse, error)
	Validation(transactionRequest interface{}) error
}