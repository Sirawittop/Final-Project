package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/shopspring/decimal"
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
			Morning   decimal.Decimal
			Afternoon decimal.Decimal
			Night     decimal.Decimal
			Dayoff    decimal.Decimal
			Oncall    decimal.Decimal
			Ot        decimal.Decimal
		}

		// Execute raw SQL query
		db.Raw(`
		SELECT
		Nursesinfo.Nurse_ID,
		Nursesinfo.Firstname,
		Nursesinfo.Lastname,
		Nursesinfo.Day1,
		Nursesinfo.Day2,
		Nursesinfo.Day3,
		Nursesinfo.Day4,
		Nursesinfo.Day5,
		Nursesinfo.Day6,
		Nursesinfo.Day7,
		Nursesinfo.Day8,
		Nursesinfo.Day9,
		Nursesinfo.Day10,
		Nursesinfo.Day11,
		Nursesinfo.Day12,
		Nursesinfo.Day13,
		Nursesinfo.Day14,
		Nursesinfo.Day15,
		Nursesinfo.Day16,
		Nursesinfo.Day17,
		Nursesinfo.Day18,
		Nursesinfo.Day19,
		Nursesinfo.Day20,
		Nursesinfo.Day21,
		Nursesinfo.Day22,
		Nursesinfo.Day23,
		Nursesinfo.Day24,
		Nursesinfo.Day25,
		Nursesinfo.Day26,
		Nursesinfo.Day27,
		Nursesinfo.Day28,
		Nursesinfo.Day29,
		Nursesinfo.Day30,
		Nursesinfo.Day31,
		calinfo.Morning,
		calinfo.Afternoon,
		calinfo.Night,
		calinfo.Dayoff,
		calinfo.Oncall,
		calinfo.Ot
	FROM
		(
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
				JOIN Plans ON NurseID = Nurses.id
				JOIN Plantypes AS Plantypes_Day1 ON Plans.Day1 = Plantypes_Day1.id
				JOIN Plantypes AS Plantypes_Day2 ON Plans.Day2 = Plantypes_Day2.id
				JOIN Plantypes AS Plantypes_Day3 ON Plans.Day3 = Plantypes_Day3.id
				JOIN Plantypes AS Plantypes_Day4 ON Plans.Day4 = Plantypes_Day4.id
				JOIN Plantypes AS Plantypes_Day5 ON Plans.Day5 = Plantypes_Day5.id
				JOIN Plantypes AS Plantypes_Day6 ON Plans.Day6 = Plantypes_Day6.id
				JOIN Plantypes AS Plantypes_Day7 ON Plans.Day7 = Plantypes_Day7.id
				JOIN Plantypes AS Plantypes_Day8 ON Plans.Day8 = Plantypes_Day8.id
				JOIN Plantypes AS Plantypes_Day9 ON Plans.Day9 = Plantypes_Day9.id
				JOIN Plantypes AS Plantypes_Day10 ON Plans.Day10 = Plantypes_Day10.id
				JOIN Plantypes AS Plantypes_Day11 ON Plans.Day11 = Plantypes_Day11.id
				JOIN Plantypes AS Plantypes_Day12 ON Plans.Day12 = Plantypes_Day12.id
				JOIN Plantypes AS Plantypes_Day13 ON Plans.Day13 = Plantypes_Day13.id
				JOIN Plantypes AS Plantypes_Day14 ON Plans.Day14 = Plantypes_Day14.id
				JOIN Plantypes AS Plantypes_Day15 ON Plans.Day15 = Plantypes_Day15.id
				JOIN Plantypes AS Plantypes_Day16 ON Plans.Day16 = Plantypes_Day16.id
				JOIN Plantypes AS Plantypes_Day17 ON Plans.Day17 = Plantypes_Day17.id
				JOIN Plantypes AS Plantypes_Day18 ON Plans.Day18 = Plantypes_Day18.id
				JOIN Plantypes AS Plantypes_Day19 ON Plans.Day19 = Plantypes_Day19.id
				JOIN Plantypes AS Plantypes_Day20 ON Plans.Day20 = Plantypes_Day20.id
				JOIN Plantypes AS Plantypes_Day21 ON Plans.Day21 = Plantypes_Day21.id
				JOIN Plantypes AS Plantypes_Day22 ON Plans.Day22 = Plantypes_Day22.id
				JOIN Plantypes AS Plantypes_Day23 ON Plans.Day23 = Plantypes_Day23.id
				JOIN Plantypes AS Plantypes_Day24 ON Plans.Day24 = Plantypes_Day24.id
				JOIN Plantypes AS Plantypes_Day25 ON Plans.Day25 = Plantypes_Day25.id
				JOIN Plantypes AS Plantypes_Day26 ON Plans.Day26 = Plantypes_Day26.id
				JOIN Plantypes AS Plantypes_Day27 ON Plans.Day27 = Plantypes_Day27.id
				JOIN Plantypes AS Plantypes_Day28 ON Plans.Day28 = Plantypes_Day28.id
				JOIN Plantypes AS Plantypes_Day29 ON Plans.Day29 = Plantypes_Day29.id
				JOIN Plantypes AS Plantypes_Day30 ON Plans.Day30 = Plantypes_Day30.id
				JOIN Plantypes AS Plantypes_Day31 ON Plans.Day31 = Plantypes_Day31.id
		) Nursesinfo
		JOIN (
			SELECT
				a.NurseID,
				a.Morning,
				b.Afternoon,
				c.Night,
				d.Dayoff,
				e.Oncall,
				f.Ot
			From
				(
					SELECT
						NurseID,
						SUM(
							CASE
								WHEN Day1 IN (1, 6, 11, 15, 19) THEN 1
								ELSE 0
							END + CASE
								WHEN Day2 IN (1, 6, 11, 15, 19) THEN 1
								ELSE 0
							END + CASE
								WHEN Day3 IN (1, 6, 11, 15, 19) THEN 1
								ELSE 0
							END + CASE
								WHEN Day4 IN (1, 6, 11, 15, 19) THEN 1
								ELSE 0
							END + CASE
								WHEN Day5 IN (1, 6, 11, 15, 19) THEN 1
								ELSE 0
							END + CASE
								WHEN Day6 IN (1, 6, 11, 15, 19) THEN 1
								ELSE 0
							END + CASE
								WHEN Day7 IN (1, 6, 11, 15, 19) THEN 1
								ELSE 0
							END + CASE
								WHEN Day8 IN (1, 6, 11, 15, 19) THEN 1
								ELSE 0
							END + CASE
								WHEN Day9 IN (1, 6, 11, 15, 19) THEN 1
								ELSE 0
							END + CASE
								WHEN Day10 IN (1, 6, 11, 15, 19) THEN 1
								ELSE 0
							END + CASE
								WHEN Day11 IN (1, 6, 11, 15, 19) THEN 1
								ELSE 0
							END + CASE
								WHEN Day12 IN (1, 6, 11, 15, 19) THEN 1
								ELSE 0
							END + CASE
								WHEN Day13 IN (1, 6, 11, 15, 19) THEN 1
								ELSE 0
							END + CASE
								WHEN Day14 IN (1, 6, 11, 15, 19) THEN 1
								ELSE 0
							END + CASE
								WHEN Day15 IN (1, 6, 11, 15, 19) THEN 1
								ELSE 0
							END + CASE
								WHEN Day16 IN (1, 6, 11, 15, 19) THEN 1
								ELSE 0
							END + CASE
								WHEN Day17 IN (1, 6, 11, 15, 19) THEN 1
								ELSE 0
							END + CASE
								WHEN Day18 IN (1, 6, 11, 15, 19) THEN 1
								ELSE 0
							END + CASE
								WHEN Day19 IN (1, 6, 11, 15, 19) THEN 1
								ELSE 0
							END + CASE
								WHEN Day20 IN (1, 6, 11, 15, 19) THEN 1
								ELSE 0
							END + CASE
								WHEN Day21 IN (1, 6, 11, 15, 19) THEN 1
								ELSE 0
							END + CASE
								WHEN Day22 IN (1, 6, 11, 15, 19) THEN 1
								ELSE 0
							END + CASE
								WHEN Day23 IN (1, 6, 11, 15, 19) THEN 1
								ELSE 0
							END + CASE
								WHEN Day24 IN (1, 6, 11, 15, 19) THEN 1
								ELSE 0
							END + CASE
								WHEN Day25 IN (1, 6, 11, 15, 19) THEN 1
								ELSE 0
							END + CASE
								WHEN Day26 IN (1, 6, 11, 15, 19) THEN 1
								ELSE 0
							END + CASE
								WHEN Day27 IN (1, 6, 11, 15, 19) THEN 1
								ELSE 0
							END + CASE
								WHEN Day28 IN (1, 6, 11, 15, 19) THEN 1
								ELSE 0
							END + CASE
								WHEN Day29 IN (1, 6, 11, 15, 19) THEN 1
								ELSE 0
							END + CASE
								WHEN Day30 IN (1, 6, 11, 15, 19) THEN 1
								ELSE 0
							END + CASE
								WHEN Day31 IN (1, 6, 11, 15, 19) THEN 1
								ELSE 0
							END
						) AS 'Morning'
					FROM
						plans
					GROUP BY
						NurseID
				) a
				JOIN (
					SELECT
						NurseID,
						SUM(
							CASE
								WHEN Day1 IN (2, 10, 13, 20) THEN 1
								ELSE 0
							END + CASE
								WHEN Day2 IN (2, 10, 13, 20) THEN 1
								ELSE 0
							END + CASE
								WHEN Day3 IN (2, 10, 13, 20) THEN 1
								ELSE 0
							END + CASE
								WHEN Day4 IN (2, 10, 13, 20) THEN 1
								ELSE 0
							END + CASE
								WHEN Day5 IN (2, 10, 13, 20) THEN 1
								ELSE 0
							END + CASE
								WHEN Day6 IN (2, 10, 13, 20) THEN 1
								ELSE 0
							END + CASE
								WHEN Day7 IN (2, 10, 13, 20) THEN 1
								ELSE 0
							END + CASE
								WHEN Day8 IN (2, 10, 13, 20) THEN 1
								ELSE 0
							END + CASE
								WHEN Day9 IN (2, 10, 13, 20) THEN 1
								ELSE 0
							END + CASE
								WHEN Day10 IN (2, 10, 13, 20) THEN 1
								ELSE 0
							END + CASE
								WHEN Day11 IN (2, 10, 13, 20) THEN 1
								ELSE 0
							END + CASE
								WHEN Day12 IN (2, 10, 13, 20) THEN 1
								ELSE 0
							END + CASE
								WHEN Day13 IN (2, 10, 13, 20) THEN 1
								ELSE 0
							END + CASE
								WHEN Day14 IN (2, 10, 13, 20) THEN 1
								ELSE 0
							END + CASE
								WHEN Day15 IN (2, 10, 13, 20) THEN 1
								ELSE 0
							END + CASE
								WHEN Day16 IN (2, 10, 13, 20) THEN 1
								ELSE 0
							END + CASE
								WHEN Day17 IN (2, 10, 13, 20) THEN 1
								ELSE 0
							END + CASE
								WHEN Day18 IN (2, 10, 13, 20) THEN 1
								ELSE 0
							END + CASE
								WHEN Day19 IN (2, 10, 13, 20) THEN 1
								ELSE 0
							END + CASE
								WHEN Day20 IN (2, 10, 13, 20) THEN 1
								ELSE 0
							END + CASE
								WHEN Day21 IN (2, 10, 13, 20) THEN 1
								ELSE 0
							END + CASE
								WHEN Day22 IN (2, 10, 13, 20) THEN 1
								ELSE 0
							END + CASE
								WHEN Day23 IN (2, 10, 13, 20) THEN 1
								ELSE 0
							END + CASE
								WHEN Day24 IN (2, 10, 13, 20) THEN 1
								ELSE 0
							END + CASE
								WHEN Day25 IN (2, 10, 13, 20) THEN 1
								ELSE 0
							END + CASE
								WHEN Day26 IN (2, 10, 13, 20) THEN 1
								ELSE 0
							END + CASE
								WHEN Day27 IN (2, 10, 13, 20) THEN 1
								ELSE 0
							END + CASE
								WHEN Day28 IN (2, 10, 13, 20) THEN 1
								ELSE 0
							END + CASE
								WHEN Day29 IN (2, 10, 13, 20) THEN 1
								ELSE 0
							END + CASE
								WHEN Day30 IN (2, 10, 13, 20) THEN 1
								ELSE 0
							END + CASE
								WHEN Day31 IN (2, 10, 13, 20) THEN 1
								ELSE 0
							END
						) AS 'Afternoon'
					FROM
						plans
					GROUP BY
						NurseID
				) b
				JOIN (
					SELECT
						nurseid As NurseID,
						SUM(
							CASE
								WHEN Day1 IN (3, 12, 14) THEN 1
								ELSE 0
							END + CASE
								WHEN Day2 IN (3, 12, 14) THEN 1
								ELSE 0
							END + CASE
								WHEN Day3 IN (3, 12, 14) THEN 1
								ELSE 0
							END + CASE
								WHEN Day4 IN (3, 12, 14) THEN 1
								ELSE 0
							END + CASE
								WHEN Day5 IN (3, 12, 14) THEN 1
								ELSE 0
							END + CASE
								WHEN Day6 IN (3, 12, 14) THEN 1
								ELSE 0
							END + CASE
								WHEN Day7 IN (3, 12, 14) THEN 1
								ELSE 0
							END + CASE
								WHEN Day8 IN (3, 12, 14) THEN 1
								ELSE 0
							END + CASE
								WHEN Day9 IN (3, 12, 14) THEN 1
								ELSE 0
							END + CASE
								WHEN Day10 IN (3, 12, 14) THEN 1
								ELSE 0
							END + CASE
								WHEN Day11 IN (3, 12, 14) THEN 1
								ELSE 0
							END + CASE
								WHEN Day12 IN (3, 12, 14) THEN 1
								ELSE 0
							END + CASE
								WHEN Day13 IN (3, 12, 14) THEN 1
								ELSE 0
							END + CASE
								WHEN Day14 IN (3, 12, 14) THEN 1
								ELSE 0
							END + CASE
								WHEN Day15 IN (3, 12, 14) THEN 1
								ELSE 0
							END + CASE
								WHEN Day16 IN (3, 12, 14) THEN 1
								ELSE 0
							END + CASE
								WHEN Day17 IN (3, 12, 14) THEN 1
								ELSE 0
							END + CASE
								WHEN Day18 IN (3, 12, 14) THEN 1
								ELSE 0
							END + CASE
								WHEN Day19 IN (3, 12, 14) THEN 1
								ELSE 0
							END + CASE
								WHEN Day20 IN (3, 12, 14) THEN 1
								ELSE 0
							END + CASE
								WHEN Day21 IN (3, 12, 14) THEN 1
								ELSE 0
							END + CASE
								WHEN Day22 IN (3, 12, 14) THEN 1
								ELSE 0
							END + CASE
								WHEN Day23 IN (3, 12, 14) THEN 1
								ELSE 0
							END + CASE
								WHEN Day24 IN (3, 12, 14) THEN 1
								ELSE 0
							END + CASE
								WHEN Day25 IN (3, 12, 14) THEN 1
								ELSE 0
							END + CASE
								WHEN Day26 IN (3, 12, 14) THEN 1
								ELSE 0
							END + CASE
								WHEN Day27 IN (3, 12, 14) THEN 1
								ELSE 0
							END + CASE
								WHEN Day28 IN (3, 12, 14) THEN 1
								ELSE 0
							END + CASE
								WHEN Day29 IN (3, 12, 14) THEN 1
								ELSE 0
							END + CASE
								WHEN Day30 IN (3, 12, 14) THEN 1
								ELSE 0
							END + CASE
								WHEN Day31 IN (3, 12, 14) THEN 1
								ELSE 0
							END
						) AS 'Night'
					FROM
						plans
					GROUP BY
						NurseID
				) c
				JOIN (
					SELECT
						nurseid As NurseID,
						SUM(
							CASE
								WHEN Day1 = 4 THEN 1
								ELSE 0
							END + CASE
								WHEN Day2 = 4 THEN 1
								ELSE 0
							END + CASE
								WHEN Day3 = 4 THEN 1
								ELSE 0
							END + CASE
								WHEN Day4 = 4 THEN 1
								ELSE 0
							END + CASE
								WHEN Day5 = 4 THEN 1
								ELSE 0
							END + CASE
								WHEN Day6 = 4 THEN 1
								ELSE 0
							END + CASE
								WHEN Day7 = 4 THEN 1
								ELSE 0
							END + CASE
								WHEN Day8 = 4 THEN 1
								ELSE 0
							END + CASE
								WHEN Day9 = 4 THEN 1
								ELSE 0
							END + CASE
								WHEN Day10 = 4 THEN 1
								ELSE 0
							END + CASE
								WHEN Day11 = 4 THEN 1
								ELSE 0
							END + CASE
								WHEN Day12 = 4 THEN 1
								ELSE 0
							END + CASE
								WHEN Day13 = 4 THEN 1
								ELSE 0
							END + CASE
								WHEN Day14 = 4 THEN 1
								ELSE 0
							END + CASE
								WHEN Day15 = 4 THEN 1
								ELSE 0
							END + CASE
								WHEN Day16 = 4 THEN 1
								ELSE 0
							END + CASE
								WHEN Day17 = 4 THEN 1
								ELSE 0
							END + CASE
								WHEN Day18 = 4 THEN 1
								ELSE 0
							END + CASE
								WHEN Day19 = 4 THEN 1
								ELSE 0
							END + CASE
								WHEN Day20 = 4 THEN 1
								ELSE 0
							END + CASE
								WHEN Day21 = 4 THEN 1
								ELSE 0
							END + CASE
								WHEN Day22 = 4 THEN 1
								ELSE 0
							END + CASE
								WHEN Day23 = 4 THEN 1
								ELSE 0
							END + CASE
								WHEN Day24 = 4 THEN 1
								ELSE 0
							END + CASE
								WHEN Day25 = 4 THEN 1
								ELSE 0
							END + CASE
								WHEN Day26 = 4 THEN 1
								ELSE 0
							END + CASE
								WHEN Day27 = 4 THEN 1
								ELSE 0
							END + CASE
								WHEN Day28 = 4 THEN 1
								ELSE 0
							END + CASE
								WHEN Day29 = 4 THEN 1
								ELSE 0
							END + CASE
								WHEN Day30 = 4 THEN 1
								ELSE 0
							END + CASE
								WHEN Day31 = 4 THEN 1
								ELSE 0
							END
						) AS 'Dayoff'
					FROM
						plans
					GROUP BY
						NurseID
				) d
				JOIN (
					SELECT
						nurseid As NurseID,
						SUM(
							CASE
								WHEN Day1 = 8 THEN 1
								ELSE 0
							END + CASE
								WHEN Day2 = 8 THEN 1
								ELSE 0
							END + CASE
								WHEN Day3 = 8 THEN 1
								ELSE 0
							END + CASE
								WHEN Day4 = 8 THEN 1
								ELSE 0
							END + CASE
								WHEN Day5 = 8 THEN 1
								ELSE 0
							END + CASE
								WHEN Day6 = 8 THEN 1
								ELSE 0
							END + CASE
								WHEN Day7 = 8 THEN 1
								ELSE 0
							END + CASE
								WHEN Day8 = 8 THEN 1
								ELSE 0
							END + CASE
								WHEN Day9 = 8 THEN 1
								ELSE 0
							END + CASE
								WHEN Day10 = 8 THEN 1
								ELSE 0
							END + CASE
								WHEN Day11 = 8 THEN 1
								ELSE 0
							END + CASE
								WHEN Day12 = 8 THEN 1
								ELSE 0
							END + CASE
								WHEN Day13 = 8 THEN 1
								ELSE 0
							END + CASE
								WHEN Day14 = 8 THEN 1
								ELSE 0
							END + CASE
								WHEN Day15 = 8 THEN 1
								ELSE 0
							END + CASE
								WHEN Day16 = 8 THEN 1
								ELSE 0
							END + CASE
								WHEN Day17 = 8 THEN 1
								ELSE 0
							END + CASE
								WHEN Day18 = 8 THEN 1
								ELSE 0
							END + CASE
								WHEN Day19 = 8 THEN 1
								ELSE 0
							END + CASE
								WHEN Day20 = 8 THEN 1
								ELSE 0
							END + CASE
								WHEN Day21 = 8 THEN 1
								ELSE 0
							END + CASE
								WHEN Day22 = 8 THEN 1
								ELSE 0
							END + CASE
								WHEN Day23 = 8 THEN 1
								ELSE 0
							END + CASE
								WHEN Day24 = 8 THEN 1
								ELSE 0
							END + CASE
								WHEN Day25 = 8 THEN 1
								ELSE 0
							END + CASE
								WHEN Day26 = 8 THEN 1
								ELSE 0
							END + CASE
								WHEN Day27 = 8 THEN 1
								ELSE 0
							END + CASE
								WHEN Day28 = 8 THEN 1
								ELSE 0
							END + CASE
								WHEN Day29 = 8 THEN 1
								ELSE 0
							END + CASE
								WHEN Day30 = 8 THEN 1
								ELSE 0
							END + CASE
								WHEN Day31 = 8 THEN 1
								ELSE 0
							END
						) AS 'Oncall'
					FROM
						plans
					GROUP BY
						NurseID
				) e
				JOIN (
					SELECT
						nurseid As NurseID,
						SUM(
							CASE
								WHEN Day1 IN (10, 11, 12, 13, 14, 15, 16, 17, 18, 20, 21, 22) THEN 1
								ELSE 0
							END + CASE
								WHEN Day2 IN (10, 11, 12, 13, 14, 15, 16, 17, 18, 20, 21, 22) THEN 1
								ELSE 0
							END + CASE
								WHEN Day3 IN (10, 11, 12, 13, 14, 15, 16, 17, 18, 20, 21, 22) THEN 1
								ELSE 0
							END + CASE
								WHEN Day4 IN (10, 11, 12, 13, 14, 15, 16, 17, 18, 20, 21, 22) THEN 1
								ELSE 0
							END + CASE
								WHEN Day5 IN (10, 11, 12, 13, 14, 15, 16, 17, 18, 20, 21, 22) THEN 1
								ELSE 0
							END + CASE
								WHEN Day6 IN (10, 11, 12, 13, 14, 15, 16, 17, 18, 20, 21, 22) THEN 1
								ELSE 0
							END + CASE
								WHEN Day7 IN (10, 11, 12, 13, 14, 15, 16, 17, 18, 20, 21, 22) THEN 1
								ELSE 0
							END + CASE
								WHEN Day8 IN (10, 11, 12, 13, 14, 15, 16, 17, 18, 20, 21, 22) THEN 1
								ELSE 0
							END + CASE
								WHEN Day9 IN (10, 11, 12, 13, 14, 15, 16, 17, 18, 20, 21, 22) THEN 1
								ELSE 0
							END + CASE
								WHEN Day10 IN (10, 11, 12, 13, 14, 15, 16, 17, 18, 20, 21, 22) THEN 1
								ELSE 0
							END + CASE
								WHEN Day11 IN (10, 11, 12, 13, 14, 15, 16, 17, 18, 20, 21, 22) THEN 1
								ELSE 0
							END + CASE
								WHEN Day12 IN (10, 11, 12, 13, 14, 15, 16, 17, 18, 20, 21, 22) THEN 1
								ELSE 0
							END + CASE
								WHEN Day13 IN (10, 11, 12, 13, 14, 15, 16, 17, 18, 20, 21, 22) THEN 1
								ELSE 0
							END + CASE
								WHEN Day14 IN (10, 11, 12, 13, 14, 15, 16, 17, 18, 20, 21, 22) THEN 1
								ELSE 0
							END + CASE
								WHEN Day15 IN (10, 11, 12, 13, 14, 15, 16, 17, 18, 20, 21, 22) THEN 1
								ELSE 0
							END + CASE
								WHEN Day16 IN (10, 11, 12, 13, 14, 15, 16, 17, 18, 20, 21, 22) THEN 1
								ELSE 0
							END + CASE
								WHEN Day17 IN (10, 11, 12, 13, 14, 15, 16, 17, 18, 20, 21, 22) THEN 1
								ELSE 0
							END + CASE
								WHEN Day18 IN (10, 11, 12, 13, 14, 15, 16, 17, 18, 20, 21, 22) THEN 1
								ELSE 0
							END + CASE
								WHEN Day19 IN (10, 11, 12, 13, 14, 15, 16, 17, 18, 20, 21, 22) THEN 1
								ELSE 0
							END + CASE
								WHEN Day20 IN (10, 11, 12, 13, 14, 15, 16, 17, 18, 20, 21, 22) THEN 1
								ELSE 0
							END + CASE
								WHEN Day21 IN (10, 11, 12, 13, 14, 15, 16, 17, 18, 20, 21, 22) THEN 1
								ELSE 0
							END + CASE
								WHEN Day22 IN (10, 11, 12, 13, 14, 15, 16, 17, 18, 20, 21, 22) THEN 1
								ELSE 0
							END + CASE
								WHEN Day23 IN (10, 11, 12, 13, 14, 15, 16, 17, 18, 20, 21, 22) THEN 1
								ELSE 0
							END + CASE
								WHEN Day24 IN (10, 11, 12, 13, 14, 15, 16, 17, 18, 20, 21, 22) THEN 1
								ELSE 0
							END + CASE
								WHEN Day25 IN (10, 11, 12, 13, 14, 15, 16, 17, 18, 20, 21, 22) THEN 1
								ELSE 0
							END + CASE
								WHEN Day26 IN (10, 11, 12, 13, 14, 15, 16, 17, 18, 20, 21, 22) THEN 1
								ELSE 0
							END + CASE
								WHEN Day27 IN (10, 11, 12, 13, 14, 15, 16, 17, 18, 20, 21, 22) THEN 1
								ELSE 0
							END + CASE
								WHEN Day28 IN (10, 11, 12, 13, 14, 15, 16, 17, 18, 20, 21, 22) THEN 1
								ELSE 0
							END + CASE
								WHEN Day29 IN (10, 11, 12, 13, 14, 15, 16, 17, 18, 20, 21, 22) THEN 1
								ELSE 0
							END + CASE
								WHEN Day30 IN (10, 11, 12, 13, 14, 15, 16, 17, 18, 20, 21, 22) THEN 1
								ELSE 0
							END + CASE
								WHEN Day31 IN (10, 11, 12, 13, 14, 15, 16, 17, 18, 20, 21, 22) THEN 1
								ELSE 0
							END
						) AS 'Ot'
					FROM
						plans
					GROUP BY
						NurseID
				) f ON a.NurseID = b.NurseID
				AND b.NurseID = c.NurseID
				AND c.NurseID = d.NurseID
				AND d.NurseID = e.NurseID
				AND e.NurseID = f.NurseID
		) calinfo ON Nursesinfo.Nurse_ID = calinfo.NurseID


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
