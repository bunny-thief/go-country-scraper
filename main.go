package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gocolly/colly"

	_ "github.com/lib/pq"
)

type country struct {
	Country    string
	Capital    string
	Population int
	Area       float32
}

func main() {
	countries := []country{}

	c := colly.NewCollector(
		colly.MaxDepth(1),
	)

	c.OnHTML(".country", func(h *colly.HTMLElement) {
		country := country{}
		country.Country = h.ChildText(".country-name")
		fmt.Println(country.Country)
		country.Capital = h.ChildText(".country-capital")
		fmt.Println(country.Capital)
		population, err := strconv.Atoi(h.ChildText(".country-population"))

		if err != nil {
			panic(err)
		}

		country.Population = population
		fmt.Println(country.Population)

		area, err := strconv.ParseFloat(h.ChildText(".country-area"), 32)

		if err != nil {
			panic(err)
		}

		// Calling ParseFloat with a bitSize of 32 returns a float64
		// that needs to be converted to float32
		country.Area = float32(area)
		fmt.Println(country.Area)

		countries = append(countries, country)
		fmt.Println("")
	})

	c.Visit("https://www.scrapethissite.com/pages/simple/")

	dbName := os.Getenv("COUNTRY_SCRAPER_DBNAME")
	username := os.Getenv("COUNTRY_SCRAPER_USERNAME")
	password := os.Getenv("COUNTRY_SCRAPER_PASSWORD")

	// open db connection
	connStr := fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", username, password, dbName)
	db, err := sql.Open("postgres", connStr)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	createCountriesTable(db)

}

func createCountriesTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS country(
		id SERIAL PRIMARY KEY,
		country_name VARCHAR(50),
		capital VARCHAR(40),
		population int,
		area		DECIMAL(8,2)
	)`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

}
