package authentication

import (
  "fmt"
  "golang.org/x/crypto/bcrypt"
  "net/http"
  // "github.com/gorilla/mux"
  "encoding/json"
  "stelita-api/user"
)

type response struct {
  AccessToken string
}

/**
 * If credentials are authenticated return JWT access token
 */
func Login(w http.ResponseWriter, r *http.Request) {
  username := r.FormValue("username")
  password := r.FormValue("password")

  user := user.GetUserByUsername(username)
  compareHashedPassword := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

  if compareHashedPassword != nil {
    w.WriteHeader(404)
    w.Write([]byte("Username or password does not match"))
    return
  }

  token := generateJWT()

  var loginResponse = response {
    AccessToken: token,
  }

  json.NewEncoder(w).Encode(loginResponse)
}


/**
 *
 */
func ProtectedResource(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Protected Resource")

  w.Write([]byte("Protected Resource"))


}
