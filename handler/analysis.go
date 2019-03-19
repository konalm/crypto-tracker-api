package handler

import (
  "fmt"
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
  "stelita-api/analysis"
)


/**
 *
 */
func GetAnalysis(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Get analysis !!")

  analysis := analysis.GetAnalysis()

  json.NewEncoder(w).Encode(analysis)
}


/**
 *
 */
func GetCryptoCurrencyAnalysis(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Get crypto currency analysis !!")

  params := mux.Vars(r)

  cryptoCurrencySymbol := params["crypto_symbol"]
  fmt.Println(cryptoCurrencySymbol)

  cryptoCurrencyAnalysis := analysis.GetCryptoCurrencyAnalysis(cryptoCurrencySymbol)

  json.NewEncoder(w).Encode(cryptoCurrencyAnalysis)
}

/**
 *
 */
func GetAnalysisItem(w http.ResponseWriter, r *http.Request) {
  fmt.Println("get analysis item !!")

  params := mux.Vars(r)
  id := params["id"]

  fmt.Println("analysis item id >>>")
  fmt.Println(id)

  analysisItem := analysis.GetAnalysisItem(id)

  json.NewEncoder(w).Encode(analysisItem)
}
