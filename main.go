package main

import (
	"fmt"
	"nurse_shift/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// connect to the database
	dsn := "top:1234@tcp(127.0.0.1:8889)/NursePlan?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}

	//query data
	var plantype []model.Plans
	db.First(&plantype)
	fmt.Println(plantype)
}
