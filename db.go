package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func openAndMigrate() *gorm.DB {
	d, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/inventory_test?parseTime=true")
	d.DropTable(&Country{})
	d.AutoMigrate(&Country{})
	d.LogMode(true)

	if err != nil {
		panic(err)
	}
	return d
}
