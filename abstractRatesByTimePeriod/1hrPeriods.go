package abstractRatesByTimePeriod

import (
  "stelita-api/structs"
)


/**
 *
 */
func OneHourPeriods(rates []structs.CryptoRate) []structs.CryptoRate {
  var ratesInOneHourPeriod []structs.CryptoRate

  for _, rate := range rates {
    if rate.Min == 0 {
      ratesInOneHourPeriod = append(ratesInOneHourPeriod, rate)
    }
  }

  return ratesInOneHourPeriod
}
