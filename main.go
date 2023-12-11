package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var tpl *template.Template

func init() {
	// Parse HTML templates
	tpl = template.Must(template.ParseGlob("templates/*.html"))
}

func main() {
	// Connect to the database
	dsn := "top:1234@tcp(127.0.0.1:8889)/NursePlan?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	// Close the database connection when the main function completes
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Error getting underlying database connection:", err)
	}
	defer sqlDB.Close()

	// Define a handler to fetch joined data and render the HTML page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Query joined data
		var plansWithNurseInfo []struct {
			Nurse_ID  int
			Firstname string
			Lastname  string
			Day1      int
			Day2      int
			Day3      int
			Day4      int
			Day5      int
			Day6      int
			Day7      int
			Day8      int
			Day9      int
			Day10     int
			Day11     int
			Day12     int
			Day13     int
			Day14     int
			Day15     int
			Day16     int
			Day17     int
			Day18     int
			Day19     int
			Day20     int
			Day21     int
			Day22     int
			Day23     int
			Day24     int
			Day25     int
			Day26     int
			Day27     int
			Day28     int
			Day29     int
			Day30     int
			Day31     int
		}

		// Execute raw SQL query
		db.Raw(`
		SELECT Nurses.id AS Nurse_ID, Nurses.Firstname, Nurses.Lastname,
		Plans.Day1, Plans.Day2, Plans.Day3, Plans.Day4, Plans.Day5,
		Plans.Day6, Plans.Day7, Plans.Day8, Plans.Day9, Plans.Day10,
		Plans.Day11, Plans.Day12, Plans.Day13, Plans.Day14, Plans.Day15,
		Plans.Day16, Plans.Day17, Plans.Day18, Plans.Day19, Plans.Day20,
		Plans.Day21, Plans.Day22, Plans.Day23, Plans.Day24, Plans.Day25,
		Plans.Day26, Plans.Day27, Plans.Day28, Plans.Day29, Plans.Day30,
		Plans.Day31
		FROM Nurses
		LEFT JOIN Plans ON Nurses.ID = Plans.NurseID;
		`).Scan(&plansWithNurseInfo)

		// Render HTML page with the joined data
		err := tpl.ExecuteTemplate(w, "index.html", plansWithNurseInfo)
		if err != nil {
			log.Println("Error executing template:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	})

	// Start the web server
	port := 5555
	fmt.Printf("Server is running on http://localhost:%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
