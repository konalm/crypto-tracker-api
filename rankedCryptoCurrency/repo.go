package rankedCryptoCurrency

import (
  "net/http"
  "encoding/json"
  _ "github.com/go-sql-driver/mysql"
  "stelita-api/structs"
  "stelita-api/utils"
  "stelita-api/db"
  "stelita-api/errorReporter"
)

/**
 *
 */
func InsertRankedCryptoCurrencies(
  cryptoCurrencies map[string] structs.RankedCryptoCurrency,
) {
  dbConn := db.Conn()
  defer dbConn.Close()

  query :=
    `INSERT INTO ranked_crypto_currencies
    (name, symbol, rank, market_cap, volume_24h)`
  queryValues := []interface{}{}

  count := 0
  for _, crypto := range cryptoCurrencies {
    if (count == 0) { query += " VALUES" }

    query += " (?,?,?,?,?),"
    quote := crypto.Quotes["USD"]

    queryValues = append(
      queryValues,
      crypto.Name, crypto.Symbol, crypto.Rank, quote.Market_cap, quote.Volume_24h,
    )

    count++;
  }

  query = utils.RemoveLastComma(query)
  stmt, _ := dbConn.Prepare(query)

  _, err := stmt.Exec(queryValues...)
  if err != nil {
    errorReporter.ReportError("Inserting ranked crypto currencies")
  }

  defer stmt.Close()
}


/**
 *
 */
func DestroyCurrentRankedCryptoCurrencies() {
  db := db.Conn()
  defer db.Close()

  destroy, err := db.Query("DELETE FROM ranked_crypto_currencies")
  if err != nil {
    panic("ERROR destroying current ranked crypto currencies")
  }

  defer destroy.Close()
}

/**
 *
 */
func GetSymbols() []string {
  dbConn := db.Conn()
  defer dbConn.Close()

  rows, err := dbConn.Query("SELECT symbol FROM ranked_crypto_currencies")
  if err != nil {
    panic(err.Error())
    panic("ERROR getting symbols from ranked crypto currencies")
  }

  var symbols []string

  for rows.Next() {
    var symbol string

    err := rows.Scan(&symbol)
    if err != nil {
      panic(err.Error())
    }

    symbols = append(symbols, symbol)
  }

  return symbols
}


/**
 *
 */
func GetCryptoCurrencyData(w http.ResponseWriter, r *http.Request) {
  db := db.Conn()
  defer db.Close()

  query :=
    `SELECT crypto.id, crypto.name, crypto.symbol, crypto.rank, crypto.market_cap,
      crypto.volume_24h, crypto.trend_statistics,
      logo.img
    FROM ranked_crypto_currencies crypto
    LEFT JOIN crypto_currency_logos logo
      ON logo.currency = crypto.name
    ORDER BY rank`

  var cryptoCurrencyData []structs.CryptoCurrencyData

  rows, err := db.Query(query)
  if err != nil {
    panic(err.Error())
  }

  for rows.Next() {
    var crypto structs.CryptoCurrencyData
    var trendStats []byte

    err := rows.Scan(
      &crypto.Id, &crypto.Name, &crypto.Symbol, &crypto.Rank, &crypto.Market_cap,
      &crypto.Volume_24h, &trendStats, &crypto.Img,
    )
    if err != nil {
      panic(err.Error())
    }

    if string(trendStats) != "" {
      if err := json.Unmarshal(trendStats, &crypto.TrendStats); err != nil {
        panic(err.Error())
      }
    }

    cryptoCurrencyData = append(cryptoCurrencyData, crypto)
  }

  json.NewEncoder(w).Encode(cryptoCurrencyData)
}
