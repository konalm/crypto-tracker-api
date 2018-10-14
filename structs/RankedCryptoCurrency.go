package structs


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
  Volumne_24h float64
  Market_cap float64
  Percent_change_1h float64
  Percent_change_24h float64
  Percent_change_7d float64
}
