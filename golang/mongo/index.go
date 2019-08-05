package main

import (
  "fmt"
  "log"
  "net/http"
  // "strconv"
  "encoding/json"

  "github.com/gorilla/mux"

  mgo "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)

const (
  CONN_HOST       = "localhost"
  CONN_PORT       = "3000"
  MONGO_DB_URL    = "127.0.0.1"
)

var session *mgo.Session
var connectionError error

type User struct {
  Firstname   string  `json:"firstname" bson:"firstname"`  
  Lastname    string  `json:"lastname" bson:"lastname"` 
  Age         int     `json:"age" bson:"age"`   
}

func init() {
  session, connectionError = mgo.Dial(MONGO_DB_URL)
  if connectionError != nil {
    log.Fatal("error connecting to database: ", connectionError)
  }
  log.Print("Connected to MongoDB")
  session.SetMode(mgo.Monotonic, true)
}

func getUsers(res http.ResponseWriter, req *http.Request) {
  var queryResult []User

  collection := session.DB("mydb").C("users")

  err := collection.Find(bson.M{}).All(&queryResult)
  if err != nil {
    log.Print("Got error when reading: ", err)
    return
  }
  json.NewEncoder(res).Encode(queryResult)
}

func main() {
  router := mux.NewRouter()
  router.HandleFunc("/users", getUsers).Methods("GET")
  fmt.Println("Server is running on "+CONN_HOST+":"+CONN_PORT)
  defer session.Close()
  err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, router)
  if err != nil {
    log.Fatal("Error starting http server: ", err)
    return
  }
}
