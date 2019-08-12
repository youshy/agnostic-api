package main

import (
	"fmt"
	"log"
	"net/http"

	// "strconv"
	"encoding/json"
	"strings"

	"github.com/ajg/form"
	"github.com/gorilla/mux"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	CONN_HOST    = "localhost"
	CONN_PORT    = "3000"
	MONGO_DB_URL = "127.0.0.1"
)

var session *mgo.Session
var connectionError error

type User struct {
	Firstname string `json:"firstname" bson:"firstname"`
	Lastname  string `json:"lastname" bson:"lastname"`
	Age       int    `json:"age" bson:"age"`
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

func getUser(res http.ResponseWriter, req *http.Request) {
	queryParams := req.URL.Query()
	// fnOk and lnOk are booleans for checking if the param is there
	firstname, fnOk := queryParams["firstname"]
	lastname, lnOk := queryParams["lastname"]

	fmt.Println(firstname, lastname)

	if fnOk && lnOk {
		queryResult := User{}

		collection := session.DB("mydb").C("users")

		err := collection.Find(bson.M{"firstname": strings.Title(firstname[0]), "lastname": strings.Title(lastname[0])}).One(&queryResult)

		if err != nil {
			log.Print("Got error when reading: ", err)
			return
		}
		json.NewEncoder(res).Encode(queryResult)
	}
}

func postUser(res http.ResponseWriter, req *http.Request) {
	var newUser User

	decoded := form.NewDecoder(req.Body)
	if err := decoded.Decode(&newUser); err != nil {
		log.Print("Could not be decoded: ", err)
		return
	}

	collection := session.DB("mydb").C("users")

	err := collection.Insert(newUser)

	if err != nil {
		log.Print("Problem adding user to db", err)
		return
	}
	log.Print("User added to the db")
	fmt.Fprint(res, "User added to the db")
}

func updateUser(res http.ResponseWriter, req *http.Request) {
	queryParams := req.URL.Query()

	firstname, fnOk := queryParams["firstname"]
	lastname, lnOk := queryParams["lastname"]

	fmt.Println(firstname, lastname)

	if fnOk && lnOk {
		var editUser User

		decoded := form.NewDecoder(req.Body)
		if err := decoded.Decode(&editUser); err != nil {
			log.Print("Could not be decoded: ", err)
			return
		}

		collection := session.DB("mydb").C("users")

		err := collection.Update(bson.M{"firstname": strings.Title(firstname[0]), "lastname": strings.Title(lastname[0])},
			bson.M{"$set": bson.M{
				"firstname": editUser.Firstname,
				"lastname":  editUser.Lastname,
				"age":       editUser.Age,
			}})
		if err != nil {
			log.Print("Problem updating user", err)
			return
		}
		log.Print("Updated user")
		fmt.Fprint(res, "Updated user")
	}
}

func deleteUser(res http.ResponseWriter, req *http.Request) {
	queryParams := req.URL.Query()

	firstname, fnOk := queryParams["firstname"]
	lastname, lnOk := queryParams["lastname"]

	if fnOk && lnOk {
		collection := session.DB("mydb").C("users")

		err := collection.Remove(bson.M{"firstname": strings.Title(firstname[0]), "lastname": strings.Title(lastname[0])})
		if err != nil {
			log.Print("Problem removing user", err)
			return
		}
		log.Print("User removed")
		fmt.Fprint(res, "User removed")
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users", getUser).Methods("GET")
	router.HandleFunc("/users", postUser).Methods("POST")
	router.HandleFunc("/users", updateUser).Methods("PUT")
	router.HandleFunc("/users", deleteUser).Methods("DELETE")
	fmt.Println("Server is running on " + CONN_HOST + ":" + CONN_PORT)
	defer session.Close()
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, router)
	if err != nil {
		log.Fatal("Error starting http server: ", err)
		return
	}
}
