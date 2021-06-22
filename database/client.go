package database

import (
	"log"
	"rest-go-loja/entity"

	"github.com/jinzhu/gorm"
)

var Connector *gorm.DB

func Connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	log.Println("Connection was successful!!")
	return nil
}

func Migrate(table *entity.Product) {
	Connector.AutoMigrate(&table)
	log.Println("Table migrated")
}
