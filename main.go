package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"encoding/json"
)


type Person struct {
	ID	string `json:"id,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname string `json:"lastname,omitempty"`
	Address *Address `json:"address,omitempty"`
}

type Address struct {
	City string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request){
	json.NewEncoder(w).Encode(people)
}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request){
	params := mux.Vars(req)
	for index, item := range people {
		if item.ID == params["id"]{
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

func main(){
	router := mux.NewRouter()
	people = append(people, Person{ID: "1", Firstname: "Sai Kiran", Lastname: "Kode", Address: &Address{City: "Somerset", State: "New Jersey"}})
	people = append(people, Person{ID: "2", Firstname: "Niharika", Lastname: "Koneru"})
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":12345", router))
}

/*
go build && go run main.go
first create functions as a base, in func main initialize new router (mux)
then sync these routes and create endpoints and tell the router which type of method it is GET/POST/DELETE
and tell the router which port to listen
as we are not using data bases we use json data
create data in main func and append new data
declare struts and a variable
fill in the function endpoints created and write the logic in functions
use postman chrome extension to test the code 
*/