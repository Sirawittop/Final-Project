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
		if planType == "ช" || planType == "T" || planType == "ช/rบ" || planType == "rด/ช" || planType == "ช2" {
			mor++
		}
		if planType == "บ" || planType == "rช/บ" || planType == "rด/บ" {
			aft++
		}
		if planType == "ด" || planType == "ด/rบ" || planType == "ด/rช" {
			nin++
		}
		if planType == "X" {
			off++
		}
		if planType == "C" {
			onc++
		}
		if planType == "rช/บ" || planType == "ช/rบ" || planType == "ด/rบ" || planType == "rด/บ" || planType == "ด/rช" || planType == "rด/ช" || planType == "x/ช" || planType == "x/บ" || planType == "x/ด" || planType == "rช/rบ" || planType == "rด/rบ" || planType == "rด/rช" {
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

func splitPlanTypeName(planTypeName string) string {
	parts := strings.Split(planTypeName, "")
	if parts[0] == "r" && parts[1] == "ด" && parts[3] == "ช" {
		return "red-ด black-ช"
	} else if parts[0] == "r" && parts[1] == "ด" && parts[3] == "บ" {
		return "red-ด black-บ"
	} else if parts[0] == "r" && parts[1] == "ช" && parts[3] == "บ" {
		return "red-ช black-บ"
	} else if parts[2] == "r" && parts[0] == "ด" && parts[3] == "ช" {
		return "black-ด red-ช"
	} else if parts[2] == "r" && parts[0] == "ด" && parts[3] == "บ" {
		return "black-ด red-บ"
	} else if parts[2] == "r" && parts[0] == "ช" && parts[3] == "บ" {
		return "black-ช red-บ"
	} else if parts[0] == "r" && parts[1] == "ด" && parts[3] == "r" && parts[4] == "ช" {
		return "red-ด red-ช"
	} else if parts[0] == "r" && parts[1] == "ด" && parts[3] == "r" && parts[4] == "บ" {
		return "red-ด red-บ"
	} else if parts[0] == "r" && parts[1] == "ช" && parts[3] == "r" && parts[4] == "บ" {
		return "red-ช red-บ"
	}
	return ""
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

	tmpl := template.Must(template.New("schedule.html").Funcs(template.FuncMap{
		"splitPlanTypeName": splitPlanTypeName,
	}).ParseFiles("templates/schedule.html"))
	router := gin.Default()
	router.SetHTMLTemplate(tmpl)
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
