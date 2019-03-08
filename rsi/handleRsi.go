package rsi

import (
  "fmt"
  "encoding/json"
  "stelita-api/cryptoRatesController"
  "stelita-api/abstractRatesByTimePeriod"
  "stelita-api/rankedCryptoCurrency"
  "stelita-api/db"
  "stelita-api/reports"
  "stelita-api/errorReporter"
)


/**
 *
 */
func HandleRsi() {
  cryptoCurrencies := rankedCryptoCurrency.GetSymbols()
  maxRoutines := 10
  handleCryptoTrendChannel := make(chan struct{}, maxRoutines)

  for _, cryptoCurrency := range cryptoCurrencies {
    handleCryptoTrendChannel <- struct{}{} // block if limit reached
    go func() {
      handleCryptoTrendStats(cryptoCurrency)
      <-handleCryptoTrendChannel
    }()
  }
}


/**
 *
 */
func handleCryptoTrendStats(cryptoCurrency string) {
  fmt.Println("handle crypto trend stats")
  fmt.Println(cryptoCurrency)

  rates := cryptoRatesController.GetCryptoCurrencyRatesForRsi(cryptoCurrency)

  /* 15 min period */
  ratesIn15MinPeriod := abstractRatesByTimePeriod.FifteenMinPeriods(rates)
  fifteenMinRsi := CalculateRsi(ratesIn15MinPeriod, 0)
  fifteenMinRsiSmoothing50 := CalculateRsi(ratesIn15MinPeriod, 50)
  fifteenMinRsiSmoothing100 := CalculateRsi(ratesIn15MinPeriod, 100)
  fifteenMinRsiSmoothing250 := CalculateRsi(ratesIn15MinPeriod, 250)

  fifteenMinRateChange := CalculateRateChange(ratesIn15MinPeriod)

  fifteenMin10Ma := CalculateMovingAverage(ratesIn15MinPeriod, 10)
  fifteenMin25Ma := CalculateMovingAverage(ratesIn15MinPeriod, 25)
  fifteenMin50Ma := CalculateMovingAverage(ratesIn15MinPeriod, 50)
  fifteenMin100Ma := CalculateMovingAverage(ratesIn15MinPeriod, 100)

  /* 1 hr period */
  ratesIn1HrPeriod := abstractRatesByTimePeriod.OneHourPeriods(rates)
  oneHrRsi := CalculateRsi(ratesIn1HrPeriod, 0)
  oneHrRsiSmoothing50 := CalculateRsi(ratesIn1HrPeriod, 50)
  oneHrRsiSmoothing100 := CalculateRsi(ratesIn1HrPeriod, 100)
  oneHrRsiSmoothing250 := CalculateRsi(ratesIn1HrPeriod, 250)

  oneHrRateChange := CalculateRateChange(ratesIn1HrPeriod)

  oneHr10Ma := CalculateMovingAverage(ratesIn1HrPeriod, 10)
  oneHr25Ma := CalculateMovingAverage(ratesIn1HrPeriod, 25)
  oneHr50Ma := CalculateMovingAverage(ratesIn1HrPeriod, 50)
  oneHr100Ma := CalculateMovingAverage(ratesIn1HrPeriod, 100)

  /* 3hr period */
  ratesIn3HrPeriod := abstractRatesByTimePeriod.ThreeHourPeriods(rates)
  threeHrRsi := CalculateRsi(ratesIn3HrPeriod, 0)
  threeHrRsiSmoothing50 := CalculateRsi(ratesIn3HrPeriod, 50)
  threeHrRsiSmoothing100 := CalculateRsi(ratesIn3HrPeriod, 100)
  threeHrRsiSmoothing250 := CalculateRsi(ratesIn3HrPeriod, 250)

  threeHrRateChange := CalculateRateChange(ratesIn3HrPeriod)

  threeHr10Ma := CalculateMovingAverage(ratesIn3HrPeriod, 10)
  threeHr25Ma := CalculateMovingAverage(ratesIn3HrPeriod, 25)
  threeHr50Ma := CalculateMovingAverage(ratesIn3HrPeriod, 50)
  threeHr100Ma := CalculateMovingAverage(ratesIn3HrPeriod, 100)

  /* 24hr period */
  ratesIn24HrPeriod := abstractRatesByTimePeriod.TwentyFourPeriods(rates)
  oneDayRsi := CalculateRsi(ratesIn24HrPeriod, 0)
  oneDayRsiSmoothing50 := CalculateRsi(ratesIn24HrPeriod, 50)
  oneDayRsiSmoothing100 := CalculateRsi(ratesIn24HrPeriod, 100)
  oneDayRsiSmoothing250 := CalculateRsi(ratesIn24HrPeriod, 250)

  oneDayRateChange := CalculateRateChange(ratesIn24HrPeriod)

  oneDay10Ma := CalculateMovingAverage(ratesIn24HrPeriod, 10)
  oneDay25Ma := CalculateMovingAverage(ratesIn24HrPeriod, 25)
  oneDay50Ma := CalculateMovingAverage(ratesIn24HrPeriod, 50)
  oneDay100Ma := CalculateMovingAverage(ratesIn24HrPeriod, 100)


  var trendStats = []TrendStat {
    TrendStat {
      Time_period: "15min",
      RsiStats: RsiStats {
        Rsi: fifteenMinRsi,
        Smoothing50: fifteenMinRsiSmoothing50,
        Smoothing100: fifteenMinRsiSmoothing100,
        Smoothing250: fifteenMinRsiSmoothing250,
      },
      RateChange: fifteenMinRateChange,
      MovingAverages: MovingAverage {
        LengthOf10: fifteenMin10Ma,
        LengthOf25: fifteenMin25Ma,
        LengthOf50: fifteenMin50Ma,
        LengthOf100: fifteenMin100Ma,
      },
    },
    TrendStat {
      Time_period: "1hr",
      RsiStats: RsiStats {
        Rsi: oneHrRsi,
        Smoothing50: oneHrRsiSmoothing50,
        Smoothing100: oneHrRsiSmoothing100,
        Smoothing250: oneHrRsiSmoothing250,
      },
      RateChange: oneHrRateChange,
      MovingAverages: MovingAverage {
        LengthOf10: oneHr10Ma,
        LengthOf25: oneHr25Ma,
        LengthOf50: oneHr50Ma,
        LengthOf100: oneHr100Ma,
      },
    },
    TrendStat {
      Time_period: "3hr",
      RsiStats: RsiStats {
        Rsi: threeHrRsi,
        Smoothing50: threeHrRsiSmoothing50,
        Smoothing100: threeHrRsiSmoothing100,
        Smoothing250: threeHrRsiSmoothing250,
      },
      RateChange: threeHrRateChange,
      MovingAverages: MovingAverage {
        LengthOf10: threeHr10Ma,
        LengthOf25: threeHr25Ma,
        LengthOf50: threeHr50Ma,
        LengthOf100: threeHr100Ma,
      },
    },
    TrendStat {
      Time_period: "24hr",
      RsiStats: RsiStats {
        Rsi: oneDayRsi,
        Smoothing50: oneDayRsiSmoothing50,
        Smoothing100: oneDayRsiSmoothing100,
        Smoothing250: oneDayRsiSmoothing250,
      },
      RateChange: oneDayRateChange,
      MovingAverages: MovingAverage {
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
  fmt.Println("update crypto trend stats !!")

  trendStatsString := string(trendStatsJson)
  successUpdate := true

  dbConn := db.Conn()
  defer dbConn.Close()

  query := "UPDATE ranked_crypto_currencies SET trend_statistics = ? WHERE symbol = ?"

  stmt, err := dbConn.Prepare(query)
  if err != nil {
    successUpdate = false
  }
  defer stmt.Close()

  _, err = stmt.Exec(trendStatsString, cryptoCurrency)
  if err != nil {
    errorReporter.ReportError("Updating crypto trand statistic")
    successUpdate = false
  }

  fmt.Println("update crypto trend stat >>>> " + cryptoCurrency)

  processListQuery := "SHOW PROCESSLIST";
  stmtProcessList, err := dbConn.Query(processListQuery)
  if err != nil {
    fmt.Println("stmt process list")
  }
  defer stmtProcessList.Close()

  i := 0
  for stmtProcessList.Next() {
    i ++
  }

  reports.InsertUpdateCryptoTrendStatReport(cryptoCurrency, successUpdate, i)
}
