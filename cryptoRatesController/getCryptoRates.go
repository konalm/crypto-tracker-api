package cryptoRatesController

import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "net/http"
  "encoding/json"
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
    var cryptoCurrency  CryptoCurrency

    err := rows.Scan(&cryptoCurrency.Currency)
    if err != nil {
      panic(err.Error())
    }

    cryptoCurrencies = append(cryptoCurrencies, cryptoCurrency)
  }

  json.NewEncoder(w).Encode(cryptoCurrencies)
}
