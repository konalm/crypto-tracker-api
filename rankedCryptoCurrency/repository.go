package rankedCryptoCurrency

import (
  "fmt"
  "encoding/json"
  // "strings"
  _ "github.com/go-sql-driver/mysql"
  "stelita-api/utils"
  "stelita-api/db"
  "stelita-api/errorReporter"
  "stelita-api/walletCurrency"
  "stelita-api/walletState"
)

/**
 *
 */
func InsertRankedCryptoCurrencies(
  cryptoCurrencies map[string] RankedCryptoCurrency,
) {
  dbConn := db.Conn()
  defer dbConn.Close()

  query :=
    `INSERT INTO ranked_crypto_currencies
    (name, symbol, rank, market_cap, volume_24h)`
  queryValues := []interface{}{}

  count := 0
  for _, crypto := range cryptoCurrencies {
    if (count == 0) { query += " VALUES" }

    query += " (?,?,?,?,?),"
    quote := crypto.Quotes["USD"]

    queryValues = append(
      queryValues,
      crypto.Name, crypto.Symbol, crypto.Rank, quote.Market_cap, quote.Volume_24h,
    )

    count++;
  }

  query = utils.RemoveLastComma(query)
  stmt, _ := dbConn.Prepare(query)

  _, err := stmt.Exec(queryValues...)
  if err != nil {
    errorReporter.ReportError("Inserting ranked crypto currencies")
  }

  defer stmt.Close()
}


/**
 *
 */
func DestroyCurrentRankedCryptoCurrencies() {
  db := db.Conn()
  defer db.Close()

  destroy, err := db.Query("DELETE FROM ranked_crypto_currencies")
  if err != nil {
    panic("ERROR destroying current ranked crypto currencies")
  }

  defer destroy.Close()
}

/**
 *
 */
func GetSymbols() []string {
  dbConn := db.Conn()
  defer dbConn.Close()

  rows, err := dbConn.Query("SELECT symbol FROM ranked_crypto_currencies")
  if err != nil {
    panic(err.Error())
    panic("ERROR getting symbols from ranked crypto currencies")
  }

  var symbols []string

  for rows.Next() {
    var symbol string

    err := rows.Scan(&symbol)
    if err != nil {
      panic(err.Error())
    }

    symbols = append(symbols, symbol)
  }

  return symbols
}


/**
 *
 */
func GetCryptoCurrencyData(
  currenciesInWallet []walletCurrency.WalletCurrencyModel,
) []CryptoCurrencyData {
  fmt.Println("get crypto currency data >>")

  db := db.Conn()
  defer db.Close()

  query :=
    `SELECT crypto.id, crypto.name, crypto.symbol, crypto.rank, crypto.market_cap,
      crypto.volume_24h, crypto.trend_statistics,
      logo.img
    FROM ranked_crypto_currencies crypto
    LEFT JOIN crypto_currency_logos logo
      ON logo.currency = crypto.name
    ORDER BY rank`

  var cryptoCurrencyData []CryptoCurrencyData

  rows, err := db.Query(query)
  if err != nil {
    panic(err.Error())
  }

  i := 0

  for rows.Next() {
    i ++

    fmt.Println("i >> ")
    fmt.Println(i)

    var crypto CryptoCurrencyData
    crypto.SellIndicator = false
    crypto.BuyIndicator = false

    var trendStats []byte

    err := rows.Scan(
      &crypto.Id, &crypto.Name, &crypto.Symbol, &crypto.Rank, &crypto.Market_cap,
      &crypto.Volume_24h, &trendStats, &crypto.Img,
    )
    if err != nil {
      panic(err.Error())
    }

    if string(trendStats) != "" {
      if err := json.Unmarshal(trendStats, &crypto.TrendStats); err != nil {
        fmt.Println("Error unmarshalling trend stat")
        fmt.Println(err.Error())
        // panic(err.Error())
      } else {
        fmt.Println("UNMARSHAL SUCCESS")
      }
    } else {
      fmt.Println("TREND STAT IS NULL")
    }

    fmt.Println("trend stats string >>>")
    fmt.Println(string(trendStats))

    crypto.InWallet =
      walletState.CheckWalletStateContainsCurrency(currenciesInWallet, crypto.Name)

      // [
      //   {
      //     "Time_period":"15min",
      //     "Rsi":0,
      //     "RsiStats": {
      //       "Rsi":8.945215690713681,
      //       "Smoothing50":23.662929054564614,
      //       "Smoothing100":24.63802365226256,
      //       "Smoothing250":0
      //     },
      //     "RateChange":0.1668312134390204,
      //     "MovingAverages":   {
      //       "LengthOf10":4087.785799999999,
      //       "LengthOf25":4153.061264,
      //       "LengthOf50":4051.5892340000005,
      //       "LengthOf100":3941.360246000001
      //     }
      //     },
      //     {
      //       "Time_period":"1hr",
      //       "Rsi":0,
      //       "RsiStats":{
      //         "Rsi":90.6777394864817,"Smoothing50":0,"Smoothing100":0,"Smoothing250":0},"RateChange":0.23625131099498817,"MovingAverages":{"LengthOf10":4184.65709,"LengthOf25":4005.598136,"LengthOf50":0,"LengthOf100":0}},{"Time_period":"3hr","Rsi":0,"RsiStats":{"Rsi":0,"Smoothing50":0,"Smoothing100":0,"Smoothing250":0},"RateChange":0.8841767659407507,"MovingAverages":{"LengthOf10":0,"LengthOf25":0,"LengthOf50":0,"LengthOf100":0}},{"Time_period":"24hr","Rsi":0,"RsiStats":{"Rsi":0,"Smoothing50":0,"Smoothing100":0,"Smoothing250":0},"RateChange":-8.965256265633617,"MovingAverages":{"LengthOf10":0,"LengthOf25":0,"LengthOf50":0,"LengthOf100":0}}]

    if crypto.TrendStats != nil {
      fifteenMinTrendStat := crypto.TrendStats[0]
      rsi := fifteenMinTrendStat.RsiStats.Rsi

      if rsi <= 30 {
        crypto.BuyIndicator = true
      }

      if rsi >= 70 {
        crypto.SellIndicator = true
      }
    }

    cryptoCurrencyData = append(cryptoCurrencyData, crypto)
  }

  fmt.Println("Get crypto currency data >> DONE !!")

  return cryptoCurrencyData
}
