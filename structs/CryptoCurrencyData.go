package structs


type CryptoCurrencyData struct {
  Name string
  Symbol string
  Rank int
  Market_cap float64
  Volume_24h float64
  Img *string
  RsiData []Rsi
}

type Rsi struct {
  Rsi float64
  Time_period string
}
