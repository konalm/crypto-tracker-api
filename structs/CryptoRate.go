package structs


type CryptoRate struct {
  Date string
  ClosingPrice float64
  ClosingPriceChange float64
  AverageGain float64
  AverageLoss float64
  Min int
  RelativeStrengthFactor float64
  RelativeStrengthIndex float64
}
