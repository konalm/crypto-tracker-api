package rsi

import (
  "fmt"
  "math"
  "stelita-api/structs"
  // "stelita-api/abstractRatesByTimePeriod"
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
func CalculateRsi(cryptoRates []structs.CryptoRate, smoothingCount int) float64 {
  fmt.Println("calculate RSI")

  if len(cryptoRates) < 15 {
    return 0.00
  }

  if smoothingCount > 0 && len(cryptoRates) < smoothingCount {
    return 0.00
  }

  cryptoRates = rateDataPoints(cryptoRates, smoothingCount)

  for index, cryptoRate := range cryptoRates {
    if index > 0 {
      cryptoRates[index].ClosingPriceChange =
        calcClosingPriceChange(
          cryptoRate.ClosingPrice, cryptoRates[index - 1].ClosingPrice,
        )
    } else {
      cryptoRates[index].ClosingPriceChange = 0.00
    }

    averageGain := calcAverageGainV2(cryptoRates, index)
    averageLoss := calcAverageLossV2(cryptoRates, index)
    relativeStrengthFactor := calculateRelativeStengthFactor(averageGain, averageLoss)

    cryptoRates[index].AverageGain = averageGain
    cryptoRates[index].AverageLoss = averageLoss
    cryptoRates[index].RelativeStrengthFactor = relativeStrengthFactor

    if (index >= 14) {
      cryptoRates[index].RelativeStrengthIndex =
        calculateRelativeStengthIndex(relativeStrengthFactor)
    }
  }

  length := len(cryptoRates)
  relativeStrengthIndex := cryptoRates[length - 1].RelativeStrengthIndex

  if math.IsNaN(relativeStrengthIndex) {
    return 0.00
  }

  return relativeStrengthIndex
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
  // periods := abstractRatesBty
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
 *
 */
func calcAverageGainV2(rates []structs.CryptoRate, index int) float64 {
  previousRates := previousRates(rates, index)
  averageGain := 0.00

  /* get average */
  if (index < 15) {
    totalGain := 0.00
    for _, rate := range previousRates {
      if rate.ClosingPriceChange > 0 {
        totalGain += rate.ClosingPriceChange
      }
    }

    averageGain = totalGain / float64(len(rates))
  } else { /* smoothen average */
      previousAverageGain := rates[index - 1].AverageGain
      currentGain := 0.00
      rate := rates[index]

      if rate.ClosingPriceChange > 0 {
        currentGain = rate.ClosingPriceChange
      }

      averageGain = ((previousAverageGain * 13) + currentGain) / 14
  }

  return averageGain
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
 *
 */
func calcAverageLossV2(rates []structs.CryptoRate, index int) float64 {
  previousRates := previousRates(rates, index)
  averageLoss := 0.00

  /* get average */
  if (index < 15) {
    totalLoss := 0.00
    for _, rate := range previousRates {
      if rate.ClosingPriceChange < 0 {
        totalLoss += (rate.ClosingPriceChange * -1)
      }
    }

    averageLoss = totalLoss / float64(len(rates))
  } else { /* smoothen average */
    rate := rates[index]
    previousAverageLoss := rates[index - 1].AverageLoss
    currentLoss := 0.00

    if rate.ClosingPriceChange < 0 {
      currentLoss = (rate.ClosingPriceChange * -1)
    }

    averageLoss = ((previousAverageLoss * 13) + currentLoss) / 14
  }

  return averageLoss
}


/**
 *
 */
func rateDataPoints(cryptoRates []structs.CryptoRate, smoothingCount int) []structs.CryptoRate {
  cryptoRatesLength := len(cryptoRates)

  if smoothingCount == 0 && cryptoRatesLength >= 15 {
    return cryptoRates[cryptoRatesLength - 15: cryptoRatesLength]
  }

  if cryptoRatesLength >= smoothingCount {
    return cryptoRates[cryptoRatesLength - smoothingCount: cryptoRatesLength]
  }

  return cryptoRates
}


/**
 * calculate relative strength factor
 */
func calculateRelativeStengthFactor(avgGain float64, avgLoss float64) float64 {
  return avgGain / avgLoss
}


/**
 * calculate relative strength index
 */
func calculateRelativeStengthIndex(rsf float64) float64 {
  return 100 - (100 / (1 + rsf))
}
