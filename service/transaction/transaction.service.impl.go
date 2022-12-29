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
	repo "github.com/adibSetiawann/transaction-api-go/repository/product"
	"github.com/go-playground/validator/v10"
)

type TransactionServiceImpl struct {
	transactionRepo repository.TransactionRepository
	productRepo repo.ProductRepository
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

func (ts *TransactionServiceImpl) Create(request model.CreateTransaction) (model.TransactionResponse, error) {
	var transactionResponse model.TransactionResponse
	
	transaction := entity.Transaction{
		Invoice:    GenerateInvoice(),
		StatusId:   1,
		CustomerId: request.CustomerId,
		ProductId:  request.ProductId,
		Qty:        request.Qty,
	}
	
	errCreate := ts.transactionRepo.Create(&transaction)
	if errCreate != nil {
		log.Println(errCreate.Error())
		return transactionResponse, errCreate
	}
	
	var productData model.ProductRelationResponse
	productId := strconv.Itoa(request.ProductId)
	
	product,_ := ts.productRepo.FindById(productId)

	productData.ID = product[0].ID
	productData.Name =product[0].Name
	productData.Description = product[0].Description
	productData.Price = product[0].Price

	transactionResponse.GrandTotal = product[0].Price * request.Qty
	
	transactionResponse.ID = transaction.ID
	transactionResponse.Invoice = transaction.Invoice
	transactionResponse.StatusId = transaction.StatusId
	transactionResponse.CustomerId = transaction.CustomerId
	transactionResponse.ProductId = transaction.ProductId
	transactionResponse.Product = productData
	transactionResponse.Qty = transaction.Qty
	return transactionResponse, nil
}

func (ts *TransactionServiceImpl) Update(invoice string, request model.UpdateTransaction) (model.TransactionResponse, error) {
	var transactionResponse model.TransactionResponse

	errCreate := ts.transactionRepo.Update(invoice, &request)

	if errCreate != nil {
		log.Println(errCreate.Error())
		return transactionResponse, errCreate
	}

	transactions, _ := ts.transactionRepo.FindByInvoice(invoice)

	var productData model.ProductRelationResponse
	productId := strconv.Itoa(request.ProductId)
	
	product,_ := ts.productRepo.FindById(productId)

	productData.ID = product[0].ID
	productData.Name =product[0].Name
	productData.Description = product[0].Description
	productData.Price = product[0].Price

	transactionResponse.GrandTotal = product[0].Price * request.Qty
	
	transactionResponse.ID = transactions[0].ID
	transactionResponse.Invoice = transactions[0].Invoice
	transactionResponse.Qty = request.Qty
	transactionResponse.ProductId = request.ProductId
	transactionResponse.Product = productData
	transactionResponse.CustomerId = transactions[0].CustomerId
	transactionResponse.Customer = transactions[0].Customer
	transactionResponse.StatusId = transactions[0].StatusId
	transactionResponse.Status = transactions[0].Status
	return transactionResponse, nil
}

func (ts *TransactionServiceImpl) Remove(invoice string) error {
	errDelete := ts.transactionRepo.DeleteByInvoice(invoice)
	if errDelete != nil {
		log.Println(errDelete.Error())
		return errDelete
	}

	return nil
}

func (ts *TransactionServiceImpl) GetByInvoice(id string) ([]model.TransactionResponse, error) {
	merchants, errFind := ts.transactionRepo.FindByInvoice(id)
	if errFind != nil {
		return nil, errFind
	}
	return merchants, nil
}

func (ts *TransactionServiceImpl) GetAllData() ([]model.TransactionResponse, error) {
	transactions, errFind := ts.transactionRepo.FindAll()
	if errFind != nil {
		return nil, errFind
	}
	return transactions, nil
}

func (ts *TransactionServiceImpl) CancelPayment(invoice string) error {
	errorCancelPayment := ts.transactionRepo.CancelPayment(invoice)
	if errorCancelPayment != nil {
		log.Println(errorCancelPayment.Error())
		return errorCancelPayment
	}

	return nil
}

func (ts *TransactionServiceImpl) SuccessPayment(invoice string) error {
	errorCancelPayment := ts.transactionRepo.SuccessPayment(invoice)
	if errorCancelPayment != nil {
		log.Println(errorCancelPayment.Error())
		return errorCancelPayment
	}

	return nil
}

func (ts *TransactionServiceImpl) Validation(userRequest interface{}) error {
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

func NewTransactionService(transactionRepo *repository.TransactionRepository, productRepo *repo.ProductRepository) TransactionService {
	return &TransactionServiceImpl{
		transactionRepo: *transactionRepo,
		productRepo: *productRepo,
	}
}
