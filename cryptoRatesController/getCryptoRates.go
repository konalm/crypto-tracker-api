package cryptoRatesController

import (
  "sort"
  _ "github.com/go-sql-driver/mysql"
  "net/http"
  "encoding/json"
  "stelita-api/rankedCryptoCurrency"
  "stelita-api/db"
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
  Img *string
  Rates []CryptoRate
}


/**
 * slice rates from crypto currencies that are greater that 15 periods ago
 */
func GetCryptoCurrencies(w http.ResponseWriter, r * http.Request) {
  type CryptoCurrency struct {
    Currency string
  }

  db := db.Conn()
  defer db.Close()

  query := `SELECT DISTINCT currency FROM crypto_rates`

  rows, err := db.Query(query)
  if err != nil {
    panic(err.Error())
  }
  defer rows.Close()

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

  /* for testing */
  // rankedCryptoCurrencySymbols = rankedCryptoCurrencySymbols[0:1]

  /* open database connection */
  db := db.Conn()
  defer db.Close()

  query :=
    `SELECT ranked_cryptos.name, ranked_cryptos.symbol, ranked_cryptos.rank,
      ranked_cryptos.market_cap, ranked_cryptos.volume_24h,
      logos.img AS img,
      crypto_rates.date, crypto_rates.closing_price, crypto_rates.min
    FROM ranked_crypto_currencies ranked_cryptos
    LEFT JOIN crypto_currency_logos logos
      ON logos.currency = ranked_cryptos.name
    LEFT JOIN crypto_rates
      ON crypto_rates.currency = ranked_cryptos.symbol`

  queryValues := []interface{}{}

  for index, symbol := range rankedCryptoCurrencySymbols {
    if index == 0 {
      query += " WHERE"
    } else {
      query += " OR"
    }

    query += " crypto_rates.currency = ?"
    queryValues = append(queryValues, symbol)
  }


  query +=
    ` AND crypto_rates.date > (NOW() - INTERVAL 16 DAY)
    ORDER BY crypto_rates.date DESC`

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
      &cryptoCurrency.Img,
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
