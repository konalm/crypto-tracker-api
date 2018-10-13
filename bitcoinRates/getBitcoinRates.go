package bitcoinRates

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"sort"
)

type BitcoinRate struct {
	Date         string
	ClosingPrice float64
}

const closingPriceQuery = `
	SELECT date, closing_price 
	FROM bitcoin_rates 
	WHERE min in (0, 15, 30, 45) 
	ORDER BY date 
	DESC LIMIT 15
`

func GetBitcoinRates(w http.ResponseWriter) {
	db := getDBConnection(
		"mysql",
		"root:root@tcp(127.0.0.1:3306)/stelita_dev",
	)
	rows, err := db.Query(closingPriceQuery)
	if err != nil {
		panic(err.Error())
	}
	bitcoinRates := processRates(rows)
	json.NewEncoder(w).Encode(bitcoinRates)
}

func getDBConnection(driverName string, dataSourceName string) *sql.DB {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func processRates(rows *sql.Rows) []BitcoinRate {
	var bitcoinRates []BitcoinRate
	for rows.Next() {
		var bitcoinRate BitcoinRate
		err := rows.Scan(&bitcoinRate.Date, &bitcoinRate.ClosingPrice)
		if err != nil {
			panic(err.Error())
		}
		bitcoinRates = append(bitcoinRates, bitcoinRate)
	}
	sort.SliceStable(bitcoinRates, func(i, j int) bool {
		return bitcoinRates[i].Date < bitcoinRates[j].Date
	})

	return bitcoinRates
}
