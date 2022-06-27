package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"Warden/model"
)

var DB *gorm.DB

func init() {
	dsn := "root@tcp(47.114.99.34:3306)/nursingHouse?charset=utf8mb4&parseTime=True&loc=Local"
	var err error;
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("mysql init error")
	}

	DB.AutoMigrate(&model.OldInformation{})
	DB.AutoMigrate(&model.CheckOutInformation{})
	DB.AutoMigrate(&model.OutInformation{})
	DB.AutoMigrate(&model.EmployeeInformation{})
	DB.AutoMigrate(&model.UserInformation{})
}
