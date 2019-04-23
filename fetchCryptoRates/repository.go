package fetchCryptoRates


type CoinMarketCapApiResponse struct {
  Status CoinMarketCapApiResponseStatus
  Data []CoinMarketCapCryptoCurrency
}

type CoinMarketCapApiResponseStatus struct {
  Timestamp string
  Error_code int
  Elapsed int
  Credit_count int
}

type CoinMarketCapCryptoCurrency struct {
  Id int
  Name string
  Symbol string
  Slug string
  Circulating_supply int
  Total_supply int
  Max_supply int
  Date_added string
  Num_market_pairs int
  Cmc_rank int
  Last_updated string
  Quote struct {
    BTC CoinMarketCapCryptoCurrencyQuote
  }
}

type CoinMarketCapCryptoCurrencyQuote struct {
  Price float64
}
