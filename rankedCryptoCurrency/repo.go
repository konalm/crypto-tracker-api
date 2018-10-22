package rankedCryptoCurrency

import (
  "net/http"
  "encoding/json"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "crypto-tracker-api/structs"
  "crypto-tracker-api/utils"
)

/**
 *
 */
func InsertRankedCryptoCurrencies(
  cryptoCurrencies map[string] structs.RankedCryptoCurrency,
) {
  /* open database connection */
  db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/stelita_dev")
  if err != nil {
    panic(err.Error())
  }

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
  stmt, _ := db.Prepare(query)

  _, err = stmt.Exec(queryValues...)
  if err != nil {
    panic("ERROR executing query" + err.Error())
  }

  defer stmt.Close()
  defer db.Close()
}


/**
 *
 */
func DestroyCurrentRankedCryptoCurrencies() {
  /* open database connection */
  db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/stelita_dev")
  if err != nil {
    panic(err.Error())
  }

  destroy, err := db.Query("DELETE FROM ranked_crypto_currencies")
  if err != nil {
    panic("ERROR destroying current ranked crpyot currencies")
  }

  defer destroy.Close()
}

/**
 *
 */
func GetSymbols() []string {
  /* open database connection */
  db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/stelita_dev")
  if err != nil {
    panic(err.Error())
  }

  rows, err := db.Query("SELECT symbol FROM ranked_crypto_currencies")
  if err != nil {
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
  /* open database connection */
  db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/stelita_dev")
  if err != nil {
    panic(err.Error())
  }

  query :=
    `SELECT crypto.name, crypto.symbol, crypto.rank, crypto.market_cap,
      crypto.volume_24h, crypto.rsi,
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
    var rsi []byte

    err := rows.Scan(
      &crypto.Name, &crypto.Symbol, &crypto.Rank, &crypto.Market_cap,
      &crypto.Volume_24h, &rsi, &crypto.Img,
    )
    if err != nil {
      panic(err.Error())
    }

    if err := json.Unmarshal(rsi, &crypto.RsiData); err != nil {
      panic(err.Error())
    }

    cryptoCurrencyData = append(cryptoCurrencyData, crypto)
  }

  json.NewEncoder(w).Encode(cryptoCurrencyData)
}
