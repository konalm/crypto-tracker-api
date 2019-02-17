package fetchCryptoRates

import (
	"stelita-api/structs"
	. "encoding/json"
	. "fmt"
	. "io/ioutil"
	"net/http"
	"stelita-api/fetchCryptoRates/fetchrate"
)

func FetchBitcoinRate() structs.BitcoinRate {
  return fetchrate(
		"GET",
		"https://rest.coinapi.io/v1/exchangerate/BTC/USD",
	)
}
