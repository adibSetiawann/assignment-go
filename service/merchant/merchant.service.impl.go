package service

import (
	"errors"
	"log"
	"strconv"

	"github.com/adibSetiawann/transaction-api-go/entity"
	"github.com/adibSetiawann/transaction-api-go/model"
	repository "github.com/adibSetiawann/transaction-api-go/repository/merchant"
	"github.com/go-playground/validator/v10"
)

type MerchantServiceImpl struct {
	merchantRepo repository.MerchantRepository
}

func (ms *MerchantServiceImpl) Create(request model.CreateMerchant) (model.MerchantResponse, error) {
	var merchantResponse model.MerchantResponse

	merchant := entity.Merchant{
		Name:    request.Name,
		Address: request.Address,
		Email:   request.Email,
		Phone:   request.Phone,
	}

	errCreate := ms.merchantRepo.Create(&merchant)
	if errCreate != nil {
		log.Println(errCreate.Error())
		return merchantResponse, errCreate
	}

	merchantResponse.ID = merchant.ID
	merchantResponse.Name = merchant.Name
	merchantResponse.Address = merchant.Address
	merchantResponse.Email = merchant.Email
	merchantResponse.Phone = merchant.Phone
	return merchantResponse, nil
}

func (ms *MerchantServiceImpl) Update(id int64, request model.UpdateMerchant) (model.MerchantResponse, error) {
	var merchantResponse model.MerchantResponse
	n := strconv.FormatInt(id, 10)

	errCreate := ms.merchantRepo.Update(id, &request)

	if errCreate != nil {
		log.Println(errCreate.Error())
		return merchantResponse, errCreate
	}

	merchants, _ := ms.merchantRepo.FindById(n)

	merchantResponse.ID = merchants[0].ID
	merchantResponse.Name = request.Name
	merchantResponse.Address = request.Address
	merchantResponse.Email = merchants[0].Email
	merchantResponse.Phone = request.Phone
	return merchantResponse, nil
}

func (ms *MerchantServiceImpl) Remove(id string) error {
	errDelete := ms.merchantRepo.Delete(id)
	if errDelete != nil {
		log.Println(errDelete.Error())
		return errDelete
	}

	return nil
}

func (ms *MerchantServiceImpl) GetById(id string) ([]model.MerchantResponse, error) {
	merchants, errFind := ms.merchantRepo.FindById(id)
	if errFind != nil {
		return nil, errFind
	}
	return merchants, nil
}

func (ms *MerchantServiceImpl) GetAllData() ([]model.MerchantResponse, error) {
	merchants, errFind := ms.merchantRepo.FindAll()
	if errFind != nil {
		return nil, errFind
	}
	return merchants, nil
}

func (ms *MerchantServiceImpl) Validation(userRequest interface{}) error {
	var messageError string
	var isError bool

	// VALIDATION USER INPUT
	validate := validator.New()
	errValidate := validate.Struct(userRequest)
	if errValidate != nil {
		messageError += errValidate.Error()
		isError = true
	}

	if isError {
		return errors.New(messageError)
	}

	return nil
}

func NewMerchantService(merchantRepo *repository.MerchantRepository) MerchantService {
	return &MerchantServiceImpl{
		merchantRepo: *merchantRepo,
	}
}
