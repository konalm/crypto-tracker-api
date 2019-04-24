package cryptoCurrency

import (
  "stelita-api/structs"
)

type CryptoCurrencyItemData struct {
  Name string
  Symbol string
  Rank int
  MarketCap float64
  Volume24h float64
  ClosingPrice float64
  RSI float64
  Img string
  Rates []structs.CryptoRate
}
