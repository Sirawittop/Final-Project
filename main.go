package main

import (
	"fmt"
	"log"
	"reflect"

	"nurse_shift/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var plans []model.Plans
var plantypes []model.Plantypes
var nurses []model.Nurses
var wards []model.Wards
var hospitals []model.Hospitals

// main function to make list
func makeList(plan model.Plans) (typeplaninplan []int) {
	val := reflect.ValueOf(plan)
	for i := 2; i < val.NumField(); i++ {
		fieldValue := val.Field(i)
		typeplaninplan = append(typeplaninplan, int(fieldValue.Int()))
	}
	return typeplaninplan
}

func main() {
	// Connect to the database
	dsn := "top:1234@tcp(127.0.0.1:8889)/NursePlan?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	// Migrate the schema
	db.Find(&plantypes)
	db.Find(&plans)
	db.Find(&nurses)
	db.Find(&wards)
	db.Find(&hospitals)

	for _, plan := range plans {
		result := makeList(plan)
		fmt.Println(result)
	}

}
