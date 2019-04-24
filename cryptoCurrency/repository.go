package cryptoCurrency

import (
  "database/sql"
  "stelita-api/db"
)

/**
 *
 */
func GetCryptoCurrencyData(symbol string) CryptoCurrencyItemData {
  conn := db.Conn()
  defer conn.Close()

  query :=  `SELECT name, symbol, rank, market_cap, volume_24h,
    logo.img
    FROM ranked_crypto_currencies
    LEFT JOIN crypto_currency_logos logo
      on logo.currency = name
    WHERE symbol = ?`

  stmt := conn.QueryRow(query, symbol)

  var cryptoCurrencyData CryptoCurrencyItemData
  err := stmt.Scan(
    &cryptoCurrencyData.Name,
    &cryptoCurrencyData.Symbol,
    &cryptoCurrencyData.Rank,
    &cryptoCurrencyData.MarketCap,
    &cryptoCurrencyData.Volume24h,
    &cryptoCurrencyData.Img,
  )

  if err != nil {
    if err == sql.ErrNoRows {
      return CryptoCurrencyItemData{
        Name: "",
        Symbol: "",
        Rank: 0,
        MarketCap: 0,
        Volume24h: 0,
      }
    }

    panic(err.Error())
  }

  return cryptoCurrencyData
}
