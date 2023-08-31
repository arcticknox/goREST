package db

import (
	"fmt"
	"goREST/models"
	"goREST/utils"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBConn *gorm.DB

func InitDB(dbName string) {
	var (
		databaseUser     string = utils.GetConfig("DB_USER")
		databasePassword string = utils.GetConfig("DB_PASSWORD")
		databaseHost     string = utils.GetConfig("DB_HOST")
		databasePort     string = utils.GetConfig("DB_PORT")
		databaseName     string = dbName
	)

	var dataSource string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", databaseUser, databasePassword, databaseHost, databasePort, databaseName)

	var err error
	DBConn, err = gorm.Open(mysql.Open(dataSource), &gorm.Config{})
	if err != nil {
		log.Panic("DB Connect failed.")
	}
	log.Println("DB Connected.")
	DBConn.AutoMigrate(&models.Item{})
}
