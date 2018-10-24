package abstractRatesByTimePeriod

import (
  // "fmt"
  // "reflect"
  "crypto-tracker-api/structs"
  "time"
)


/**
 *
 */
func ThreeHourPeriods(rates []structs.CryptoRate) []structs.CryptoRate {
  // fmt.Println("three hour periods")

  var ratesIn3HrPeriod []structs.CryptoRate

  for _, rate := range rates {
    t, err := time.Parse("2006-01-02 15:04:05", rate.Date)
    if err != nil {
      panic(err.Error())
    }

    hour := t.Hour()

    if (hour == 0 ||
        hour == 3 ||
        hour == 6 ||
        hour == 9 ||
        hour == 12 ||
        hour == 15 ||
        hour == 18 ||
        hour == 21) && rate.Min == 0 {
      ratesIn3HrPeriod = append(ratesIn3HrPeriod, rate)
    }
  }

  return ratesIn3HrPeriod
}
