package service

import (
	"errors"
	"log"
	"strconv"

	"github.com/adibSetiawann/transaction-api-go/entity"
	"github.com/adibSetiawann/transaction-api-go/model"
	repository "github.com/adibSetiawann/transaction-api-go/repository/product"
	"github.com/go-playground/validator/v10"
)

type ProductServiceImpl struct {
	productRepo repository.ProductRepository
}

func (ps *ProductServiceImpl) Create(request model.CreateProduct) (model.ProductResponse, error) {
	var productResponse model.ProductResponse

	product := entity.Product{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		Stock:       request.Stock,
		MerchantId:  request.MerchantId,
	}

	errCreate := ps.productRepo.Create(&product)
	if errCreate != nil {
		log.Println(errCreate.Error())
		return productResponse, errCreate
	}

	productResponse.ID = product.ID
	productResponse.Name = product.Name
	productResponse.Description = product.Description
	productResponse.Price = product.Price
	productResponse.Stock = product.Stock
	productResponse.MerchantId = product.MerchantId
	return productResponse, nil
}

func (ms *ProductServiceImpl) Update(id int64, request model.UpdateProduct) (model.ProductResponse, error) {
	var productResponse model.ProductResponse
	n := strconv.FormatInt(id, 10)

	errCreate := ms.productRepo.Update(id, &request)

	if errCreate != nil {
		log.Println(errCreate.Error())
		return productResponse, errCreate
	}

	products, _ := ms.productRepo.FindById(n)

	productResponse.ID = products[0].ID
	productResponse.Name = request.Name
	productResponse.Description = request.Description
	productResponse.Stock = request.Stock
	productResponse.Price = request.Price
	productResponse.MerchantId = products[0].MerchantId
	productResponse.Merchant = products[0].Merchant
	return productResponse, nil
}

func (ms *ProductServiceImpl) Remove(id string) error {
	errDelete := ms.productRepo.Delete(id)
	if errDelete != nil {
		log.Println(errDelete.Error())
		return errDelete
	}

	return nil
}

func (ms *ProductServiceImpl) GetById(id string) ([]model.ProductResponse, error) {
	merchants, errFind := ms.productRepo.FindById(id)
	if errFind != nil {
		return nil, errFind
	}
	return merchants, nil
}

func (ps *ProductServiceImpl) GetAllData() ([]model.ProductResponse, error) {
	products, errFind := ps.productRepo.FindAll()
	if errFind != nil {
		return nil, errFind
	}
	return products, nil
}

func (ms *ProductServiceImpl) Validation(userRequest interface{}) error {
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

func NewProductService(productRepo *repository.ProductRepository) ProductService {
	return &ProductServiceImpl{
		productRepo: *productRepo,
	}
}
