package main

import (
	"fmt"
	"html/template"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Define your model struct
type Plan_Type struct {
	gorm.Model
	TypeID     int    `gorm:"column:Type_ID;AUTO_INCREMENT"`
	TypeName   string `gorm:"column:Type_Name"`
	Morning    string `gorm:"column:ช"`
	Afternoon  string `gorm:"column:บ"`
	Night      string `gorm:"column:ด"`
	X          string `gorm:"column:x"`
	V          string `gorm:"column:v"`
	Leave      string `gorm:"column:ล"`
	C          string `gorm:"column:c"`
	Part       string `gorm:"column:น"`
	OtM        string `gorm:"column:Otช"`
	OtA        string `gorm:"column:Otบ"`
	OtN        string `gorm:"column:Otด"`
	StatusS    string `gorm:"column:Status_s"`
	StatusOt12 string `gorm:"column:status Ot12"`
	StatusOt8  string `gorm:"column:status Ot8"`
}

func main() {
	// Replace 'username:password@tcp(host:port)/database' with your MySQL connection details
	dsn := "top:1234@tcp(127.0.0.1:8889)/Nurse_Shift_scheduling?charset=utf8mb4&parseTime=True&loc=Local"

	// Open a connection to the database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}

	// AutoMigrate will create the table based on the YourModel struct
	err = db.AutoMigrate(&Plan_Type{})
	if err != nil {
		fmt.Println("Error migrating the database:", err)
		return
	}

	// Create a simple HTTP server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Fetch data from the database
		var plans []Plan_Type
		db.Find(&plans)

		// Render the data in an HTML template
		tmpl, err := template.ParseFiles("templates/index.html")
		if err != nil {
			fmt.Println("Error parsing template:", err)
			return
		}

		err = tmpl.Execute(w, plans)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})

	// Start the HTTP server
	http.ListenAndServe(":5555", nil)
}
