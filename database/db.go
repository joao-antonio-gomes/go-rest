package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
	err error
)

func ConnectDatabase() {
	connectionString := "host=localhost user=postgres password=postgres dbname=go-rest port=5433 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Panic("Erro ao conectar com o banco de dados")
	}
}