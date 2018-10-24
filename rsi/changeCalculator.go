package rsi

import (
  // "math"
  // "fmt"
  "crypto-tracker-api/structs"
)

func CalculateRateChange(cryptoRates []structs.CryptoRate) float64 {
  cryptoRatesLength := len(cryptoRates)

  if cryptoRatesLength <= 1 {
    return 0.00
  }

  latestClosingPrice := cryptoRates[cryptoRatesLength - 1].ClosingPrice
  previousClosingPrice := cryptoRates[cryptoRatesLength - 2].ClosingPrice

  diff := previousClosingPrice - latestClosingPrice
  diffPercent := diff / previousClosingPrice * 100;

  return diffPercent
}
