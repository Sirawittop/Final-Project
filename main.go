package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"nurse_shift/model"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var plans []model.Plans
var planTypes []model.Plantypes
var hospitals []model.Hospitals
var wards []model.Wards
var nurses []model.Nurses

func mapPlanTypeToClass(planTypeName string) string {
	switch {
	case strings.Contains(planTypeName, "r") && strings.Contains(planTypeName, "ช"):
		return "red-ช"
	case strings.Contains(planTypeName, "r") && strings.Contains(planTypeName, "บ"):
		return "red-บ"
	case strings.Contains(planTypeName, "r") && strings.Contains(planTypeName, "ด"):
		return "red-ด"
	case planTypeName == "x/ช" || planTypeName == "x/บ" || planTypeName == "x/ด":
		return "red"
	default:
		return ""
	}
}

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

func calculateNursesShift(nurseIndex int, planWithType []string) []int {
	planWithTypeForNurse := getPlanWithTypeForNurse(nurseIndex, planWithType)
	var calnurseshift []int
	mor := 0
	aft := 0
	nin := 0
	off := 0
	onc := 0
	otv := 0
	lev := 0
	cav := 0
	cat := 0
	for _, planType := range planWithTypeForNurse {
		if planType == "ช" || planType == "T" || planType == "ช/rบ" || planType == "ช/rด" || planType == "ช2" {
			mor++
		}
		if planType == "บ" || planType == "rช/บ" || planType == "บ/rด" {
			aft++
		}
		if planType == "ด" || planType == "rบ/ด" || planType == "rช/ด" {
			nin++
		}
		if planType == "X" {
			off++
		}
		if planType == "C" {
			onc++
		}
		if planType == "rช/บ" || planType == "ช/rบ" || planType == "rบ/ด" || planType == "บ/rด" || planType == "rช/ด" || planType == "ช/rด" || planType == "x/ช" || planType == "x/บ" || planType == "x/ด" || planType == "rช/rบ" || planType == "rบ/rด" || planType == "rช/rด" {
			otv++
		}
		if planType == "ล" {
			lev++
		}
		if planType == "v" {
			cav++
		}
		if planType == "T" {
			cat++
		}
	}
	calnurseshift = append(calnurseshift, nin, mor, aft, off, onc, otv, lev, cav, cat)
	return calnurseshift
}

func main() {
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
	var allNurseShifts [][]int
	for i := 0; i < len(plans); i++ {
		nurseshiftcal := calculateNursesShift(i, planWithType)
		allNurseShifts = append(allNurseShifts, nurseshiftcal)
	}

	router := gin.Default()
	router.SetFuncMap(template.FuncMap{
		"mapPlanTypeToClass": mapPlanTypeToClass,
	})
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "schedule.html", gin.H{
			"Nurses":          nurses,
			"PlanWithType":    planWithType,
			"GetPlanWithType": getPlanWithTypeForNurse,
			"CalNurseShifts":  allNurseShifts,
		})
	})
	router.Run(":5555")
}
