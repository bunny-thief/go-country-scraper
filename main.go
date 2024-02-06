package main

import (
	"database/sql"
	"fmt"
	"log"
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

	connStr := "postgres://country-scraper:123abc@localhost/countries?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

}
