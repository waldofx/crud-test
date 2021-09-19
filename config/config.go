package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	//connectionString := "root:@tcp(0.0.0.0:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	connectionString := "https://virtserver.swaggerhub.com/sepulsa/RentABook-API/1.0.0/"

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func InitMigrate() {
	DB.AutoMigrate(&models.User{})
}