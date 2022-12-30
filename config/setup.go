package config

import (
	"fmt"

	"github.com/adibSetiawann/transaction-api-go/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dbHost := GetEnvVariable("DB_HOST")
	dbUser := GetEnvVariable("DB_USER")
	dbPass := GetEnvVariable("DB_PASSWORD")
	dbName := GetEnvVariable("DB_NAME")
	dbPort := GetEnvVariable("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		dbHost,
		dbUser,
		dbPass,
		dbName,
		dbPort,
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
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
