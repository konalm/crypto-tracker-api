package cryptoRatesController

import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "net/http"
  "encoding/json"
  "fmt"
)


/**
 *
 */
func GetCryptoCurrencies(w http.ResponseWriter, r * http.Request) {
  type CryptoCurrency struct {
    Currency string
  }

  /* open database connection */
  db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/stelita_dev")
  if err != nil {
    panic(err.Error())
  }

  query := `SELECT DISTINCT currency FROM crypto_rates`

  rows, err := db.Query(query)
  if err != nil {
    panic(err.Error())
  }

  var cryptoCurrencies []CryptoCurrency

  for rows.Next() {
    var cryptoCurrency CryptoCurrency

    err := rows.Scan(&cryptoCurrency.Currency)
    if err != nil {
      panic(err.Error())
    }

    cryptoCurrencies = append(cryptoCurrencies, cryptoCurrency)
  }

  json.NewEncoder(w).Encode(cryptoCurrencies)
}


/**
 *
 */
func GetCryptoCurrencyRates(w http.ResponseWriter, r *http.Request) {
  fmt.Println("get crypto currency rates")

  type CryptoCurrencyRate struct {
    Currency string
    Date string
    Closing_price float64
    Min int
  }

  /* open database connection */
  db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/stelita_dev")
  if err != nil {
    panic(err.Error())
  }

  query := `SELECT currency, date, closing_price, min FROM crypto_rates LIMIT 40000`

  rows, err := db.Query(query)
  if err != nil {
    panic(err.Error())
  }

  var cryptoCurrencyRates = make(map[string][]CryptoCurrencyRate)

  for rows.Next() {
    var cryptoCurrencyRate CryptoCurrencyRate

    err := rows.Scan(
      &cryptoCurrencyRate.Currency,
      &cryptoCurrencyRate.Date,
      &cryptoCurrencyRate.Closing_price,
      &cryptoCurrencyRate.Min,
    )
    if err != nil {
      panic(err.Error())
    }

    cryptoCurrencyRates[cryptoCurrencyRate.Currency] =
      append(cryptoCurrencyRates[cryptoCurrencyRate.Currency], cryptoCurrencyRate)
  }

  json.NewEncoder(w).Encode(cryptoCurrencyRates)
}
