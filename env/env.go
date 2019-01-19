package env

import "os"

func SetEnvVariables() {
  os.Setenv("FOO", "1")
  os.Setenv("BOO", "2")
}
