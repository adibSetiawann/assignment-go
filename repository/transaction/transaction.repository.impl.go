package repository

import (
	"github.com/adibSetiawann/transaction-api-go/config"
	"github.com/adibSetiawann/transaction-api-go/entity"
	"github.com/adibSetiawann/transaction-api-go/model"
)

type TransactionRepositoryImplement struct {
}

func (*TransactionRepositoryImplement) Create(transaction *entity.Transaction) error {

	var customers []model.CustomerResponse
	var prdocuts []model.ProductResponse

	errorCustomerNotFound := config.DB.Debug().Preload("Gender").First(&customers, "id=?", transaction.CustomerId)
	if errorCustomerNotFound.Error != nil {
		return errorCustomerNotFound.Error
	}

	errorProductNotFound := config.DB.Debug().Preload("Merchant").First(&prdocuts, "id=?", transaction.ProductId)
	if errorProductNotFound.Error != nil {
		return errorProductNotFound.Error
	}

	db := config.DB.Debug().Create(&transaction)
	if db.Error != nil {
		return db.Error
	}

	return nil
}

func (*TransactionRepositoryImplement) Update(invoice string, transaction *model.UpdateTransaction) error {
	var transactionData entity.Transaction

	err := config.DB.Debug().First(&transactionData, "invoice=?", invoice)
	if err.Error != nil {
		return err.Error
	}

	transactionData.ProductId = transaction.ProductId
	transactionData.Qty = transaction.Qty

	errUpdate := config.DB.Debug().Save(&transactionData).Error
	if errUpdate != nil {
		return errUpdate
	}

	return nil
}
func (*TransactionRepositoryImplement) DeleteByInvoice(invoice string) error {
	var transaction entity.Transaction

	err := config.DB.Debug().First(&transaction, "invoice=?", invoice).Error
	if err != nil {
		return err
	}

	errDelete := config.DB.Debug().Delete(&transaction).Error
	if errDelete != nil {
		return err
	}

	return nil
}

func (*TransactionRepositoryImplement) FindByInvoice(invoice string) ([]model.TransactionResponse, error) {

	var transactions []model.TransactionResponse

	err := config.DB.Debug().Preload("Customer").Preload("Product").Preload("Status").First(&transactions, "invoice=?", invoice)
	if err.Error != nil {
		return nil, err.Error
	}

	return transactions, nil
}

func (*TransactionRepositoryImplement) FindAll() ([]model.TransactionResponse, error) {

	var transactions []model.TransactionResponse

	db := config.DB.Debug().Preload("Customer").Preload("Product").Preload("Status").Find(&transactions)
	if db.Error != nil {
		return nil, db.Error
	}

	return transactions, nil
}

func NewTransactionRepository() TransactionRepository {
	return &TransactionRepositoryImplement{}
}
