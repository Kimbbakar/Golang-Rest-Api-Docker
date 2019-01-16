package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/kimbbakar/rest-api/api-1/mux"
	"github.com/kimbbakar/rest-api/api-1/TextFileRead"
//	"github.com/rest-api/kimbbakar/api-1/InMemoryfile/"
	)

type Person struct {
	FirstName string	`json: "firstname: omitempty"`
	LastName string		`json: "lastname: omitempty"`
	ID 		string 		`json: "id"`
}

type IO interface {
	ReadFile(map[string] string)	[]byte
	WriteFile(map[string]interface{} )
 
}

//var db IO  = &InMemoryfile.InMemoryfile{}
var db IO = &TextFileRead.TextFileRead{}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	
	parameter := mux.Vars(r)

	var person Person
	content:= db.ReadFile(parameter)

	err := json.Unmarshal(content,&person)

	if err != nil {
		log.Fatal(err)
	}
 

	json.NewEncoder(w).Encode(person)	
 
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	log.Println("Creating person")

    var person map[string]interface{}

	_ = json.NewDecoder(r.Body).Decode(&person)
  
	db.WriteFile(person)	
 
 
	json.NewEncoder(w).Encode(person) 

}



func main() {
	router := mux.NewRouter() 
	router.HandleFunc("/person/{id}",GetPerson ).Methods("GET")
	router.HandleFunc("/person",CreatePerson ).Methods("POST")
	log.Println("Listening...")
	http.ListenAndServe(":8080", router)
}