package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	dsn := "top:1234@tcp(127.0.0.1:8889)/nurse_scheduler?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}
	//show struct in database models.go hospital
	var Hospital Hospital
	DB.First(&Hospital)
	fmt.Println(Hospital)
}
