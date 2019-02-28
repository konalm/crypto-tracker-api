package email

import (
  "fmt"
  "log"
  "os"
  "net/smtp"
)

/**
 *
 */
func send(message string) {
  fmt.Println("init email")

  from := os.Getenv("EMAIL_FROM")
  pass := os.Getenv("EMAIL_PASS")
  to := os.Getenv("EMAIL_TO")

  subject := "Stelita RSI Notification"
  // dateTime := time.Now()

  msg := "From: " + from + "\n" +
    "To: " + to + "\n" +
    "Subject: " + subject + "\n\n" +
    message

  err := smtp.SendMail("smtp.gmail.com:587",
    smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
    from, []string{to}, []byte(msg))

  if err != nil {
    log.Printf("smtp error: %s", err)
    return
  }
}

/**
 *
 */
func TradeCurrencyNotification(currency string, rsi float64, recommendation string) {
  fmt.Println("email >> sell currency notification")

  rsiString := fmt.Sprintf("%f", rsi)
  message := "Currency: " + currency + "\n" +
    "RSI: " + rsiString + "\n" +
    "Stelita recommends " + recommendation

  send(message)
}
