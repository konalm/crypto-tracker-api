package main

import (
  "log"
  "net/http"
  "github.com/gorilla/handlers"
  "stelita-api/router"
  "stelita-api/cronJobs"
  "stelita-api/env"
  "stelita-api/config"
  "fmt"
  "os"
)


func main() {
  fmt.Println("config port >>>")
  fmt.Println(config.PORT)

  fmt.Println("config allowed client >>>")
  fmt.Println(config.ALLOWED_CLIENT)


  /* log message */
  f, err := os.OpenFile("text.log",
  	os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
  if err != nil {
  	log.Println(err)
  }
  defer f.Close()

  logger := log.New(f, "prefix", log.LstdFlags)
  logger.Println("text to append")
  logger.Println("more text to append")
  logger.Println(config.ALLOWED_CLIENT)


  env.SetEnvVariables()

  appRouter := router.Index()

	cronJobs.HandleBitcoinRate()
  cronJobs.HandleRankedCryptoCurrencyUpdate()

  originsAllowed := handlers.AllowedOrigins([]string{config.ALLOWED_CLIENT})
  headersAllowed := handlers.AllowedHeaders([]string{"X-Requested-With"})
  methodsAllowed := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

  log.Fatal(http.ListenAndServe(":" + config.PORT, handlers.CORS(originsAllowed, headersAllowed, methodsAllowed)(appRouter)))
}
