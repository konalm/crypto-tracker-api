package main

import (
  "log"
  "net/http"
  "fmt"
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

  fmt.Println("Allowed Client >>>")
  fmt.Println(config.ALLOWED_CLIENT)

  originsAllowed := handlers.AllowedOrigins([]string{config.ALLOWED_CLIENT, config.ALLOWED_CLIENT_2})
  headersAllowed := handlers.AllowedHeaders([]string{"pkm-client", "X-Requested-With", "content-type", "Authorization", "Access-Control-Allow-Origin"})
  methodsAllowed := handlers.AllowedMethods([]string{"POST", "HEAD", "GET", "PUT", "OPTIONS"})

  log.Fatal(http.ListenAndServe(":" + config.PORT, handlers.CORS(originsAllowed, headersAllowed, methodsAllowed)(appRouter)))
}
