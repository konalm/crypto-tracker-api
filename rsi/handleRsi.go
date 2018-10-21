package rsi

import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "fmt"
  "crypto-tracker-api/cryptoRatesController"
  "crypto-tracker-api/abstractRatesByTimePeriod"
  "crypto-tracker-api/rankedCryptoCurrency"
  "encoding/json"
)

type Rsi struct {
  Rsi float64
  Time_period string
}


func HandleRsi() {
  fmt.Println("handle RSI")

  cryptoCurrencies := rankedCryptoCurrency.GetSymbols()

  for _, cryptoCurrency := range cryptoCurrencies {
    fmt.Println("call handle rsi for >>> " + cryptoCurrency)

    go handleCryptoRsi(cryptoCurrency)
  }
}


/**
 *
 */
func handleCryptoRsi(cryptoCurrency string) {
  fmt.Println("handle crypto rsi >>> " + cryptoCurrency)
  rates := cryptoRatesController.GetCryptoCurrencyRatesForRsi(cryptoCurrency)

  ratesIn15MinPeriod := abstractRatesByTimePeriod.FifteenMinPeriods(rates)
  fifteenMinRsi := CalculateRsi(ratesIn15MinPeriod)

  ratesIn1HrPeriod := abstractRatesByTimePeriod.OneHourPeriods(rates)
  oneHrRsi := CalculateRsi(ratesIn1HrPeriod)

  ratesIn3HrPeriod := abstractRatesByTimePeriod.ThreeHourPeriods(rates)
  threeHrRsi := CalculateRsi(ratesIn3HrPeriod)

  ratesIn24HrPeriod := abstractRatesByTimePeriod.TwentyFourPeriods(rates)
  oneDayRsi := CalculateRsi(ratesIn24HrPeriod)

  var rsiData = []Rsi {
    Rsi {Time_period: "15min", Rsi: fifteenMinRsi,},
    Rsi {Time_period: "1hr", Rsi: oneHrRsi,},
    Rsi {Time_period: "3hr", Rsi: threeHrRsi,},
    Rsi {Time_period: "24hr", Rsi: oneDayRsi,},
  }

  rsiJson, err := json.Marshal(rsiData)
  if err != nil {
    panic(err.Error())
  }

  updateCryptoRsi(cryptoCurrency, rsiJson)
}


/**
 *
 */
func updateCryptoRsi(cryptoCurrency string, rsiJson []byte) {
  fmt.Println("update crypo rsi")

  if !rsiJson {
    return
  }

  /* open database connection */
  db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/stelita_dev")
  if err != nil {
    panic(err.Error())
  }

  query := "UPDATE ranked_crypto_currencies SET rsi = ? WHERE symbol = ?"

  stmt, err := db.Prepare(query)
  if err != nil {
    panic(err.Error())
  }

  _, err = stmt.Exec(rsiJson, cryptoCurrency)
  if err != nil {
    panic(err.Error())
  }

  defer stmt.Close()
  defer db.Close()
}
