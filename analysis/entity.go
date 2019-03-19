package analysis


type CryptoCurrencyAnalysis struct {
  CryptoCurrency string
  CryptoCurrencyLogoImgPath string
  TotalGainPercent float64
  TotalLossPercent float64
  AveragePercent float64
  InProgressCount int
  AnalysisItems []Analysis
}


type Analysis struct {
  TimeInterval string
  Smoothing string
  StartPrice float64
  StartRsi float64
  StartDate string
  EndPrice float64
  EndRsi float64
  EndDate string
  GainPercent float64
  LossPercent float64
  DurationHours int
  Complete bool
  DataReport []AnalysisReportDataPoint
}


type AnalysisReportDataPoint struct {
  Rsi float64
  ClosingPrice float64
  Date string
}
