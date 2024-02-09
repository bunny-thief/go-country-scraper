# Go Country Scraper

![](Go-Logo_Blue.png)

A Go 1.21.6 project that scrapes information about countries of the World from https://www.scrapethissite.com/pages/simple/.

## Country Example:

**Andorra**  
**Capital**: Andorra la Vella  
**Population**: 84000  
**Area (km2)**: 468.0

Scrapping is done with [Colly](https://go-colly.org/). Once the information is scrapped it is inserted into a local PostgresQL database. The scraped data was used to create a Spring/Spring Boot REST API project called [countries-api](https://github.com/bunny-thief/countries-api 'countries-api repo').

This project is a reimplementation of a previous Javascript project ([country-scraper](https://github.com/bunny-thief/country-scraper)) that used Puppeter to scrape date from the same site.
