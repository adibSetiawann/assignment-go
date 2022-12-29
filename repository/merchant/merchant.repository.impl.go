package repository

import (
	"github.com/adibSetiawann/transaction-api-go/config"
	"github.com/adibSetiawann/transaction-api-go/entity"
	"github.com/adibSetiawann/transaction-api-go/model"
)

type MerchantRepositoryImplement struct {
}

func (*MerchantRepositoryImplement) Create(merchant *entity.Merchant) error {
	db := config.DB.Debug().Create(&merchant)
	if db.Error != nil {
		return db.Error
	}

	return nil
}

func (*MerchantRepositoryImplement) Update(id int64, merchant *model.UpdateMerchant) error {
	var merchantData entity.Merchant

	err := config.DB.Debug().First(&merchantData, "id=?", id)
	if err.Error != nil {
		return err.Error
	}
	
	merchantData.Name = merchant.Name
	merchantData.Address = merchant.Address
	merchantData.Phone = merchant.Phone

	errUpdate := config.DB.Debug().Save(&merchantData).Error
	if errUpdate != nil {
		return errUpdate
	}

	return nil
}

func (*MerchantRepositoryImplement) Delete(id string) error {
	var merchant entity.Merchant

	err := config.DB.Debug().First(&merchant, "id=?", id).Error
	if err != nil {
		return err
	}

	errDelete := config.DB.Debug().Delete(&merchant).Error
	if errDelete != nil {
		return err
	}

	return nil
}

func (*MerchantRepositoryImplement) FindById(id string) ([]model.MerchantResponse, error) {

	var merchants []model.MerchantResponse

	err := config.DB.Debug().First(&merchants, "id=?", id)
	if err.Error != nil {
		return nil, err.Error
	}

	return merchants, nil
}

func (*MerchantRepositoryImplement) FindAll() ([]model.MerchantResponse, error) {

	var merchants []model.MerchantResponse

	db := config.DB.Debug().Find(&merchants)
	if db.Error != nil {
		return nil, db.Error
	}

	return merchants, nil
}

func NewMerchantRepository() MerchantRepository {
	return &MerchantRepositoryImplement{}
}
