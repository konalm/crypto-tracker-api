package rankedCryptoCurrency

import (
  "net/http"
  "fmt"
  "io/ioutil"
  "encoding/json"
  "crypto-tracker-api/structs"
)

type CoinMarketGapApiResponse struct {
  Data map[string] structs.RankedCryptoCurrency
}


/**
 *
 */
func FetchRankedCryptoCurrencies() map[string] structs.RankedCryptoCurrency {
  client := &http.Client{}
  url := "https://api.coinmarketcap.com/v2/ticker/?limit=4"
  request, err := http.NewRequest("GET", url, nil)

  resp, err := client.Do(request)
  if err != nil {
    fmt.Println(err.Error())
  }

  if resp.StatusCode != 200 {
    panic("ERROR fetching ranked crypto currencies")
  }

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    panic("Error reading response body")
  }

  jsonBody := string(body)
  var apiResponse CoinMarketGapApiResponse
  json.Unmarshal([]byte(jsonBody), &apiResponse)


  return apiResponse.Data
}
