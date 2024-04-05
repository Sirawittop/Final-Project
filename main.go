package main

import (
	"fmt"
	"log"
	"nurse_shift/model"
	"reflect"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var plans []model.Plans
var planTypes []model.Plantypes
var hospitals []model.Hospitals
var wards []model.Wards
var nurses []model.User
var OTs []model.OTs

// func calculateOT(nurseIndex int, planWithType []string, OTs []model.OTs, plans []model.Plans) int {
// 	planWithTypeForNurse := getPlanWithTypeForNurse(nurseIndex, planWithType)
// 	otmor := 0
// 	otaft := 0
// 	otnin := 0

// 	// Ot คิดยังไงเมื่อ OT/OT เช่น rด/rช คิดเป็น OT 1 หรือ OT 2
// 	for _, planType := range planWithTypeForNurse {
// 		if planType == "rด/บ" || planType == "x/ด" || planType == "rด/rบ" || planType == "rด/rช" {
// 			otnin++
// 		}
// 		if planType == "rช/บ" || planType == "ด/rช" || planType == "x/ช" || planType == "rช/rบ" || planType == "rด/rช" {
// 			otmor++
// 		}
// 		if planType == "ช/rบ" || planType == "ด/rบ" || planType == "rด/ช" || planType == "x/บ" || planType == "rช/rบ" || planType == "rด/rบ" {
// 			otaft++
// 		}
// 	}
// 	if plans[nurseIndex].OT == 1 {
// 		return (OTs[0].Morning * otmor) + (OTs[0].Afternoon * otaft) + (OTs[0].Night * otnin)
// 	} else if plans[nurseIndex].OT == 2 {
// 		return (OTs[1].Morning * otmor) + (OTs[1].Afternoon * otaft) + (OTs[1].Night * otnin)
// 	} else if plans[nurseIndex].OT == 3 {
// 		return (OTs[2].Morning * otmor) + (OTs[2].Afternoon * otaft) + (OTs[2].Night * otnin)
// 	}
// 	return 0
// }
// func splitPlanTypeName(planTypeName string) string {
// 	parts := strings.Split(planTypeName, "")
// 	if parts[0] == "r" && parts[1] == "ด" && parts[3] == "ช" {
// 		return "red-ด black-ช"
// 	} else if parts[0] == "r" && parts[1] == "ด" && parts[3] == "บ" {
// 		return "red-ด black-บ"
// 	} else if parts[0] == "r" && parts[1] == "ช" && parts[3] == "บ" {
// 		return "red-ช black-บ"
// 	} else if parts[2] == "r" && parts[0] == "ด" && parts[3] == "ช" {
// 		return "black-ด red-ช"
// 	} else if parts[2] == "r" && parts[0] == "ด" && parts[3] == "บ" {
// 		return "black-ด red-บ"
// 	} else if parts[2] == "r" && parts[0] == "ช" && parts[3] == "บ" {
// 		return "black-ช red-บ"
// 	} else if parts[0] == "r" && parts[1] == "ด" && parts[3] == "r" && parts[4] == "ช" {
// 		return "red-ด red-ช"
// 	} else if parts[0] == "r" && parts[1] == "ด" && parts[3] == "r" && parts[4] == "บ" {
// 		return "red-ด red-บ"
// 	} else if parts[0] == "r" && parts[1] == "ช" && parts[3] == "r" && parts[4] == "บ" {
// 		return "red-ช red-บ"
// 	}
// 	return ""
// }

func mapPlanwithPlantype(plans []model.Plans, planTypes []model.Plantypes) map[int]string {
	planNumberMapping := make(map[int]string)
	for _, planType := range planTypes {
		planNumberMapping[int(planType.ID)] = planType.Type
	}
	return planNumberMapping
}

func generatePlanWithTypeMapping(plans []model.Plans, planNumberMapping map[int]string) map[uint][]string {
	result := make(map[uint][]string)
	for _, plan := range plans {
		nurseID := plan.Nurse_id
		var planWithType []string
		for j := 1; j <= 31; j++ {
			dayField := fmt.Sprintf("Day%d", j)
			field := reflect.ValueOf(plan).FieldByName(dayField)
			planType := int(field.Uint())
			planWithType = append(planWithType, planNumberMapping[planType])
		}
		result[nurseID] = planWithType
	}
	return result
}

func calculateNursesShift(planWithType []string) []int {
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
	for _, PlanID := range planWithType {
		if PlanID == "ช" || PlanID == "ช/rบ" || PlanID == "rด/ช" || PlanID == "ช2" {
			mor++
		}
		if PlanID == "บ" || PlanID == "rช/บ" || PlanID == "rด/บ" {
			aft++
		}
		if PlanID == "ด" || PlanID == "ด/rบ" || PlanID == "ด/rช" {
			nin++
		}
		if PlanID == "X" {
			off++
		}
		if PlanID == "C" {
			onc++
		}
		if PlanID == "rช/บ" || PlanID == "ช/rบ" || PlanID == "ด/rบ" || PlanID == "rด/บ" || PlanID == "ด/rช" || PlanID == "rด/ช" || PlanID == "x/ช" || PlanID == "x/บ" || PlanID == "x/ด" || PlanID == "rช/rบ" || PlanID == "rด/rบ" || PlanID == "rด/rช" {
			otv++
		}
		if PlanID == "ล" {
			lev++
		}
		if PlanID == "V" {
			cav++
		}
		if PlanID == "T" {
			cat++
		}
	}
	calnurseshift = append(calnurseshift, nin, mor, aft, off, onc, otv, lev, cav, cat)
	return calnurseshift
}

func mapNurseUser(nurse []model.User) map[uint]string {
	nurseMapping := make(map[uint]string)
	for _, nurseUser := range nurse {
		nurseMapping[nurseUser.ID] = nurseUser.Firstname + " " + nurseUser.Lastname
	}

	return nurseMapping
}

func generatePlanWithUser(plans []model.Plans, planNumberMapping map[int]string, nurseMapping map[uint]string) map[uint]map[string][]string {
	result := make(map[uint]map[string][]string)
	for _, plan := range plans {
		nurseID := plan.Nurse_id
		nurseName := nurseMapping[nurseID]

		var planWithType []string
		for j := 1; j <= 31; j++ {
			dayField := fmt.Sprintf("Day%d", j)
			field := reflect.ValueOf(plan).FieldByName(dayField)
			planType := int(field.Uint())
			planWithType = append(planWithType, planNumberMapping[planType])
		}

		if _, ok := result[nurseID]; !ok {
			result[nurseID] = make(map[string][]string)
		}
		result[nurseID][nurseName] = planWithType
	}
	return result
}

// make function to map nurse with plan

func main() {
	dsn := "sirawittop:1234@tcp(127.0.0.1:3306)/NursePlans?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	db.Find(&plans)
	db.Find(&planTypes)
	db.Find(&hospitals)
	db.Find(&wards)
	db.Find(&nurses)
	db.Find(&OTs)

	planNumberMapping := mapPlanwithPlantype(plans, planTypes)

	planWithType := generatePlanWithTypeMapping(plans, planNumberMapping)

	fmt.Println("planwithtype", planWithType)

	nurseMapping := mapNurseUser(nurses)

	planWithUser := generatePlanWithUser(plans, planNumberMapping, nurseMapping)

	fmt.Println("planWithUser", planWithUser)

	// nursename := mapNurseUser(nurses)

	// Map nurses with their IDs
	// nursename := mapNurseUser(nurses)

	// // Print the mapped nurses
	// fmt.Println("Nurse Name Map:")
	// for id, name := range nursename {
	//     fmt.Printf("ID: %d, Name: %s\n", id, name)
	// }

	// if planWithType[1][0] == "ช" {
	// 	fmt.Println("yes")
	// }

	// var allNurseShifts [][]int
	// for _, plan := range plans {
	// 	nurseshiftcal := calculateNursesShift(planWithType[plan.Nurse_id])
	// 	allNurseShifts = append(allNurseShifts, nurseshiftcal)
	// }
	// fmt.Println("allNurseShifts", allNurseShifts)

	// fmt.Println("planwithtype", planWithType)
	// var nurseOT []int
	// for i := 0; i < len(plans); i++ {
	// 	nurseOT = append(nurseOT, calculateOT(i, planWithType, OTs, plans))
	// }

	// tmpl := template.Must(template.New("schedule.html").Funcs(template.FuncMap{
	// 	"splitPlanTypeName": splitPlanTypeName,
	// }).ParseFiles("templates/schedule.html"))

	// router := gin.Default()
	// router.SetHTMLTemplate(tmpl)

	// router.GET("/", func(c *gin.Context) {
	// 	c.HTML(http.StatusOK, "schedule.html", gin.H{
	// 		"Nurses":          nurses,
	// 		"PlanWithType":    planWithType,
	// 		"GetPlanWithType": getPlanWithTypeForNurse,
	// 		"CalNurseShifts":  allNurseShifts,
	// 		"nurseOT":         nurseOT,
	// 	})
	// 	db.Find(&plans)
	// 	db.Find(&planTypes)
	// 	db.Find(&hospitals)
	// 	db.Find(&wards)
	// 	db.Find(&nurses)
	// 	db.Find(&OTs)

	// 	planNumberMapping := mapPlanwithPlantype(plans, planTypes)
	// 	planWithType = generatePlanWithTypeMapping(plans, planNumberMapping)
	// 	allNurseShifts = [][]int{}
	// 	for i := 0; i < len(plans); i++ {
	// 		nurseshiftcal := calculateNursesShift(i, planWithType)
	// 		allNurseShifts = append(allNurseShifts, nurseshiftcal)
	// 	}

	// 	nurseOT = []int{}
	// 	for i := 0; i < len(plans); i++ {
	// 		nurseOT = append(nurseOT, calculateOT(i, planWithType, OTs, plans))
	// 	}
	// })

	// router.Run(":5555")
}
