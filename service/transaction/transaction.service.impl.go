package service

import (
	"errors"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/adibSetiawann/transaction-api-go/entity"
	"github.com/adibSetiawann/transaction-api-go/model"
	repository "github.com/adibSetiawann/transaction-api-go/repository/transaction"
	"github.com/go-playground/validator/v10"
)

type TransactionServiceImpl struct {
	transactionRepo repository.TransactionRepository
}

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}

func GenerateInvoice() string {
	rndmString := strings.ToUpper(String(10))
	_, month, day := time.Now().Date()

	var str strings.Builder
	str.WriteString("TRSC-")
	str.WriteString(strconv.Itoa(int(month)))
	str.WriteString(strconv.Itoa(day))
	str.WriteString(rndmString)

	return str.String()
}

func (ps *TransactionServiceImpl) Create(request model.CreateTransaction) (model.TransactionResponse, error) {
	var transactionResponse model.TransactionResponse

	transaction := entity.Transaction{
		Invoice:    GenerateInvoice(),
		StatusId:   1,
		CustomerId: request.CustomerId,
		ProductId:  request.ProductId,
		Qty:        request.Qty,
	}

	errCreate := ps.transactionRepo.Create(&transaction)
	if errCreate != nil {
		log.Println(errCreate.Error())
		return transactionResponse, errCreate
	}

	transactionResponse.ID = transaction.ID
	transactionResponse.Invoice = transaction.Invoice
	transactionResponse.StatusId = transaction.StatusId
	transactionResponse.CustomerId = transaction.CustomerId
	transactionResponse.ProductId = transaction.ProductId
	transactionResponse.Qty = transaction.Qty
	return transactionResponse, nil
}

func (ms *TransactionServiceImpl) Update(invoice string, request model.UpdateTransaction) (model.TransactionResponse, error) {
	var transactionResponse model.TransactionResponse

	errCreate := ms.transactionRepo.Update(invoice, &request)

	if errCreate != nil {
		log.Println(errCreate.Error())
		return transactionResponse, errCreate
	}

	transactions, _ := ms.transactionRepo.FindByInvoice(invoice)

	transactionResponse.ID = transactions[0].ID
	transactionResponse.Invoice = transactions[0].Invoice
	transactionResponse.Qty = request.Qty
	transactionResponse.ProductId = request.ProductId
	transactionResponse.Product = transactions[0].Product
	transactionResponse.CustomerId = transactions[0].CustomerId
	transactionResponse.Customer = transactions[0].Customer
	transactionResponse.StatusId = transactions[0].StatusId
	transactionResponse.Status = transactions[0].Status
	return transactionResponse, nil
}

func (ms *TransactionServiceImpl) Remove(invoice string) error {
	errDelete := ms.transactionRepo.DeleteByInvoice(invoice)
	if errDelete != nil {
		log.Println(errDelete.Error())
		return errDelete
	}

	return nil
}

func (ms *TransactionServiceImpl) GetByInvoice(id string) ([]model.TransactionResponse, error) {
	merchants, errFind := ms.transactionRepo.FindByInvoice(id)
	if errFind != nil {
		return nil, errFind
	}
	return merchants, nil
}

func (ps *TransactionServiceImpl) GetAllData() ([]model.TransactionResponse, error) {
	transactions, errFind := ps.transactionRepo.FindAll()
	if errFind != nil {
		return nil, errFind
	}
	return transactions, nil
}

func (ms *TransactionServiceImpl) Validation(userRequest interface{}) error {
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

func NewTransactionService(transactionRepo *repository.TransactionRepository) TransactionService {
	return &TransactionServiceImpl{
		transactionRepo: *transactionRepo,
	}
}
