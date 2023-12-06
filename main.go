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
	hospital_name string
}

type ward_data struct {
	gorm.Model
	HospitalCode string
	WardName     string
}

type nurse_data struct {
	gorm.Model
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

func main() {
	// Replace these with your actual database credentials
	dsn := "top:1234@tcp(127.0.0.1:8889)/ORM_Data?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}

	// Create the table
	db.AutoMigrate(&plan_types{}, &hospitals{}, &ward_data{}, &nurse_data{})
}
