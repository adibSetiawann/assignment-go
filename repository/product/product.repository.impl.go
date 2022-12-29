package repository

import (
	"github.com/adibSetiawann/transaction-api-go/config"
	"github.com/adibSetiawann/transaction-api-go/entity"
	"github.com/adibSetiawann/transaction-api-go/model"
)

type ProductRepositoryImplement struct {
	
}

func (*ProductRepositoryImplement) Create(product *entity.Product) error{
	db := config.DB.Debug().Create(&product)
	if db.Error != nil {
		return db.Error
	}

	return nil
}

func (*ProductRepositoryImplement) Update(id int64, product *model.UpdateProduct) error {
	var productData entity.Product

	err := config.DB.Debug().First(&productData, "id=?", id)
	if err.Error != nil {
		return err.Error
	}
	
	productData.Name = product.Name
	productData.Description = product.Description
	productData.Stock = product.Stock
	productData.Price = product.Price

	errUpdate := config.DB.Debug().Save(&productData).Error
	if errUpdate != nil {
		return errUpdate
	}

	return nil
}
func (*ProductRepositoryImplement) Delete(id string) error {
	var product entity.Product

	err := config.DB.Debug().First(&product, "id=?", id).Error
	if err != nil {
		return err
	}

	errDelete := config.DB.Debug().Delete(&product).Error
	if errDelete != nil {
		return err
	}

	return nil
}

func (*ProductRepositoryImplement) FindById(id string) ([]model.ProductResponse, error) {

	var products []model.ProductResponse

	err := config.DB.Debug().Preload("Merchant").First(&products, "id=?", id)
	if err.Error != nil {
		return nil, err.Error
	}

	return products, nil
}

func(*ProductRepositoryImplement) FindAll() ([]model.ProductResponse, error) {

	var products []model.ProductResponse

	db := config.DB.Debug().Preload("Merchant").Find(&products)
	if db.Error != nil {
		return nil, db.Error
	}

	return products, nil
}


func NewProductRepository()  ProductRepository{
	return &ProductRepositoryImplement{}
}