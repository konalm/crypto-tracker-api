package main

import (
  "fmt"
  // "net/http"
  "encoding/json"
  // "reflect"
  // "io/ioutil"
)

type User struct {
  userId int
  Id int
  Title string
  Body string
}

type Bird struct {
  Species string
  Description string
}

type Foo struct {
  Value interface{}
}


func main() {
  userJson := `{"userId": 1, "id": 1, "title": "test title", "body": "this is the body of user"}`
  var user User
  json.Unmarshal([]byte(userJson), &user)

  fmt.Println("user properties >>>>>>>>>>")
  fmt.Println(user.userId)
  fmt.Println(user.Id)
  fmt.Println(user.Title)
  fmt.Println(user.Body)

  // birdJson := `{"species": "pigeon","description": "likes to perch on rocks"}`
  birdJson := `[{"species":"pigeon","decription":"likes to perch on rocks"},{"species":"eagle","description":"bird of prey"}]`

  var birds []Bird

  json.Unmarshal([]byte(birdJson), &birds)

  fmt.Printf("Birds : %+v", birds)

  // fmt.Println(bird.Description)
  // fmt.Printf("Species: %s, Description: %s", bird.Species, bird.Description)

  // req, err := http.Get("https://jsonplaceholder.typicode.com/posts")
  //
  // if err != nil {
  //   fmt.Println("http request error")
  // }
  //
  //
  // body, err := ioutil.ReadAll(req.Body)
  // fmt.Println("body string >>>>>>>>")
  // fmt.Println( string (body) )
  //
  // var f Foo
  // // s := string(body)
  //
  // json.Unmarshal([]byte(`{ "value": 123 }`), &f);
  //
  //


  //
  //
  //
  // var users []User
  //
  // fmt.Println("users before decode >>>>>>>>>>")
  // fmt.Println(users)
  //
  // // json.NewDecoder(req.Body).Decode(users)
  // // fmt.Println(&users)
  //
  // // defer req.Body.Close()
  //
  // json.NewDecoder(body).Decode(&users)
  // // err = d.Decode(&users)
  //
  //
  //
  // fmt.Println("users after decode >>>>>>>>>>>>>")
  // fmt.Println(users)

  // for _,x := range users {
    // fmt.Println("in users loop >>>>>>>>>>>>")
    // fmt.Println(x)
    //
    // fmt.Println(reflect.TypeOf(x))
    //
    // fmt.Println( json.Marshal(x) )
    //
    // fmt.Printf("%+v\n", x)
    //
    // res2B, _ := json.Marshal(x)
    // fmt.Println(string(res2B))

    //
    // var u User
    // err := json.NewDecoder(x).Decoder(&u)
    //
    // fmt.Println("user after decode >>>>>")
    // fmt.Println(u)
  // }
}
