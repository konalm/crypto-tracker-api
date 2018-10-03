package main

import (
  "fmt"
  "time"
  // "strings"
)


func main() {
  // originalString := "2018-09-29T10:40:34.6013152Z"
  s := "2018-09-29 10:40:34.6013152";
  // x := strings.Replace(originalString, "T", " ", -1)

  // var replacer = strings.NewReplacer("T", " ", "Z", "")
  // y := replacer.Replace(x)

  fd, err := time.Parse("2006-01-02 15:04:05", s)

  if err != nil {
    fmt.Printf("ERROR >>>")
    fmt.Println(err.Error())
  }

  fmt.Println("formatted date >>>>>>>")
  fmt.Println(fd)

  t := time.Now()

  fmt.Println("time >>>")
  fmt.Println(t)

  fmt.Println( t.Minute() )
}



// request.Header.Set("X-CoinAPI-Key", `B59AE558-D76F-4AB3-ACE8-D4DFB30D6C59`)
