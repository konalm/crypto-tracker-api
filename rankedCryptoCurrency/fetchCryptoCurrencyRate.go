package rankedCryptoCurrency

import (
  "net/http"
  "fmt"
  "io/ioutil"
  "errors"
  "encoding/json"
  "stelita-api/errorReporter"
)

type CryptoCurrency struct {
  Time string
  Asset_id_base string
  Asset_id_quote string
  Rate float64
}


/**
 *
 */
func FetchCryptoCurrencyExchangeRate(withdrawCurrency string, depositCurrency string) (float64, error) {
  client := &http.Client{}
  url := "https://rest.coinapi.io/v1/exchangerate/" + withdrawCurrency + "/" + depositCurrency

  request, err := http.NewRequest("GET", url, nil)
  request.Header.Set("X-CoinApi-Key", `0E8FD9DE-805B-4DE8-A8B7-2458CA1B7D83`)

  resp, err := client.Do(request)
  if err != nil {
    panic(err.Error())
  }

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    panic("Error reading response body")
  }

  if resp.StatusCode != 200 {
    errorReporter.ReportError("Fetching crypto currency rate from coinapi")
    return -1.00, errors.New("Error Fetching crypto currency rate")
  }

  fmt.Println("fetch crypto currency exchange rate >> response status")
  fmt.Println(resp.StatusCode)

  jsonBody := string(body)
  var apiResponse CryptoCurrency
  json.Unmarshal([]byte(jsonBody), &apiResponse)

  return apiResponse.Rate, nil
}
