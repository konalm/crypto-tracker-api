package logger

import (
  "fmt"
  "log"
  "os"
  "os/exec"
  "stelita-api/config"
  "stelita-api/db"
)

/**
 *
 */
func WriteLog(message string, tag string) {
  fmt.Println("write log")
  fmt.Println(message)

  message = "Stelita API -- " + message

  /* log message */
  f, err := os.OpenFile(config.LOG_FILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
  if err != nil {
    log.Println(err)
  }
  defer f.Close()

  fmt.Println("Write log >>>")
  fmt.Println(message)

  logger := log.New(f, "", log.LstdFlags)
  logger.Println(message)

  // StoreLog(message, "")
}

/**
 *
 */
func StoreLog(message string, tag string) {
  fmt.Println("store log")
  fmt.Println(message)

  conn := db.Conn()
  defer conn.Close()

  query := `INSERT INTO logs (id, message, tag) VALUES (?,?,?)`
  stmt, err := conn.Prepare(query)
  if err != nil {
    panic(err.Error())
  }

  uuidOut, err := exec.Command("uuidgen").Output()
  if err != nil {
    panic(err.Error())
  }
  uuid := fmt.Sprintf("%s", uuidOut)

  _, err = stmt.Exec(uuid, message, tag)
  if err != nil {
    panic(err.Error())
  }

  defer stmt.Close()
  defer conn .Close()
}
