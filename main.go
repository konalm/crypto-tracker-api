package main

import (
  "log"
  "net/http"
  "github.com/gorilla/handlers"
  "stelita-api/router"
  "stelita-api/cronJobs"
  "stelita-api/env"
  "stelita-api/config"
)


func main() {
  env.SetEnvVariables()

  appRouter := router.Index()

	cronJobs.HandleBitcoinRate()
  cronJobs.HandleRankedCryptoCurrencyUpdate()

  originsAllowed := handlers.AllowedOrigins([]string{config.ALLOWED_CLIENT})
  headersAllowed := handlers.AllowedHeaders([]string{"X-Requested-With"})
  methodsAllowed := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

  log.Fatal(http.ListenAndServe(":" + config.PORT, handlers.CORS(originsAllowed, headersAllowed, methodsAllowed)(appRouter)))
}
