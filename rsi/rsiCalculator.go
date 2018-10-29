package rsi

import (
  "math"
  "stelita-api/structs"
  "stelita-api/abstractRatesByTimePeriod"
)


var DummyCryptoRates = []structs.CryptoRate {
  structs.CryptoRate {
    Date: "01/01/2018",
    ClosingPrice: 44.34,
  },
  structs.CryptoRate {
    Date: "02/01/2018",
    ClosingPrice: 44.09,
  },
  structs.CryptoRate {
    Date: "03/01/2018",
    ClosingPrice: 44.15,
  },
  structs.CryptoRate {
    Date: "04/01/2018",
    ClosingPrice: 43.61,
  },
  structs.CryptoRate {
    Date: "05/01/2018",
    ClosingPrice: 44.33,
  },
  structs.CryptoRate {
    Date: "06/01/2018",
    ClosingPrice: 44.83,
  },
  structs.CryptoRate {
    Date: "07/01/2018",
    ClosingPrice: 45.10,
  },
  structs.CryptoRate {
    Date: "08/01/2018",
    ClosingPrice: 45.42,
  },
  structs.CryptoRate {
    Date: "09/01/2018",
    ClosingPrice: 45.84,
  },
  structs.CryptoRate {
    Date: "10/01/2018",
    ClosingPrice: 46.08,
  },
  structs.CryptoRate {
    Date: "11/01/2018",
    ClosingPrice: 45.89,
  },
  structs.CryptoRate {
    Date: "12/01/2018",
    ClosingPrice: 46.03,
  },
  structs.CryptoRate {
    Date: "13/01/2018",
    ClosingPrice: 45.61,
  },
  structs.CryptoRate {
    Date: "14/01/2018",
    ClosingPrice: 46.28,
  },
  structs.CryptoRate {
    Date: "15/01/2018",
    ClosingPrice: 46.28,
  },
}


/**
 *
 */
func CalculateRsi(_cryptoRates []structs.CryptoRate) float64 {
  if len(_cryptoRates) < 15 {
    return 0.00
  }

  var cryptoRates = abstractRatesByTimePeriod.LimitRates(_cryptoRates, 15)

  avgGain := calcAverageGain(cryptoRates)
  avgLoss := calcAverageLoss(cryptoRates)
  rsf := calcRsf(avgGain, avgLoss)
  rsi := calcRsi(rsf)

  if math.IsNaN(rsi) {
    return 0.00
  }

  return rsi
}


/**
 *
 */
func previousRates(cryptoRates []structs.CryptoRate, index int) []structs.CryptoRate {
  return cryptoRates[0:index]
}


/**
 *
 */
func calcClosingPriceChange(closingPrice float64, previousClosingPrice float64) float64 {
  return closingPrice - previousClosingPrice
}


/**
 * calculate the average gain of all crypto rates
 */
func calcAverageGain(rates []structs.CryptoRate) float64 {
  totalGain := 0.00

  for index, rate := range rates {
    if index == 0 { continue }

    closingPriceChange :=
      calcClosingPriceChange(rate.ClosingPrice, rates[index - 1].ClosingPrice)

    if closingPriceChange > 0 {
      totalGain += closingPriceChange
    }
  }

  return totalGain / float64(len(rates))
}


/**
 * calculate the average loss of all crypto rates
 */
func calcAverageLoss(rates []structs.CryptoRate) float64 {
  totalLoss := 0.00

  for i, rate := range rates {
    if i == 0 { continue }

    closingPriceChange :=
      calcClosingPriceChange(rate.ClosingPrice, rates[i - 1].ClosingPrice)

    if closingPriceChange < 0 {
      totalLoss += closingPriceChange
    }
  }

  return (totalLoss * -1) / float64(len(rates))
}


/**
 * calculate relative strength factor
 */
func calcRsf(avgGain float64, avgLoss float64) float64 {
  return avgGain / avgLoss
}


/**
 * calculate relative strength index
 */
func calcRsi(rsf float64) float64 {
  return 100 - (100 / (1 + rsf))
}
