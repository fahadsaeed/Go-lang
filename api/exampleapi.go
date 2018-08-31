package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Person struct {
	ID        string   `json:"id,omitempty"`
	FirstName string   `json:firstname,omitempty"`
	LastName  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:state,omitempty"`
}

var people []Person

func GetPeopleEndpount(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}
func GetPersonEndpount(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	fmt.Println("params.....", params)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func CreatePersonEndpount(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)

	person.ID = params["id"]

	fmt.Println("person post ", person)

	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}
func DeletePersonEndpount(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(people)

}

func main() {
	router := mux.NewRouter()

	people = append(people, Person{ID: "1", FirstName: "Fahad", LastName: "Saeed", Address: &Address{City: "Karachi", State: "Sindhi"}})
	people = append(people, Person{ID: "2", FirstName: "Sami", LastName: "Saeed"})

	router.HandleFunc("/people", GetPeopleEndpount).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpount).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndpount).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndpount).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":1234", router))

}
