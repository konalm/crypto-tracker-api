package bitcoinRates

import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "net/http"
  "encoding/json"
  "sort"
)


/**
 *
 */
func GetBitcoinRates(w http.ResponseWriter, r * http.Request) {
  type BitcoinRate struct {
    Date string
    Closing_price float64
  }

  /* open database connection */
  db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/stelita_dev")
  if err != nil {
    panic(err.Error())
  }

  /* get by 15min periods */
  query :=
    `SELECT date, closing_price
    FROM bitcoin_rates
    WHERE min = 0
      OR min = 15
      OR min = 30
      OR min = 45
    ORDER BY date DESC
    LIMIT 15
  `

  rows, err := db.Query(query)
  if err != nil {
    panic(err.Error())
  }

  var bitcoinRates []BitcoinRate

  for rows.Next() {
    var bitcoinRate BitcoinRate

    err := rows.Scan(&bitcoinRate.Date, &bitcoinRate.Closing_price)
    if err != nil {
      panic(err.Error())
    }

    bitcoinRates = append(bitcoinRates, bitcoinRate)
  }

  /* sort by lowest dates */
  sort.SliceStable(bitcoinRates, func(i, j int) bool {
    return bitcoinRates[i].Date < bitcoinRates[j].Date
  })

  json.NewEncoder(w).Encode(bitcoinRates)
}


// func main() {
//   router := mux.NewRouter()
//   router.HandleFunc("/rates", getLatestBitcoinRates).Methods("GET")
//
//   testpackage.PrintTest()
//
//   log.Fatal(http.ListenAndServe(":8484", router))
// }
