package repository

import (
	"github.com/adibSetiawann/transaction-api-go/entity"
	"github.com/adibSetiawann/transaction-api-go/model"
)

type MerchantRepository interface {
	Create(merchant *entity.Merchant) error
	Update(id int64,merchant *model.UpdateMerchant) error
	FindAll() ([]model.MerchantResponse, error)
	FindById(id string) ([]model.MerchantResponse, error)
	Delete(id string) error
}
