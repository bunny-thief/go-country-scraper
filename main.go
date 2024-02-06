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

type Country struct {
	Country    string
	Capital    string
	Population int
	Area       float32
}

func main() {
	countries := []Country{}

	c := createCollector()

	//
	c.OnHTML(".country", func(h *colly.HTMLElement) {
		country := Country{}
		country.Country = h.ChildText(".country-name")
		country.Capital = h.ChildText(".country-capital")
		population, err := strconv.Atoi(h.ChildText(".country-population"))

		if err != nil {
			panic(err)
		}

		country.Population = population

		area, err := strconv.ParseFloat(h.ChildText(".country-area"), 32)

		if err != nil {
			panic(err)
		}

		// Calling ParseFloat with a bitSize of 32 returns a float64
		// that needs to be converted to float32
		country.Area = float32(area)

		countries = append(countries, country)
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

	dropTable(db)
	createCountriesTable(db)

	// insert countries into db
	for _, country := range countries {
		insertCountry(db, country)
	}

}

func createCollector() *colly.Collector {
	return colly.NewCollector(
		colly.MaxDepth(1),
	)
}

func createCountriesTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS countries(
		id SERIAL PRIMARY KEY,
		country_name VARCHAR(50),
		capital VARCHAR(40),
		population int
	)`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}

}

func insertCountry(db *sql.DB, country Country) int {
	query := `INSERT INTO countries (country_name, capital, population)
	values($1, $2, $3) RETURNING id`

	var pk int

	err := db.QueryRow(query, country.Country, country.Capital, country.Population).Scan(&pk)

	if err != nil {
		log.Fatal(err)
	}

	return pk
}

func dropTable(db *sql.DB) {
	query := `DROP TABLE countries`

	_, err := db.Exec(query)

	if err != nil {
		log.Fatal(err)
	}

}
