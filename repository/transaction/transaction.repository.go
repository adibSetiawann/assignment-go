package repository

import (
	"github.com/adibSetiawann/transaction-api-go/entity"
	"github.com/adibSetiawann/transaction-api-go/model"
)

type TransactionRepository interface {
	Create(transaction *entity.Transaction) error
	Update(invoice string, transaction *model.UpdateTransaction) error
	FindAll() ([]model.TransactionResponse, error)
	FindByInvoice(invoice string) ([]model.TransactionResponse, error)
	DeleteByInvoice(invoice string) error
}
