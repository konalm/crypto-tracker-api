package structs


type ProcessList struct {
  Id int
  User string
  Host string
  Db *string
  Command string
  Time int
  State *string
  Info *string
}
