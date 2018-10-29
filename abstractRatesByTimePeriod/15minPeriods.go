package abstractRatesByTimePeriod

import (
  "stelita-api/structs"
)


/**
 *
 */
func FifteenMinPeriods(rates []structs.CryptoRate) []structs.CryptoRate {
  var ratesIn15MinPeriod []structs.CryptoRate

  for _, rate := range rates {
    if rate.Min == 0 || rate.Min == 15 || rate.Min == 30 || rate.Min == 45 {
      ratesIn15MinPeriod = append(ratesIn15MinPeriod, rate)
    }
  }

  return ratesIn15MinPeriod
}
