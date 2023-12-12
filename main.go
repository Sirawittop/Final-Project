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
			Day1      string
			Day2      string
			Day3      string
			Day4      string
			Day5      string
			Day6      string
			Day7      string
			Day8      string
			Day9      string
			Day10     string
			Day11     string
			Day12     string
			Day13     string
			Day14     string
			Day15     string
			Day16     string
			Day17     string
			Day18     string
			Day19     string
			Day20     string
			Day21     string
			Day22     string
			Day23     string
			Day24     string
			Day25     string
			Day26     string
			Day27     string
			Day28     string
			Day29     string
			Day30     string
			Day31     string
		}

		// Execute raw SQL query
		db.Raw(`
		SELECT 
    		Nurses.id AS Nurse_ID,
    		Nurses.Firstname, 
    		Nurses.Lastname, 
    		Plantypes_Day1.TypeName AS Day1, 
    		Plantypes_Day2.TypeName AS Day2, 
    		Plantypes_Day3.TypeName AS Day3,
    		Plantypes_Day4.TypeName AS Day4,
    		Plantypes_Day5.TypeName AS Day5,
    		Plantypes_Day6.TypeName AS Day6,
    		Plantypes_Day7.TypeName AS Day7,
    		Plantypes_Day8.TypeName AS Day8,
    		Plantypes_Day9.TypeName AS Day9,
    		Plantypes_Day10.TypeName AS Day10,
    		Plantypes_Day11.TypeName AS Day11,
    		Plantypes_Day12.TypeName AS Day12,
    		Plantypes_Day13.TypeName AS Day13,
    		Plantypes_Day14.TypeName AS Day14,
    		Plantypes_Day15.TypeName AS Day15,
    		Plantypes_Day16.TypeName AS Day16,
    		Plantypes_Day17.TypeName AS Day17,
    		Plantypes_Day18.TypeName AS Day18,
    		Plantypes_Day19.TypeName AS Day19,
    		Plantypes_Day20.TypeName AS Day20,
    		Plantypes_Day21.TypeName AS Day21,
    		Plantypes_Day22.TypeName AS Day22,
    		Plantypes_Day23.TypeName AS Day23,
    		Plantypes_Day24.TypeName AS Day24,
    		Plantypes_Day25.TypeName AS Day25,
    		Plantypes_Day26.TypeName AS Day26,
    		Plantypes_Day27.TypeName AS Day27,
    		Plantypes_Day28.TypeName AS Day28,
    		Plantypes_Day29.TypeName AS Day29,
    		Plantypes_Day30.TypeName AS Day30,
    		Plantypes_Day31.TypeName AS Day31
		FROM 
    		Nurses
		JOIN 
		    Plans ON NurseID = Nurses.id
		JOIN 
		    Plantypes AS Plantypes_Day1 ON Plans.Day1 = Plantypes_Day1.id
		JOIN 
		    Plantypes AS Plantypes_Day2 ON Plans.Day2 = Plantypes_Day2.id
		JOIN 
		    Plantypes AS Plantypes_Day3 ON Plans.Day3 = Plantypes_Day3.id
		JOIN 
		    Plantypes AS Plantypes_Day4 ON Plans.Day4 = Plantypes_Day4.id
		JOIN
		    Plantypes AS Plantypes_Day5 ON Plans.Day5 = Plantypes_Day5.id
		JOIN    
		    Plantypes AS Plantypes_Day6 ON Plans.Day6 = Plantypes_Day6.id
		JOIN    
		    Plantypes AS Plantypes_Day7 ON Plans.Day7 = Plantypes_Day7.id
		JOIN
		    Plantypes AS Plantypes_Day8 ON Plans.Day8 = Plantypes_Day8.id
		JOIN    
		    Plantypes AS Plantypes_Day9 ON Plans.Day9 = Plantypes_Day9.id
		JOIN    
		    Plantypes AS Plantypes_Day10 ON Plans.Day10 = Plantypes_Day10.id
		JOIN
		    Plantypes AS Plantypes_Day11 ON Plans.Day11 = Plantypes_Day11.id
		JOIN
		    Plantypes AS Plantypes_Day12 ON Plans.Day12 = Plantypes_Day12.id
		JOIN
		    Plantypes AS Plantypes_Day13 ON Plans.Day13 = Plantypes_Day13.id
		JOIN
		    Plantypes AS Plantypes_Day14 ON Plans.Day14 = Plantypes_Day14.id
		JOIN
		    Plantypes AS Plantypes_Day15 ON Plans.Day15 = Plantypes_Day15.id
		JOIN
		    Plantypes AS Plantypes_Day16 ON Plans.Day16 = Plantypes_Day16.id
		JOIN
		    Plantypes AS Plantypes_Day17 ON Plans.Day17 = Plantypes_Day17.id
		JOIN
		    Plantypes AS Plantypes_Day18 ON Plans.Day18 = Plantypes_Day18.id
		JOIN
		    Plantypes AS Plantypes_Day19 ON Plans.Day19 = Plantypes_Day19.id
		JOIN
		    Plantypes AS Plantypes_Day20 ON Plans.Day20 = Plantypes_Day20.id
		JOIN    
		    Plantypes AS Plantypes_Day21 ON Plans.Day21 = Plantypes_Day21.id
		JOIN
		    Plantypes AS Plantypes_Day22 ON Plans.Day22 = Plantypes_Day22.id
		JOIN
		    Plantypes AS Plantypes_Day23 ON Plans.Day23 = Plantypes_Day23.id
		JOIN
		    Plantypes AS Plantypes_Day24 ON Plans.Day24 = Plantypes_Day24.id
		JOIN
		    Plantypes AS Plantypes_Day25 ON Plans.Day25 = Plantypes_Day25.id
		JOIN
		    Plantypes AS Plantypes_Day26 ON Plans.Day26 = Plantypes_Day26.id
		JOIN    
		    Plantypes AS Plantypes_Day27 ON Plans.Day27 = Plantypes_Day27.id
		JOIN
		    Plantypes AS Plantypes_Day28 ON Plans.Day28 = Plantypes_Day28.id
		JOIN
		    Plantypes AS Plantypes_Day29 ON Plans.Day29 = Plantypes_Day29.id
		JOIN
		    Plantypes AS Plantypes_Day30 ON Plans.Day30 = Plantypes_Day30.id
		JOIN 
		    Plantypes AS Plantypes_Day31 ON Plans.Day31 = Plantypes_Day31.id;

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
