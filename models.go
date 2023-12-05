package main

import (
	"time"

	"gorm.io/gorm"
)

type PlanType struct {
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
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Hospital struct {
	gorm.Model
	HospitalCode string `gorm:"primaryKey"`
	HospitalName string
}

type Ward struct {
	gorm.Model
	WardID       int `gorm:"primaryKey;autoIncrement"`
	Hospital     Hospital
	HospitalCode string
	WardName     string
}

type Nurse struct {
	gorm.Model
	MedicalPID int `gorm:"primaryKey;autoIncrement"`
	Firstname  string
	Lastname   string
}
