package main

import (
  "log"
  "net/http"
  "os"
  "github.com/joho/godotenv"
  "github.com/gorilla/handlers"
  "crypto-tracker-api/router"
  "crypto-tracker-api/cronJobs"
)

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  appRouter := router.Index()

	cronJobs.HandleBitcoinRate()

  originsAllowed := handlers.AllowedOrigins([]string{os.Getenv("ALLOWED_CLIENT")})
  headersAllowed := handlers.AllowedHeaders([]string{"X-Requested-With"})
  methodsAllowed := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

  log.Fatal(http.ListenAndServe(":8484", handlers.CORS(originsAllowed, headersAllowed, methodsAllowed)(appRouter)))
}
