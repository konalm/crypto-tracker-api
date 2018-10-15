package rankedCryptoCurrency

import (
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
