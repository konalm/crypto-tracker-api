package authentication

import (
  "os"
  "time"
  "github.com/dgrijalva/jwt-go"
)

func generateJWT() string {
  signedKey := []byte(os.Getenv("TOKEN_SECRET"))
  tokenExpireTime := time.Now().Add(time.Hour * 48).Unix()

  // Create the claims
  claims := &jwt.StandardClaims {
    ExpiresAt: tokenExpireTime,
    Issuer: "Stelita API",
  }

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  signedToken, err := token.SignedString(signedKey)
  if err != nil {
    panic(err.Error())
  }

  return signedToken
}
