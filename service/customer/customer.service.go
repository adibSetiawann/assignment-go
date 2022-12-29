package service

import (
	"github.com/adibSetiawann/transaction-api-go/model"
)

type CustomerService interface {
	Login(formLogin *model.LoginForm) (string, error)
	Create(customerRequest model.CreateCustomer) (model.CustomerResponse, error)
	Update(id int64, customerRequest model.UpdateCustomer) (model.CustomerResponse, error)
	Remove(id string) error
	GetAllData() ([]model.CustomerResponse, error)
	GetById(id string) ([]model.CustomerResponse, error)
	Validation(customerRequest interface{}) error
}