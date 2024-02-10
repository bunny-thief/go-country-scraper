![](Go-Logo_Blue_scraper.png)

A Go project that scrapes information about countries of the World from https://www.scrapethissite.com/pages/simple/ and inserts the data into a PostgresQL table.

### Sample Country Data:

**Andorra**  
**Capital**: Andorra la Vella  
**Population**: 84000  
**Area (km2)**: 468.0

## Description

Scrapping is done with [Colly](https://go-colly.org/ 'go-colly.org'). Once the information is scrapped it is inserted into a local PostgresQL database. The scraped data was used to create a Spring/Spring Boot REST API project called [countries-api](https://github.com/bunny-thief/countries-api 'countries-api repo').

This project is a reimplementation of a previous Javascript project ([country-scraper](https://github.com/bunny-thief/country-scraper 'country-scraper')) that used Puppeter to scrape date from the same site.

## Getting Started

### Dependencies

You need four things to make sure the script executes correctly:

1. PostegresQL installed and running.
2. PostgresQL username and password.
3. Create a database using the user created in step 2.
4. Three environment variables to initialize the database name, username and password which are required to establish a connection to the database.

The main.go file assumes you have PostgresQL installed and that it its running on the computer on which you execute it.

Once you have PostgresQL up and running, you will need to create a PostgresQL user and password.

You can then create a database with the user created in the previous step. Creating the database with this user will ensure that the script can connect to the database with their credentials from the script.

Finally, in order to connect to the database server, you will need to add three environment variables; one each for the database name, PostgresQL username and password. The Go script will load the values of those environment variables into strings which are then concatenated to form the connection string.

```go
dbName := os.Getenv("COUNTRY_SCRAPER_DBNAME")
username := os.Getenv("COUNTRY_SCRAPER_USERNAME")
password := os.Getenv("COUNTRY_SCRAPER_PASSWORD")
```

Therefore, you need to name the environment variables COUNTRY_SCRAPER_DBNAME, COUNTRY_SCRAPER_USERNAME and COUNTRY_SCRAPER_PASSWORD. Optionally, you can simply edit those three lines and place the pertinent information directly in the script if you don't want to use environment variables.

```go
dbName := "<database name>"
username := "<username>"
password := "<password>"
```

### Installing

Clone the Github repository.

```
git clone git@github.com:bunny-thief/go-country-scraper.git
```

### Executing program

`cd` into the project directory.

```
cd go-country-scraper
```

Execute the **main.go** file.

```
go run main.go
```

## Authors

![](Mastodon_logo.png) [@bunnythief@hachyderm.io](https://hachyderm.io/@bunnythief)

## License

This project is licensed under the MIT License - see the LICENSE.md file for details
