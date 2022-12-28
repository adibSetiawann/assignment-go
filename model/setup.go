package model

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host= localhost user=postgres password=root dbname=assignment-test port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	database := db.AutoMigrate(
		&Customer{},
		&Gender{},
		&Merchant{},
		&Product{},
		&Transaction{},
		&Status{},
	)
	if database != nil {
		fmt.Println("Can't running migration")
	}

	DB = db
}
