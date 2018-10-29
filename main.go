package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	Id        string   `json:"Id,omitempty"`
	FirstName string   `json:"FirstName,omitempty"`
	LastName  string   `json:"LastName,omitempty"`
	Address   *Address `json:"Address,omitempty"`
}

type Address struct {
	City  string `json:"City,omitempty"`
	State string `json:"State,omitempty"`
}

var people []Person

func GetPeople(w http.ResponseWriter, r *http.Request) {
	//encoding results in json
	json.NewEncoder(w).Encode(people)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, item := range people {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Person{})
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.Id = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if item.Id == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(people)
	}
}

func main() {
	router := mux.NewRouter()

	//manually adding data to people array struct
	people = append(people, Person{Id: "1", FirstName: "John", LastName: "Doe", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{Id: "2", FirstName: "Giles", LastName: "Doe", Address: &Address{City: "City Z", State: "State Y"}})

	//routes
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
