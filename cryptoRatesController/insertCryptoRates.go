package cryptoRatesController

import (
	"crypto-tracker-api/structs"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	"time"
)


/**
 *
 */
func InsertCryptoRates(rates []structs.BitcoinRate) {
  /* open database connection */
  db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/stelita_dev")
  if err != nil {
    panic(err.Error())
  }

  query := `INSERT INTO crypto_rates (currency, date, closing_price, min)`
  var queryValues = []interface{}{}

  timeOfCronJob := time.Now()
  minOfCronJob := timeOfCronJob.Minute()

  for i, rate := range rates {
    query += buildRateModalQuery(rate, i)
    formattedDate := formatDateForMysql(rate.Time)

		if rate.Asset_id_quote == "USD" { rate.Asset_id_quote = "BTC" }

    queryValues = append(queryValues, rate.Asset_id_quote, formattedDate, rate.Rate, minOfCronJob)
  }

  query = removeLastComma(query)
  stmt, _ := db.Prepare(query)

  _, err = stmt.Exec(queryValues...)
  if err != nil {
    panic(err.Error())
  }

  stmt.Close()
  db.Close()
}


/**
 * build query to insert single rate modal
 */
func buildRateModalQuery(rate structs.BitcoinRate, i int) string {
  var query string

  if i == 0 {
    query += " VALUES"
  }

  query += " (?,?,?,?),"

  return query
}


/**
 *
 */
func formatDateForMysql(date string) time.Time {
  var replacer = strings.NewReplacer("T", " ", "Z", "")
  preparedDate := replacer.Replace(date)

  formattedDate, err := time.Parse("2006-01-02 15:04:05", preparedDate)
  if err != nil {
    fmt.Println("Error: could not format date")
  }

  return formattedDate
}


/**
 * Remove the last comma from the query string
 */
func removeLastComma(query string) string  {
  queryx := len(query)

  if queryx > 0 && query[queryx-1] == ',' {
    query = query[:queryx-1]
  }

  return query
}
