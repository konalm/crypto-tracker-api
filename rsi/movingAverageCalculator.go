package rsi

import (
  "fmt"
  "crypto-tracker-api/structs"
  "crypto-tracker-api/abstractRatesByTimePeriod"
)


/**
 *
 */
func CalculateMovingAverage(_cryptoRates []structs.CryptoRate, length int) float64 {
  fmt.Println("calculate moving average")

  if len(_cryptoRates) < length {
    return 0.00
  }

  var cryptoRates = abstractRatesByTimePeriod.LimitRates(_cryptoRates, length)
  rateTotal := 0.00

  for _, cryptoRate := range cryptoRates {
    fmt.Println("calc moving average -> in loop")

    rateTotal += cryptoRate.ClosingPrice
  }

  return rateTotal / float64(length)
}
