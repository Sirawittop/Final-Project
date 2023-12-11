package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ... (your existing code)
type plan_types struct {
	gorm.Model
	TypeName   string
	Morning    bool
	Afternoon  bool
	Night      bool
	X          bool
	V          bool
	Nleave     bool
	C          bool
	Part       bool
	OtM        bool
	OtA        bool
	OtN        bool
	StatusS    bool
	StatusOt12 bool
	StatusOt8  bool
}

type hospitals struct {
	gorm.Model
	Hospital_name string
}

type ward_data struct {
	gorm.Model
	HospitalCode hospitals `gorm:"foreignKey:ID"`
	WardName     string
}

type nurse_data struct {
	gorm.Model
	Nurse_ID                  int
	MedicalP_ID               int
	CitizenI_Num              int
	ProfessionalM_license_Num string
	Firstname                 string
	Lastname                  string
	Birthday                  time.Time
	Position                  string
	Ward_ID                   int
	Employment_Year           int
	Hospital_ID               int
	Salary                    int
	Gender                    string
	Phone_Number              string
	Number_Children           int
	MaritalS_ID               int
	Religion                  string
	Email                     string
	Line_ID                   string
	RelativeP_Num             string
	Village                   string
	House_number              string
	Zip_Code                  string
	Province                  string
	District                  string
	Subdistrict               string
}

type NursingPlan struct {
	gorm.Model
	Nures nurse_data `gorm:"foreignKey:ID"`
	Day1  plan_types `gorm:"foreignKey:ID"`
	Day2  plan_types `gorm:"foreignKey:ID"`
	Day3  plan_types `gorm:"foreignKey:ID"`
	Day4  plan_types `gorm:"foreignKey:ID"`
	Day5  plan_types `gorm:"foreignKey:ID"`
	Day6  plan_types `gorm:"foreignKey:ID"`
	Day7  plan_types `gorm:"foreignKey:ID"`
	Day8  plan_types `gorm:"foreignKey:ID"`
	Day9  plan_types `gorm:"foreignKey:ID"`
	Day10 plan_types `gorm:"foreignKey:ID"`
	Day11 plan_types `gorm:"foreignKey:ID"`
	Day12 plan_types `gorm:"foreignKey:ID"`
	Day13 plan_types `gorm:"foreignKey:ID"`
	Day14 plan_types `gorm:"foreignKey:ID"`
	Day15 plan_types `gorm:"foreignKey:ID"`
	Day16 plan_types `gorm:"foreignKey:ID"`
	Day17 plan_types `gorm:"foreignKey:ID"`
	Day18 plan_types `gorm:"foreignKey:ID"`
	Day19 plan_types `gorm:"foreignKey:ID"`
	Day20 plan_types `gorm:"foreignKey:ID"`
	Day21 plan_types `gorm:"foreignKey:ID"`
	Day22 plan_types `gorm:"foreignKey:ID"`
	Day23 plan_types `gorm:"foreignKey:ID"`
	Day24 plan_types `gorm:"foreignKey:ID"`
	Day25 plan_types `gorm:"foreignKey:ID"`
	Day26 plan_types `gorm:"foreignKey:ID"`
	Day27 plan_types `gorm:"foreignKey:ID"`
	Day28 plan_types `gorm:"foreignKey:ID"`
	Day29 plan_types `gorm:"foreignKey:ID"`
	Day30 plan_types `gorm:"foreignKey:ID"`
	Day31 plan_types `gorm:"foreignKey:ID"`
}

func main() {
	// connect to the database
	dsn := "top:1234@tcp(127.0.0.1:8889)/Nurse_Shift_scheduling?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}

	// migrate
	db.AutoMigrate(&NursingPlan{}, &ward_data{})

	// Handle HTTP requests
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	// Query data from the database
	// 	var nursingPlans []NursingPlan
	// 	db.Find(&nursingPlans)

	// 	// Render HTML template with the data
	// 	renderTemplate(w, nursingPlans)
	// })

	// // Start the HTTP server
	// http.ListenAndServe(":5555", nil)
}

// Render HTML template with data
// func renderTemplate(w http.ResponseWriter, nursingPlans []NursingPlan) {
// 	// Read the HTML template from the file
// 	tmpl, err := template.ParseFiles("template.html")
// 	if err != nil {
// 		http.Error(w, "Error parsing template", http.StatusInternalServerError)
// 		return
// 	}

// 	// Execute the template with the data and write to the response
// 	err = tmpl.Execute(w, nursingPlans)
// 	if err != nil {
// 		http.Error(w, "Error executing template", http.StatusInternalServerError)
// 		return
// 	}
// }
