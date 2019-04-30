package cryptoRates

import (
  "fmt"
  "stelita-api/db"
)

/**
 *
 */
func GetCryptoRates(symbol string, amount int) []CryptoRate {
   conn := db.Conn()
   defer conn.Close()

   query := `SELECT date, closing_price
    FROM crypto_rates
    WHERE currency = ?
    ORDER BY date DESC
    LIMIT ?`


  rows, err := conn.Query(query, symbol, amount)
  if err != nil {
    fmt.Println(err.Error())
    panic("Error executing query to  get crypto rates")
  }

  var cryptoRates []CryptoRate

  for rows.Next() {
    var cryptoRate CryptoRate

    err := rows.Scan(
      &cryptoRate.Date,
      &cryptoRate.ClosingPrice,
    )
    if err != nil { panic(err.Error())  }

    cryptoRates = append(cryptoRates, cryptoRate)
  }

  return cryptoRates
}
