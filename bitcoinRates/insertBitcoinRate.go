package bitcoinRates

import (
  "fmt"
  "strings"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "time"
  "crypto-tracker-api/structs"
)


/**
 *
 */
func InsertBitcoinRate(rate structs.BitcoinRate) {
  fmt.Println("insert bitcoin rate >>")
  fmt.Println(rate)

  var replacer = strings.NewReplacer("T", " ", "Z", "")
  preparedDate := replacer.Replace(rate.Time)

  timeOfCronJob := time.Now()
  minOfCronJob := timeOfCronJob.Minute()

  formattedDate, err := time.Parse("2006-01-02 15:04:05", preparedDate)
  if err != nil {
    fmt.Println(err.Error())
  }

  /* open database connection */
  db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/stelita_dev")
  if err != nil {
    panic(err.Error())
  }

  var query string =
    `INSERT INTO bitcoin_rates (currency, date, closing_price, min)
    VALUES (?,?,?,?)`

  stmt, err := db.Prepare(query)
  if err != nil {
    panic(err.Error())
  }

  _, err = stmt.Exec(rate.Asset_id_quote, formattedDate, rate.Rate, minOfCronJob)
  if err != nil {
    panic(err.Error())
  }

  defer stmt.Close()
  defer db.Close()
}
