// +build prod

package config

const (
  PORT = "8484"
  ALLOWED_CLIENT = "https://stelita.app"
  ALLOWED_CLIENT_2 = "https://monitor.stelita.app"
  LOG_FILE = "/var/log/stelita-prod.log"
  DB = "stelita_prod"
)
