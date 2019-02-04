package authentication

import (
  "fmt"
  "net/http"
  "github.com/dgrijalva/jwt-go"
  "os"
)

/**
 *
 */
func AuthCheck(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Auth Check")

  authToken := r.Header.Get("Authorization")
  tokenString := authToken
  mySigningKey := []byte(os.Getenv("TOKEN_SECRET"))


  // Parse takes the token string and a function for looking up the key. The latter is especially
  // useful if you use multiple keys for your application.  The standard is to use 'kid' in the
  // head of the token to identify which key to use, but the parsed token (head and claims) is provided
  // to the callback, providing flexibility.
  token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
      // Don't forget to validate the alg is what you expect:
      if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
          return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
      }

      // hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
      return mySigningKey, nil
  })

  if !token.Valid {
    w.WriteHeader(406)
    w.Write([]byte("Not Authorized"))
    return
  }

  w.WriteHeader(200)
  w.Write([]byte("Authorized"))
}
