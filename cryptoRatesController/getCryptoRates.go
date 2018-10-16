package cryptoRatesController

import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "net/http"
  "encoding/json"
  "crypto-tracker-api/rankedCryptoCurrency"
  // "fmt"
  "sort"
)


type CryptoRate struct {
  Currency string
  Date string
  Closing_price float64
  Min int
}

type CryptoCurrency struct {
  Name string
  Symbol string
  Rank int
  Market_cap float64
  Volume_24h float64
  Rates []CryptoRate
}


/**
 * slice rates from crypto currencies that are greater that 15 periods ago
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
  rankedCryptoCurrencySymbols := rankedCryptoCurrency.GetSymbols()

  /* open database connection */
  db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/stelita_dev")
  if err != nil {
    panic(err.Error())
  }

  query :=
    `SELECT ranked_cryptos.name, ranked_cryptos.symbol, ranked_cryptos.rank,
      ranked_cryptos.market_cap, ranked_cryptos.volume_24h,
      crypto_rates.date, crypto_rates.closing_price, crypto_rates.min
    FROM ranked_crypto_currencies ranked_cryptos
    LEFT JOIN crypto_rates
      ON crypto_rates.currency = ranked_cryptos.symbol`

  queryValues := []interface{}{}

  for index, symbol := range rankedCryptoCurrencySymbols {
    if index == 0 {
      query += " WHERE"
    } else {
      query += " OR"
    }

    query += " currency = ?"
    queryValues = append(queryValues, symbol)
  }

  query +=
    ` ORDER BY date DESC
    LIMIT 1500`

  rows, err := db.Query(query, queryValues...)
  if err != nil {
    panic(err.Error())
  }

  var cryptoCurrencies = make(map[string]CryptoCurrency)

  for rows.Next() {
    var cryptoCurrency CryptoCurrency
    var cryptoRate CryptoRate

    err := rows.Scan(
      &cryptoCurrency.Name,
      &cryptoCurrency.Symbol,
      &cryptoCurrency.Rank,
      &cryptoCurrency.Market_cap,
      &cryptoCurrency.Volume_24h,
      &cryptoRate.Date,
      &cryptoRate.Closing_price,
      &cryptoRate.Min,
    )
    if err != nil {
      panic(err.Error())
    }

    if _, ok := cryptoCurrencies[cryptoCurrency.Name]; ok == false {
      cryptoCurrencies[cryptoCurrency.Name] = cryptoCurrency
    }

    /* cannot assign to struct field in map workaround */
    var x = cryptoCurrencies[cryptoCurrency.Name]
    x.Rates = append(x.Rates, cryptoRate)
    cryptoCurrencies[cryptoCurrency.Name] = x
  }

  cryptoCurrencies = limitCryptoCurrencyRates(cryptoCurrencies)

  json.NewEncoder(w).Encode(cryptoCurrencies)
}


/**
 *
 */
func limitCryptoCurrencyRates(cryptoCurrencies map[string]CryptoCurrency) map[string]CryptoCurrency {
  for currency, cryptoCurrency := range cryptoCurrencies {
    sort.Slice(cryptoCurrency.Rates, func(x, y int) bool {
      return cryptoCurrency.Rates[x].Date < cryptoCurrency.Rates[y].Date
    })

    rateCount := len(cryptoCurrency.Rates)

    if rateCount > 15 {
      sliceAmount := rateCount - 15

      cryptoCurrency.Rates = cryptoCurrency.Rates[sliceAmount:rateCount]
      cryptoCurrencies[currency] = cryptoCurrency
    }
  }

  return cryptoCurrencies
}
