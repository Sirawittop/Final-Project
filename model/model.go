package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Plantypes struct {
	ID         int
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

type Hospitals struct {
	ID      int
	Name    string
	Address string
	Phone   string
}

type Wards struct {
	Id         int
	Name       string
	HospitalID int
}

type Nurses struct {
	ID                        int
	MedicalP_ID               int
	CitizenI_Num              int
	ProfessionalM_license_Num string
	Firstname                 string
	Lastname                  string
	Birthday                  time.Time
	Position                  string
	WardID                    int
	Employment_Year           int
	HospitalID                int
	Salary                    decimal.Decimal
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

type Plans struct {
	Id         int
	HospitalID int
	WardID     int
	Month      int
	NurseID    int
	Day1       int
	Day2       int
	Day3       int
	Day4       int
	Day5       int
	Day6       int
	Day7       int
	Day8       int
	Day9       int
	Day10      int
	Day11      int
	Day12      int
	Day13      int
	Day14      int
	Day15      int
	Day16      int
	Day17      int
	Day18      int
	Day19      int
	Day20      int
	Day21      int
	Day22      int
	Day23      int
	Day24      int
	Day25      int
	Day26      int
	Day27      int
	Day28      int
	Day29      int
	Day30      int
	Day31      int
}
