package abstractRatesByTimePeriod

import (
  // "fmt"
  "time"
  "crypto-tracker-api/structs"
)


/**
 *
 */
func TwentyFourPeriods(rates []structs.CryptoRate) []structs.CryptoRate {
  // fmt.Println("24 hour periods")

  var ratesIn24HrPeriod []structs.CryptoRate

  for _, rate := range rates {
    t, err := time.Parse("2006-01-02 15:04:05", rate.Date)
    if err != nil {
      panic(err.Error())
    }

    hour := t.Hour()

    if hour == 0 && rate.Min == 0 {
      ratesIn24HrPeriod = append(ratesIn24HrPeriod, rate)
    }
  }

  return ratesIn24HrPeriod
}
