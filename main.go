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
	Plan_Type_ID int
	TypeName     string
	Morning      bool
	Afternoon    bool
	Night        bool
	X            bool
	V            bool
	Nleave       bool
	C            bool
	Part         bool
	OtM          bool
	OtA          bool
	OtN          bool
	StatusS      bool
	StatusOt12   bool
	StatusOt8    bool
}

type hospitals struct {
	gorm.Model
	HospitalCode  string
	hospital_name string
}

type ward_data struct {
	gorm.Model
	Ward_ID      int
	HospitalCode string
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
	gender                    string
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

type plan struct {
	gorm.Model
	plan_types plan_types
	hospitals  hospitals
	ward_data  ward_data
	nurse_data nurse_data
	day1       plan_types
	day2       plan_types
	day3       plan_types
	day4       plan_types
	day5       plan_types
	day6       plan_types
	day7       plan_types
	day8       plan_types
	day9       plan_types
	day10      plan_types
	day11      plan_types
	day12      plan_types
	day13      plan_types
	day14      plan_types
	day15      plan_types
	day16      plan_types
	day17      plan_types
	day18      plan_types
	day19      plan_types
	day20      plan_types
	day21      plan_types
	day22      plan_types
	day23      plan_types
	day24      plan_types
	day25      plan_types
	day26      plan_types
	day27      plan_types
	day28      plan_types
	day29      plan_types
	day30      plan_types
	day31      plan_types
}

func main() {
	// connect to the database
	dsn := "top:1234@tcp(127.0.0.1:8889)/ORM_Data?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}

	// migrate
	db.AutoMigrate(&plan{})

	// query
	var plan_types []plan
	db.Find(&plan_types)
	fmt.Println(plan_types)

}
