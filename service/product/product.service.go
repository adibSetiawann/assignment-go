package service

import (
	"github.com/adibSetiawann/transaction-api-go/model"
)

type ProductService interface {
	Create(productRequest model.CreateProduct) (model.ProductResponse, error)
	Update(id int64, productRequest model.UpdateProduct) (model.ProductResponse, error)
	Remove(id string) error
	GetAllData() ([]model.ProductResponse, error)
	GetById(id string) ([]model.ProductResponse, error)
	Validation(productRequest interface{}) error
}