package structs


type CryptoCurrencyData struct {
  Name string
  Symbol string
  Rank int
  Market_cap float64
  Volume_24h float64
  Img *string
  TrendStats []TrendStat
}

type TrendStat struct {
  Time_period string
  Rsi float64
  RateChange float64
  MovingAverages MovingAverage
}

type MovingAverage struct {
  LengthOf10 float64
  LengthOf25 float64
  LengthOf50 float64
  LengthOf100 float64
}
