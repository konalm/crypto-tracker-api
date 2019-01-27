package main

import (
  "fmt"
  "time"
  "github.com/dgrijalva/jwt-go"
)

func main() {
  fmt.Println("JWT TOKENS")

  mySigningKey := []byte("super_secret")

  expireTime := time.Now().Add(time.Hour * 48).Unix()

  // Create the Claims
  claims := &jwt.StandardClaims{
    ExpiresAt: expireTime,
    Issuer: "test",
  }


  token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
  ss, err := token.SignedString(mySigningKey)
  if err != nil {
    panic(err.Error())
  }

  fmt.Println("token >>>>")
  fmt.Println(ss)

  validateToken(ss)
}


func validateToken(tokenString string) {
  fmt.Println("Validate Token")

  mySigningKey := []byte("super_secret")

  // Parse takes the token string and a function for looking up the key. The latter is especially
  // useful if you use multiple keys for your application.  The standard is to use 'kid' in the
  // head of the token to identify which key to use, but the parsed token (head and claims) is provided
  // to the callback, providing flexibility.
  token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
      // Don't forget to validate the alg is what you expect:
      if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
          return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
      }

      // hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
      return mySigningKey, nil
  })

  if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
      fmt.Println(claims["foo"], claims["nbf"])
  } else {
      fmt.Println(err)
  }

  fmt.Println("END OF CHECK >>")
}
