package cryptoRatesController

import (
  "fmt"
  "strings"
  _ "github.com/go-sql-driver/mysql"
  "database/sql"
  "time"
  "rest/structs"
)


type X struct {
  currency string
  date time.Time
  closingPrice float64
  min int
}


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
  var queryValues []X

  for _, rate := range rates {
    var replacer = strings.NewReplacer("T", " ", "Z", "")
    preparedDate := replacer.Replace(rate.Time)
    formattedDate, err := time.Parse("2006-01-02 15:04:05", preparedDate)
    if err != nil {
      fmt.Println(err.Error())
    }

    timeOfCronJob := time.Now()
    minOfCronJob := timeOfCronJob.Minute()

    var queryV X
    queryV.currency = rate.Asset_id_quote
    queryV.date = formattedDate
    queryV.closingPrice = rate.Rate
    queryV.min = minOfCronJob

    query += " VALUES(?,?,?,?),"
    queryValues = append(queryValues, queryV)
  }

  queryx := len(query)
  if queryx > 0 && query[queryx-1] == ',' {
    query = query[:queryx-1]
  }

  fmt.Println("Query >>>>")
  fmt.Println(query)

  stmt, err := db.Prepare(query)
  if err != nil {
    panic(err.Error())
  }

  _, err = stmt.Exec(queryValues)
  if err != nil {
    panic(err.Error())
  }

  defer stmt.Close()
  defer db.Close()
}
