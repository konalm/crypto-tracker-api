package abstractRatesByTimePeriod

import (
  "crypto-tracker-api/structs"
)

/**
 *
 */
func Latest15Rates (rates []structs.CryptoRate)  []structs.CryptoRate {
  ratesLength := len(rates)

  if ratesLength < 15 {
    return rates
  }

  return rates[ratesLength - 15 : ratesLength]
}
