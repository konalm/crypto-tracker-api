package abstractRatesByTimePeriod

import (
  // "fmt"
  "crypto-tracker-api/structs"
)


/**
 *
 */
func OneHourPeriods(rates []structs.CryptoRate) []structs.CryptoRate {
  // fmt.Println("one hour periods")

  var ratesInOneHourPeriod []structs.CryptoRate

  for _, rate := range rates {
    if rate.Min == 0 {
      ratesInOneHourPeriod = append(ratesInOneHourPeriod, rate)
    }
  }

  return Latest15Rates(ratesInOneHourPeriod)
}
