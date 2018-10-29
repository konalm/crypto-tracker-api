package main

import (
  "log"
  "net/http"
  "os"
  // "fmt"
  "github.com/joho/godotenv"
  "github.com/gorilla/handlers"
  "stelita-api/router"
  "stelita-api/cronJobs"
  // "database/sql"
  // _ "github.com/go-sql-driver/mysql"
  // "stelita-api/structs"
  // "stelita-api/cryptoRatesController"
  // "stelita-api/abstractRatesByTimePeriod"
  "stelita-api/rsi"
)


func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  rsi.HandleRsi()
  // return

  appRouter := router.Index()

	cronJobs.HandleBitcoinRate()
  cronJobs.HandleRankedCryptoCurrencyUpdate()

  originsAllowed := handlers.AllowedOrigins([]string{os.Getenv("ALLOWED_CLIENT")})
  headersAllowed := handlers.AllowedHeaders([]string{"X-Requested-With"})
  methodsAllowed := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

  log.Fatal(http.ListenAndServe(":8484", handlers.CORS(originsAllowed, headersAllowed, methodsAllowed)(appRouter)))
}
