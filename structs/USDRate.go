package structs

type USDRate struct {
  Name string
  Symbol string
  Slug string
  Circulating_supply float64
  Total_supply float64
  Max_supply float64
  Date_added string
  Quotes []USDQuote
}


type USDQuote struct {
  Price float64
  Volumne_24h float64
  Percent_change_1h float64
  Percent_change_24h float64
  Percent_change_7d float64
  Market_cap float64
  Last_updated float64
}
