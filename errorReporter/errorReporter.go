package errorReporter

import (
  "fmt"
  "log"
  "time"
  "os"
  "net/smtp"
  "stelita-api/config"
)

/**
 *
 */
func ReportError(errorMessage string) {
  fmt.Println("Report Error >>>")
  fmt.Println(errorMessage)

  // currentDateTime := time.Now()

  /* log message */
  f, err := os.OpenFile(config.LOG_FILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
  if err != nil {
    log.Println(err)
  }
  defer f.Close()

  logger := log.New(f, "", log.LstdFlags)
  logger.Println(">> ", errorMessage)
}


func send(body string) {
	from := "connorlloydmoore@gmail.com"
  pass := "tvaupqoyhqqzwqkq"
	to := "connor@codegood.co"
  subject := "Stelita API Error"
  dateTime := time.Now()

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		dateTime.Format("2006-01-02 15:04:05") + "\n\n " + body

	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}

	log.Print("sent email")
}
