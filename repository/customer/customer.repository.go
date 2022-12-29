package repository

import (
	"github.com/adibSetiawann/transaction-api-go/entity"
	"github.com/adibSetiawann/transaction-api-go/model"
)

type CustomerRepository interface {
	Login(form *model.LoginForm) (string, error)
	Create(customer *entity.Customer) error
	Update(id int64, product *model.UpdateCustomer) error
	FindAll() ([]model.CustomerResponse, error)
	FindById(id string) ([]model.CustomerResponse, error)
	Delete(id string) error
}
