package rankedCryptoCurrency


type RankedCryptoCurrency struct {
  Id int
  Name string
  Symbol string
  Website_slug string
  Rank int
  Circulating_supply float64
  Total_supply float64
  Max_supply float64
  Quotes map[string]Quote
}

type Quote struct {
  Price float64
  Volume_24h float64
  Market_cap float64
  Percent_change_1h float64
  Percent_change_24h float64
  Percent_change_7d float64
}

type CryptoCurrencyData struct {
  Id int
  Name string
  Symbol string
  Rank int
  Market_cap float64
  Volume_24h float64
  Img *string
  SellIndicator bool
  BuyIndicator bool
  InWallet bool
  TrendStats []TrendStat
}

type TrendStat struct {
  Time_period string
  RsiStats RsiStats
  RateChange float64
  MovingAverages MovingAverage
}

type MovingAverage struct {
  LengthOf10 float64
  LengthOf25 float64
  LengthOf50 float64
  LengthOf100 float64
}


type RsiStats struct {
  Rsi float64
  Smoothing50 float64
  Smoothing100 float64
  Smoothing250 float64
}
