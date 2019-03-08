package rsi

type TrendStat struct {
  Time_period string
  Rsi float64
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
