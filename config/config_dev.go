// +build dev

package config

const (
  PORT = "8383"
  ALLOWED_CLIENT = "http://localhost:8080"
  ALLOWED_CLIENT_2 = "http://localhost:8081"
  LOG_FILE = "/var/log/stelita/dev.log"
  DB = "stelita_dev"
  ANALYSIS_API_URL = "http://138.68.167.173:3000"
  EVENT_REPORTER_API_URL = "http://138.68.167.173:3005"
)
