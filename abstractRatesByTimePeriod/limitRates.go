package abstractRatesByTimePeriod

import (
  "stelita-api/structs"
)

/**
 *
 */
func LimitRates (rates []structs.CryptoRate, limit int)  []structs.CryptoRate {
  ratesLength := len(rates)

  if ratesLength < limit {
    return rates
  }

  return rates[ratesLength - limit : ratesLength]
}
