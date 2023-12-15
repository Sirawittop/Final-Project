package main

import (
	"fmt"
	"log"
	"net/http"
	"nurse_shift/model"
	"reflect"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var plans []model.Plans
var planTypes []model.Plantypes
var hospitals []model.Hospitals
var wards []model.Wards
var nurses []model.Nurses

func mapPlanwithPlantype(plans []model.Plans, planTypes []model.Plantypes) map[int]string {
	planNumberMapping := make(map[int]string)
	for _, planType := range planTypes {
		planNumberMapping[planType.ID] = planType.TypeName
	}
	return planNumberMapping
}

func generatePlanWithTypeMapping(plans []model.Plans, planNumberMapping map[int]string) []string {
	var planWithType []string
	for i := 0; i < len(plans); i++ {
		for j := 1; j <= 31; j++ {
			dayField := fmt.Sprintf("Day%d", j)
			field := reflect.ValueOf(plans[i]).FieldByName(dayField)
			planType := int(field.Int())
			planWithType = append(planWithType, planNumberMapping[planType])
		}
	}
	return planWithType
}

func getPlanWithTypeForNurse(nurseIndex int, planWithType []string) []string {
	start := nurseIndex * 31
	end := start + 31
	if end > len(planWithType) {
		end = len(planWithType)
	}
	return planWithType[start:end]
}

func main() {
	// Connect to the database
	dsn := "top:1234@tcp(127.0.0.1:8889)/NursePlan?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	db.Find(&plans)
	db.Find(&planTypes)
	db.Find(&hospitals)
	db.Find(&wards)
	db.Find(&nurses)

	planNumberMapping := mapPlanwithPlantype(plans, planTypes)
	planWithType := generatePlanWithTypeMapping(plans, planNumberMapping)

	// Initialize Gin
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "schedule.html", gin.H{
			"Nurses":          nurses,
			"PlanWithType":    planWithType,
			"GetPlanWithType": getPlanWithTypeForNurse,
		})
	})

	router.Run(":5555")
}
