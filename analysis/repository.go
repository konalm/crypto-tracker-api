package analysis

import (
  "fmt"
  "encoding/json"
  "stelita-api/db"
)

/**
 *
 */
func GetAnalysis() []CryptoCurrencyAnalysis {
  fmt.Println("REPO >> Get Analysis")

  conn := db.Conn()
  defer conn.Close()

  query :=
    `SELECT crypto_currency,
      logo.img logo_img_path,
      IF (SUM(gain_percent) IS NULL, 0.00, SUM(gain_percent)) total_gain_percent,
      IF (SUM(loss_percent) IS NULL, 0.00, SUM(loss_percent)) total_loss_percent,
      SUM(CASE WHEN complete = 0 THEN 1 ELSE 1 END) is_complete_count
    FROM analysis

    INNER JOIN ranked_crypto_currencies ranked_crypto
      ON ranked_crypto.symbol = analysis.crypto_currency

    INNER JOIN crypto_currency_logos logo
      ON logo.currency = ranked_crypto.name

    GROUP BY crypto_currency, logo_img_path`

  rows, err := conn.Query(query)
  if err != nil {
    panic("error executing select stmt for analysis")
  }

  var analysis []CryptoCurrencyAnalysis

  for rows.Next() {
    var cryptoCurrencyAnalysis CryptoCurrencyAnalysis

    err := rows.Scan(
      &cryptoCurrencyAnalysis.CryptoCurrency,
      &cryptoCurrencyAnalysis.CryptoCurrencyLogoImgPath,
      &cryptoCurrencyAnalysis.TotalGainPercent,
      &cryptoCurrencyAnalysis.TotalLossPercent,
      &cryptoCurrencyAnalysis.InProgressCount,
    )
    if err != nil {
      panic(err.Error())
    }

    cryptoCurrencyAnalysis.AveragePercent =
      cryptoCurrencyAnalysis.TotalGainPercent -
      cryptoCurrencyAnalysis.TotalLossPercent

    analysis = append(analysis, cryptoCurrencyAnalysis)
  }

  return analysis
}

/**
 *
 */
func GetCryptoCurrencyAnalysis(cryptoCurrency string) CryptoCurrencyAnalysis {
  conn := db.Conn()
  defer conn.Close()

  query :=
    `SELECT
      IF(ranked_crypto.name IS NULL, "", ranked_crypto.name) name,
      IF(logo.img IS NULL, "", logo.img) logo_img_path,
      a.start_date,
      a.start_value,
      IF(a.end_value IS NULL, 0.00, a.end_value) end_value,
      IF(a.gain_percent IS NULL, 0.00, a.gain_percent) gain_percent,
      IF(a.loss_percent IS NULL, 0.00, a.loss_percent) loss_percent
    FROM analysis a

    LEFT JOIN ranked_crypto_currencies ranked_crypto
      ON ranked_crypto.symbol = a.crypto_currency

    LEFT JOIN crypto_currency_logos logo
      ON logo.currency = ranked_crypto.name

    WHERE a.crypto_currency = ?`

  rows, err := conn.Query(query, cryptoCurrency)
  if err != nil {
    panic("Error executing select query for crypto currency analysis")
  }

  var cryptoCurrencyAnalysis CryptoCurrencyAnalysis

  for rows.Next() {
    var analysis Analysis

    err := rows.Scan(
      &cryptoCurrencyAnalysis.CryptoCurrency,
      &cryptoCurrencyAnalysis.CryptoCurrencyLogoImgPath,
      &analysis.StartDate,
      &analysis.StartPrice,
      &analysis.EndPrice,
      &analysis.GainPercent,
      &analysis.LossPercent,
    )
    if err != nil { panic(err.Error()) }

    cryptoCurrencyAnalysis.TotalGainPercent += analysis.GainPercent
    cryptoCurrencyAnalysis.TotalLossPercent += analysis.LossPercent
    cryptoCurrencyAnalysis.AnalysisItems =
      append(cryptoCurrencyAnalysis.AnalysisItems, analysis)
  }

  cryptoCurrencyAnalysis.AveragePercent =
    cryptoCurrencyAnalysis.TotalGainPercent - cryptoCurrencyAnalysis.TotalLossPercent

  return cryptoCurrencyAnalysis
}


/**
 *
 */
func GetAnalysisItem(id string) Analysis {
  conn := db.Conn()
  defer conn.Close()

  query :=
    `SELECT time_interval,
      smoothing,
      start_value,
      rsi,
      start_date,
      IF (end_value IS NULL, "", end_value) end_value,
      IF (end_date IS NULL, "", end_date) end_date,
      IF (gain_percent IS NULL, 0.00, gain_percent) gain_percent,
      IF (loss_Percent IS NULL, 0.00, loss_percent) loss_percent,
      IF (duration_hours IS NULL, 0, duration_hours) duration_hours,
      complete,
      reports_json

    FROM analysis a
    WHERE id = ?`

  stmt := conn.QueryRow(query, id)

  var analysis Analysis
  var dataReportJson string

  err := stmt.Scan(
           &analysis.TimeInterval,
           &analysis.Smoothing,
           &analysis.StartPrice,
           &analysis.StartRsi,
           &analysis.StartDate,
           &analysis.EndPrice,
           &analysis.EndDate,
           &analysis.GainPercent,
           &analysis.LossPercent,
           &analysis.DurationHours,
           &analysis.Complete,
           &dataReportJson,
         )
  if err != nil { panic(err.Error()) }

  if dataReportJson != "" {
    err := json.Unmarshal([]byte(dataReportJson), &analysis.DataReport)

    if err != nil { panic(err.Error()) }
  }


  return analysis
}
