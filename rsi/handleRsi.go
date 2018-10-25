package rsi

import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "fmt"
  "encoding/json"
  "crypto-tracker-api/cryptoRatesController"
  "crypto-tracker-api/abstractRatesByTimePeriod"
  "crypto-tracker-api/rankedCryptoCurrency"
  "crypto-tracker-api/structs"
  // "reflect"
)

type TrendStat struct {
  Time_period string
  Rsi float64
  RateChange float64
  MovingAverages structs.MovingAverage
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

  fifteenMin10Ma := CalculateMovingAverage(ratesIn15MinPeriod, 10)
  fifteenMin25Ma := CalculateMovingAverage(ratesIn15MinPeriod, 25)
  fifteenMin50Ma := CalculateMovingAverage(ratesIn15MinPeriod, 50)
  fifteenMin100Ma := CalculateMovingAverage(ratesIn15MinPeriod, 100)


  /* 1 hr period */
  ratesIn1HrPeriod := abstractRatesByTimePeriod.OneHourPeriods(rates)
  oneHrRsi := CalculateRsi(ratesIn1HrPeriod)
  oneHrRateChange := CalculateRateChange(ratesIn1HrPeriod)

  oneHr10Ma := CalculateMovingAverage(ratesIn1HrPeriod, 10)
  oneHr25Ma := CalculateMovingAverage(ratesIn1HrPeriod, 25)
  oneHr50Ma := CalculateMovingAverage(ratesIn1HrPeriod, 50)
  oneHr100Ma := CalculateMovingAverage(ratesIn1HrPeriod, 100)


  /* 3hr period */
  ratesIn3HrPeriod := abstractRatesByTimePeriod.ThreeHourPeriods(rates)
  threeHrRsi := CalculateRsi(ratesIn3HrPeriod)
  threeHrRateChange := CalculateRateChange(ratesIn3HrPeriod)

  threeHr10Ma := CalculateMovingAverage(ratesIn3HrPeriod, 10)
  threeHr25Ma := CalculateMovingAverage(ratesIn3HrPeriod, 25)
  threeHr50Ma := CalculateMovingAverage(ratesIn3HrPeriod, 50)
  threeHr100Ma := CalculateMovingAverage(ratesIn3HrPeriod, 100)


  /* 24hr period */
  ratesIn24HrPeriod := abstractRatesByTimePeriod.TwentyFourPeriods(rates)
  oneDayRsi := CalculateRsi(ratesIn24HrPeriod)
  oneDayRateChange := CalculateRateChange(ratesIn24HrPeriod)

  oneDay10Ma := CalculateMovingAverage(ratesIn24HrPeriod, 10)
  oneDay25Ma := CalculateMovingAverage(ratesIn24HrPeriod, 25)
  oneDay50Ma := CalculateMovingAverage(ratesIn24HrPeriod, 50)
  oneDay100Ma := CalculateMovingAverage(ratesIn24HrPeriod, 100)


  var trendStats = []TrendStat {
    TrendStat {
      Time_period: "15min",
      Rsi: fifteenMinRsi,
      RateChange: fifteenMinRateChange,
      MovingAverages: structs.MovingAverage {
        LengthOf10: fifteenMin10Ma,
        LengthOf25: fifteenMin25Ma,
        LengthOf50: fifteenMin50Ma,
        LengthOf100: fifteenMin100Ma,
      },
    },
    TrendStat {
      Time_period: "1hr",
      Rsi: oneHrRsi,
      RateChange: oneHrRateChange,
      MovingAverages: structs.MovingAverage {
        LengthOf10: oneHr10Ma,
        LengthOf25: oneHr25Ma,
        LengthOf50: oneHr50Ma,
        LengthOf100: oneHr100Ma,
      },
    },
    TrendStat {
      Time_period: "3hr",
      Rsi: threeHrRsi,
      RateChange: threeHrRateChange,
      MovingAverages: structs.MovingAverage {
        LengthOf10: threeHr10Ma,
        LengthOf25: threeHr25Ma,
        LengthOf50: threeHr50Ma,
        LengthOf100: threeHr100Ma,
      },
    },
    TrendStat {
      Time_period: "24hr",
      Rsi: oneDayRsi,
      RateChange: oneDayRateChange,
      MovingAverages: structs.MovingAverage {
        LengthOf10: oneDay10Ma,
        LengthOf25: oneDay25Ma,
        LengthOf50: oneDay50Ma,
        LengthOf100: oneDay100Ma,
      },
    },
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

  stmt.Close()
  db.Close()
}
