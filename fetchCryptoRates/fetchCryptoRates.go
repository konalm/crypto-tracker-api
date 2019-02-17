package fetchCryptoRates

import (
  "net/http"
  "encoding/json"
  "fmt"
  "io/ioutil"
  "stelita-api/structs"
	"stelita-api/fetchCryptoRates/fetchrate"
)


type ApiResponse struct {
  Asset_id_base string
  Rates []structs.BitcoinRate
}

/**
 *
 */
func FetchCryptoRates() []structs.BitcoinRate {
  return fetchrate(
		"GET",
		"https://rest.coinapi.io/v1/exchangerate/BTC",
	)
}
