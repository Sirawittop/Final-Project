package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"nurse_shift/model"

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

	// Define a handler to fetch data and render the HTML page
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Query data
		var plans []model.Plans
		db.Find(&plans)

		// Render HTML page with the data
		err := tpl.ExecuteTemplate(w, "index.html", plans)
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
