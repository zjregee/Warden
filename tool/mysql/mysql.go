package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	dsn := "root@tcp(47.114.99.34:3306)/nursingHouse?charset=utf8mb4&parseTime=True&loc=Local"
	var err error;
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("mysql init error")
	}
}
