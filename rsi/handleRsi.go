package rsi

import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "fmt"
  "crypto-tracker-api/cryptoRatesController"
  "crypto-tracker-api/abstractRatesByTimePeriod"
  "crypto-tracker-api/rankedCryptoCurrency"
  "encoding/json"
  // "reflect"
)

type TrendStat struct {
  Time_period string
  Rsi float64
  RateChange float64
}


func HandleRsi() {
  fmt.Println("handle RSI")

  cryptoCurrencies := rankedCryptoCurrency.GetSymbols()

  // cryptoCurrencies = cryptoCurrencies[0:3]

  for _, cryptoCurrency := range cryptoCurrencies {
    fmt.Println("call handle trend stats for >>> " + cryptoCurrency)

    go handleCryptoTrendStats(cryptoCurrency)
  }
}


/**
 *
 */
func handleCryptoTrendStats(cryptoCurrency string) {
  rates := cryptoRatesController.GetCryptoCurrencyRatesForRsi(cryptoCurrency)

  /* 15 min period */
  ratesIn15MinPeriod := abstractRatesByTimePeriod.FifteenMinPeriods(rates)
  fifteenMinRsi := CalculateRsi(ratesIn15MinPeriod)
  fifteenMinRateChange := CalculateRateChange(ratesIn15MinPeriod)

  /* 1 hr period */
  ratesIn1HrPeriod := abstractRatesByTimePeriod.OneHourPeriods(rates)
  oneHrRsi := CalculateRsi(ratesIn1HrPeriod)
  oneHrRateChange := CalculateRateChange(ratesIn1HrPeriod)

  /* 3hr period */
  ratesIn3HrPeriod := abstractRatesByTimePeriod.ThreeHourPeriods(rates)
  threeHrRsi := CalculateRsi(ratesIn3HrPeriod)
  threeHrRateChange := CalculateRateChange(ratesIn3HrPeriod)

  /* 24hr period */
  ratesIn24HrPeriod := abstractRatesByTimePeriod.TwentyFourPeriods(rates)
  oneDayRsi := CalculateRsi(ratesIn24HrPeriod)
  oneDayRateChange := CalculateRateChange(ratesIn24HrPeriod)

  var trendStats = []TrendStat {
    TrendStat {Time_period: "15min", Rsi: fifteenMinRsi, RateChange: fifteenMinRateChange},
    TrendStat {Time_period: "1hr", Rsi: oneHrRsi, RateChange: oneHrRateChange},
    TrendStat {Time_period: "3hr", Rsi: threeHrRsi, RateChange: threeHrRateChange},
    TrendStat {Time_period: "24hr", Rsi: oneDayRsi, RateChange: oneDayRateChange},
  }

  trendStatsJson, err := json.Marshal(trendStats)
  if err != nil {
    fmt.Println("ERROR JSON MARSHAL !!!!!!")
    fmt.Println(trendStats)
  }

  updateCryptoTrendStats(cryptoCurrency, trendStatsJson)
}


/**
 *
 */
func updateCryptoTrendStats(cryptoCurrency string, trendStatsJson []byte) {
  trendStatsString := string(trendStatsJson)

  /* open database connection */
  db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/stelita_dev")
  if err != nil {
    panic(err.Error())
  }

  query := "UPDATE ranked_crypto_currencies SET trend_statistics = ? WHERE symbol = ?"

  stmt, err := db.Prepare(query)
  if err != nil {
    fmt.Print("PREPARE ERROR !!")
    panic(err.Error())
  }

  _, err = stmt.Exec(trendStatsString, cryptoCurrency)
  if err != nil {
    fmt.Println("EXEC ERROR !!")
    panic(err.Error())
  }

  defer stmt.Close()
  defer db.Close()
}
