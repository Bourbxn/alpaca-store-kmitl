package config

import (
	"github.com/Bourbxn/alpaca-store-kmitl/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Database *gorm.DB
var DATABASE_URI string = "root:Boss2546@tcp(localhost:3306)/alpaca?charset=utf8mb4&parseTime=True&loc=Local"

func Connect() error {
    var err error

    Database, err = gorm.Open(mysql.Open(DATABASE_URI), &gorm.Config{
        SkipDefaultTransaction: true,
        PrepareStmt:            true,
    })

    if err != nil {
        panic(err)
    }

    Database.AutoMigrate(&models.User{}, &models.Product{}, &models.ProductOption{})

    return nil
}
