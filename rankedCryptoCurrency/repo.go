package rankedCryptoCurrency

import (
  "database/sql"
  "fmt"
  _ "github.com/go-sql-driver/mysql"
  "crypto-tracker-api/structs"
)


/**
 *
 */
func InsertRankedCryptoCurrencies(cryptoCurrencies []structs.RankedCryptoCurrency) {
  fmt.Println("insert ranked cryptos")
  
  query := `INSERT INTO ranked_crypto_currencies ()`

  /* open database connection */
  db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/stelita_dev")
  if err != nil {
    panic(err.Error())
  }

  for _, crypto := range cryptoCurrencies {
    fmt.Println("in crypto loop")
    fmt.Println(crypto)
  }
}
