package main

import (
  "fmt"
  "net/http"
  _ "github.com/go-sql-driver/mysql"
  "io/ioutil"
  "encoding/json"
  // "strconv"
  "github.com/robfig/cron"
  "database/sql"
  "time"
  // "reflect"
  "strings"
)

type testStruct struct {
  Test string
}

type User struct {
  UserId int
  Id int
  Title string
  Body string
}

type CryptoRate struct {
  Time string
  Asset_id_quote string
  Rate float64
}

type CryptoRatesApiResponse struct {
  Asset_id_base string
  Rates []CryptoRate
}

type BitcoinRate struct {
  Time string
  Asset_id_base string
  Asset_id_quote string
  Rate float64
}


/**
 *
 */
func getBitcoinRate() BitcoinRate {
  client := &http.Client{}

  request, err := http.NewRequest("GET", "https://rest.coinapi.io/v1/exchangerate/BTC/USD", nil)
  request.Header.Set("X-CoinApi-Key", `E4C3D4AE-29D8-4A9F-BD36-EB367D836532`)

  resp, err := client.Do(request)
  if err != nil {
    fmt.Println("handle http error")
  }

  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    fmt.Println("ERROR reading response body")
  }

  jsonBody := string(body)

  fmt.Println("api reponse >>")
  fmt.Println(jsonBody)

  var bitcoinRate BitcoinRate

  // json.Unmarshal([]byte(jsonBody), &apiResponse)
  json.Unmarshal([]byte(jsonBody), &bitcoinRate)

  return bitcoinRate
}


/**
 *
 */
func insertBitcoinRate(rate BitcoinRate) {
  fmt.Println("insert bitcoin rate")

  var replacer = strings.NewReplacer("T", " ", "Z", "")
  preparedDate := replacer.Replace(rate.Time)

  timeOfCronJob := time.Now()
  minOfCronJob := timeOfCronJob.Minute()

  formattedDate, err := time.Parse("2006-01-02 15:04:05", preparedDate)
  if err != nil {
    fmt.Println("Error: unable to format date")
    fmt.Println(err.Error())
  }

  // /* open database connection */
  db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/stelita_dev")
  if err != nil {
    panic(err.Error())
  }

  var query string =
    `INSERT INTO bitcoin_rates (currency, date, closing_price, min)
    VALUES (?,?,?,?)`

  stmt, err := db.Prepare(query)
  if err != nil {
    panic(err.Error())
  }

  _, err = stmt.Exec(rate.Asset_id_quote, formattedDate, rate.Rate, minOfCronJob)
  if err != nil {
    panic(err.Error())
  }

  defer stmt.Close()
  defer db.Close()
}


/**
 *
 */
func main() {
  c := cron.New()
  c.Start()

  c.AddFunc("0 */5 * * * *", func() {
    fmt.Println("Every 5th min")
    fmt.Println("call coin api every 5 mins")

    bitcoinRate := getBitcoinRate()
    insertBitcoinRate(bitcoinRate)
  })

  err := http.ListenAndServe(":8000", nil)
  if err != nil {
    panic(err)
  }

  fmt.Println("Listening on port 8000")
}
