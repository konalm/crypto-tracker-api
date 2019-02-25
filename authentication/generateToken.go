package authentication

import (
  "os"
  "time"
  "fmt"
  "github.com/dgrijalva/jwt-go"
)

type TokenClaims struct {
  UserId int
  jwt.StandardClaims
}

func generateJWT(userId int) string {
  signedKey := []byte(os.Getenv("TOKEN_SECRET"))
  tokenExpireTime := time.Now().Add(time.Hour * 48).Unix()

  fmt.Println("user id >>>")
  fmt.Println(userId)

  // Create the claims
  claims := TokenClaims {
    userId,
    jwt.StandardClaims {
      ExpiresAt: tokenExpireTime,
      Issuer: "Stelita API",
    },
  }

  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  signedToken, err := token.SignedString(signedKey)
  if err != nil {
    panic(err.Error())
  }

  return signedToken
}
