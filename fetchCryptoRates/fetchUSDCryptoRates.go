package fetchCryptoRates

import (
  "net/http"
  "encoding/json"
  "fmt"
  "io/ioutil"
  "stelita-api/structs"
)


type CoinMarketGapApiResponse struct {
  Data []structs.USDRate
}


/**
 *
 */
func FetchUSDCryptoRates() []structs.USDRate {
  client := &http.Client{}
  requestUrl :=  "https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest"
  request, err := http.NewRequest("GET", requestUrl, nil)
  request.Header.Set("X-CMC_PRO_API_KEY", `a65f7aff-3f3b-4471-8d62-63fd0b7729a6`)

  response, err := client.Do(request)
  if err != nil {
    fmt.Println(err.Error())
  }

  if response.StatusCode != 200 {
    panic("ERROR fetching usd crypto rates")
  }

  body, err := ioutil.ReadAll(response.Body)
  if err != nil {
    fmt.Println("ERROR reading response body")
  }

  jsonBody := string(body)
  var coinMarketGapApiResponse CoinMarketGapApiResponse

  json.Unmarshal([]byte(jsonBody), &coinMarketGapApiResponse)

  return coinMarketGapApiResponse.Data
}
