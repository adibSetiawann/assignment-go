package config

import (
	"fmt"

	"github.com/adibSetiawann/transaction-api-go/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host= localhost user=postgres password=root dbname=assignment-test port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	database := db.AutoMigrate(
		&entity.Customer{},
		&entity.Gender{},
		&entity.Merchant{},
		&entity.Product{},
		&entity.Transaction{},
		&entity.Status{},
	)
	if database != nil {
		fmt.Println("Can't running migration")
	}

	DB = db
}
