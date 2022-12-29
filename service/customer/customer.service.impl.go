package service

import (
	"errors"
	"log"
	"strconv"

	"github.com/adibSetiawann/transaction-api-go/entity"
	"github.com/adibSetiawann/transaction-api-go/model"
	repository "github.com/adibSetiawann/transaction-api-go/repository/customer"
	"github.com/adibSetiawann/transaction-api-go/utils"
	"github.com/go-playground/validator/v10"
)

type CustomerServiceImpl struct {
	customerRepo repository.CustomerRepository
}

func (cs *CustomerServiceImpl) Login(formLogin *model.LoginForm) (string, error) {
	token, err := cs.customerRepo.Login(formLogin)
	return token, err
}

func (cs *CustomerServiceImpl) Create(request model.CreateCustomer) (model.CustomerResponse, error) {
	var customerResponse model.CustomerResponse

	hashedPassword, err := utils.HashPassword(request.Password)
	if err != nil {
		log.Println(err)
		return customerResponse, nil
	}
	customer := entity.Customer{
		Name:     request.Name,
		Email:    request.Email,
		Password: hashedPassword,
		Phone:    request.Phone,
		Address:  request.Address,
		GenderID: request.GenderID,
	}

	errCreate := cs.customerRepo.Create(&customer)
	if errCreate != nil {
		log.Println(errCreate.Error())
		return customerResponse, errCreate
	}

	customerResponse.ID = customer.ID
	customerResponse.Name = customer.Name
	customerResponse.Email = customer.Email
	customerResponse.Phone = customer.Phone
	customerResponse.Address = customer.Address
	customerResponse.GenderID = customer.GenderID
	return customerResponse, nil
}

func (ms *CustomerServiceImpl) Update(id int64, request model.UpdateCustomer) (model.CustomerResponse, error) {
	var customerResponse model.CustomerResponse
	n := strconv.FormatInt(id, 10)

	errCreate := ms.customerRepo.Update(id, &request)

	if errCreate != nil {
		log.Println(errCreate.Error())
		return customerResponse, errCreate
	}

	customers, _ := ms.customerRepo.FindById(n)

	customerResponse.ID = customers[0].ID
	customerResponse.Name = request.Name
	customerResponse.Email = customers[0].Email
	customerResponse.Address = request.Address
	customerResponse.Phone = request.Phone
	customerResponse.GenderID = customers[0].GenderID
	customerResponse.Gender = customers[0].Gender
	return customerResponse, nil
}

func (ms *CustomerServiceImpl) Remove(id string) error {
	errDelete := ms.customerRepo.Delete(id)
	if errDelete != nil {
		log.Println(errDelete.Error())
		return errDelete
	}

	return nil
}

func (ms *CustomerServiceImpl) GetById(id string) ([]model.CustomerResponse, error) {
	merchants, errFind := ms.customerRepo.FindById(id)
	if errFind != nil {
		return nil, errFind
	}
	return merchants, nil
}

func (cs *CustomerServiceImpl) GetAllData() ([]model.CustomerResponse, error) {
	customers, errFind := cs.customerRepo.FindAll()
	if errFind != nil {
		return nil, errFind
	}
	return customers, nil
}

func (ms *CustomerServiceImpl) Validation(userRequest interface{}) error {
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

func NewCustomerService(customerRepo *repository.CustomerRepository) CustomerService {
	return &CustomerServiceImpl{
		customerRepo: *customerRepo,
	}
}
