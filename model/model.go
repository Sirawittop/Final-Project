package model

import "gorm.io/gorm"

type Hospitals struct {
	gorm.Model
	ID      uint `gorm:"primaryKey;autoIncrement;not null"`
	Name    string
	Address string
	Phone   string
}

type Login struct {
	gorm.Model
	ID       uint `gorm:"primaryKey;autoIncrement;not null"`
	Username string
	Password string
	RoleID   uint
}

type NurseLevel struct {
	gorm.Model
	ID    uint `gorm:"primaryKey;autoIncrement;not null"`
	Level string
}

type OTs struct {
	gorm.Model
	ID           uint `gorm:"primaryKey;autoIncrement;not null"`
	Name         string
	Hospitals_id uint
	Ward_id      uint
	Morning      uint
	Afternoon    uint
	Night        uint
}

type Plans struct {
	gorm.Model
	ID           int `gorm:"primaryKey;autoIncrement;not null"`
	Hospitals_id uint
	Ward_id      uint
	Nurse_id     uint
	OT_id        uint
	Day1         uint
	Day2         uint
	Day3         uint
	Day4         uint
	Day5         uint
	Day6         uint
	Day7         uint
	Day8         uint
	Day9         uint
	Day10        uint
	Day11        uint
	Day12        uint
	Day13        uint
	Day14        uint
	Day15        uint
	Day16        uint
	Day17        uint
	Day18        uint
	Day19        uint
	Day20        uint
	Day21        uint
	Day22        uint
	Day23        uint
	Day24        uint
	Day25        uint
	Day26        uint
	Day27        uint
	Day28        uint
	Day29        uint
	Day30        uint
	Day31        uint
}

type Plantypes struct {
	gorm.Model
	ID         int `gorm:"primaryKey;autoIncrement;not null"`
	Type       string
	Morning    bool
	Afternoon  bool
	Night      bool
	X          bool
	V          bool
	Nleave     bool
	C          bool
	Part       bool
	Otm        bool
	Ota        bool
	Otn        bool
	StatusS    bool
	StatusOT8  bool
	StatusOT12 bool
}

type Roles struct {
	gorm.Model
	ID   uint `gorm:"primaryKey;autoIncrement;not null"`
	Name string
}

type User struct {
	gorm.Model
	ID              uint `gorm:"primaryKey;autoIncrement;not null"`
	Hospital_id     uint
	Ward_id         uint
	Login_id        uint
	Nurse_level     uint
	Firstname       string
	Lastname        string
	Gender          string
	Phone           string
	Email           string
	Employment_year string
}

type Wards struct {
	gorm.Model
	ID           uint `gorm:"primaryKey;autoIncrement;not null"`
	Hospitals_id uint
	Name         string
}
