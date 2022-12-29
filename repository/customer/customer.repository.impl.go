package repository

import (
	"time"

	"github.com/adibSetiawann/transaction-api-go/config"
	"github.com/adibSetiawann/transaction-api-go/entity"
	"github.com/adibSetiawann/transaction-api-go/model"
	"github.com/adibSetiawann/transaction-api-go/utils"
	"github.com/golang-jwt/jwt"
)

type CustomerRepositoryImplement struct {
	
}

func (*CustomerRepositoryImplement) Login(loginForm *model.LoginForm) (string, error){
	var customerData entity.Customer

	err := config.DB.Debug().First(&customerData, "email=?", loginForm.Email)
	if err.Error != nil {
		return "user not found in database", err.Error
	}
	isValid, errValid := utils.ConfirmPassword(loginForm.Password, customerData.Password)

	if !isValid {
		return "please input correct password", errValid
	}

	claims := jwt.MapClaims{}
	claims["name"] = customerData.Name
	claims["email"] = customerData.Email
	claims["exp"] = time.Now().Add(time.Minute * 1500).Unix()
	if customerData.Name == "admin" {
		claims["role"] = "admin"
	} else {
		claims["role"] = "customer"
	}

	token, errToken := utils.GenerateToken(&claims)

	if errToken != nil {
		return "failed generate token", errToken
	}


	return token, nil
}

func (*CustomerRepositoryImplement) Create(customer *entity.Customer) error{
	db := config.DB.Debug().Create(&customer)
	if db.Error != nil {
		return db.Error
	}

	return nil
}

func (*CustomerRepositoryImplement) Update(id int64, customer *model.UpdateCustomer) error {
	var customerData entity.Customer

	err := config.DB.Debug().First(&customerData, "id=?", id)
	if err.Error != nil {
		return err.Error
	}
	
	customerData.Name = customer.Name
	customerData.Address = customer.Address
	customerData.Phone = customer.Phone

	errUpdate := config.DB.Debug().Save(&customerData).Error
	if errUpdate != nil {
		return errUpdate
	}

	return nil
}
func (*CustomerRepositoryImplement) Delete(id string) error {
	var customer entity.Customer

	err := config.DB.Debug().First(&customer, "id=?", id).Error
	if err != nil {
		return err
	}

	errDelete := config.DB.Debug().Delete(&customer).Error
	if errDelete != nil {
		return err
	}

	return nil
}

func (*CustomerRepositoryImplement) FindById(id string) ([]model.CustomerResponse, error) {

	var customers []model.CustomerResponse

	err := config.DB.Debug().Preload("Gender").First(&customers, "id=?", id)
	if err.Error != nil {
		return nil, err.Error
	}

	return customers, nil
}

func(*CustomerRepositoryImplement) FindAll() ([]model.CustomerResponse, error) {

	var customers []model.CustomerResponse

	db := config.DB.Debug().Preload("Gender").Find(&customers)
	if db.Error != nil {
		return nil, db.Error
	}

	return customers, nil
}


func NewCustomerRepository()  CustomerRepository{
	return &CustomerRepositoryImplement{}
}