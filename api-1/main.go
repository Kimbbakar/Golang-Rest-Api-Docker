package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/kimbbakar/rest-api/api-1/mux"
//	"github.com/kimbbakar/rest-api/api-1/TextFileRead"
	"github.com/kimbbakar/rest-api/api-1/InMemoryfile"
	)

type Person struct {
	FirstName string	`json: "firstname: omitempty"`
	LastName string		`json: "lastname: omitempty"`
	ID 		string 		`json: "id"`
}

type IO interface {
	ReadFile(map[string] string)	[]byte
	WriteFile(map[string]interface{} )
	GetPeople()                      []byte
 
}

var db IO  = &InMemoryfile.InMemoryfile{}
//var db IO = &TextFileRead.TextFileRead{}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	
	parameter := mux.Vars(r)

	var person Person
	content:= db.ReadFile(parameter)

	if string (content) == "Person not found"{
		json.NewEncoder(w).Encode("Person not found")	

		return
	}

	err := json.Unmarshal(content,&person)

	if err != nil {
		log.Fatal(err)
	}
 

	json.NewEncoder(w).Encode(person)	
 
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
 
    var person map[string]interface{}

	_ = json.NewDecoder(r.Body).Decode(&person)
  
	db.WriteFile(person)	
 
 
	json.NewEncoder(w).Encode(person) 

}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	log.Println("collecting list")

    var person map[string]interface{}

	_ = json.NewDecoder(r.Body).Decode(&person)
  
	list := db.GetPeople()	
	
 
	json.NewEncoder(w).Encode(string(list)) 

}



func main() {
	router := mux.NewRouter() 
	router.HandleFunc("/person/{id}",GetPerson ).Methods("GET")
	router.HandleFunc("/person",CreatePerson ).Methods("POST")
	router.HandleFunc("/person",GetPeople ).Methods("GET")
	log.Println("Listening...")
	http.ListenAndServe(":8080", router)
}