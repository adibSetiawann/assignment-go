package service

import (
	"github.com/adibSetiawann/transaction-api-go/model"
)

type MerchantService interface {
	Create(merchantRequest model.CreateMerchant) (model.MerchantResponse, error)
	Update(id int64, merchantRequest model.UpdateMerchant) (model.MerchantResponse, error)
	Remove(id string) error
	GetAllData() ([]model.MerchantResponse, error)
	GetById(id string) ([]model.MerchantResponse, error)
	Validation(merchantRequest interface{}) error
}
