package repository

import (
	"github.com/adibSetiawann/transaction-api-go/entity"
	"github.com/adibSetiawann/transaction-api-go/model"
)

type ProductRepository interface {
	Create(product *entity.Product) error
	Update(id int64, product *model.UpdateProduct) error
	FindAll() ([]model.ProductResponse, error)
	FindById(id string) ([]model.ProductResponse, error)
	Delete(id string) error
}
