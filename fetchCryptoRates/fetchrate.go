package fetchCryptoRates

import (
	"stelita-api/structs"
	. "encoding/json"
	. "fmt"
	. "io/ioutil"
	"net/http"
)

func FetchRate(httpMethod, url, panicMessage) structs.BitcoinRate {
	client := &http.Client{}
	request, err := http.NewRequest(httpMethod, url, nil)
	request.Header.Set("X-CoinApi-Key", `E4C3D4AE-29D8-4A9F-BD36-EB367D836532`)

	resp, err := client.Do(request)
	if err != nil {
		Println("handle http error")
	}

	if resp.StatusCode != 200 {
		panic(panicMessage)
	}

	body, err := ReadAll(resp.Body)
	if err != nil {
		Println("ERROR reading response body")
	}
	jsonBody := string(body)
	var rate structs.BitcoinRate
	json.Unmarshal([]byte(jsonBody), &rate)

	return rate 
}
