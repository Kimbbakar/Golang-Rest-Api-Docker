package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	)

type Person struct{
	FirstName string	`json: "firstname: omitempty"`
	LastName string		`json: "lastname: omitempty"`
	ID 		string 		`json: "id"`
}

var People []Person

func GetPerson(w http.ResponseWriter, r *http.Request) {
	
	parameter := mux.Vars(r)
	
	found := false 

	for _,item:=range People{
		if item.ID == parameter["id"]{ 
			json.NewEncoder(w).Encode(item)
			found = true
			break
		}
	}

	if found==false{
				
	}
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	log.Println("Creating person")
	var person Person
    _ = json.NewDecoder(r.Body).Decode(&person)
	
	People = append(People,person)

	json.NewEncoder(w).Encode(person)
	json.NewEncoder(w).Encode(People)

}



func main() {
	router := mux.NewRouter()
	router.HandleFunc("/person/{id}",GetPerson ).Methods("GET")
	router.HandleFunc("/person",CreatePerson ).Methods("POST")
	log.Println("Listening...")
	http.ListenAndServe(":8080", router)
}