package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/kimbbakar/rest-api/api-1/mux"
	"github.com/kimbbakar/rest-api/api-1/TextFileRead"
//	"github.com/kimbbakar/rest-api/api-1/InMemoryfile"
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
	Update(string,map[string] string) []byte		
 
}

//var db IO  = &InMemoryfile.InMemoryfile{}
var db IO = &TextFileRead.TextFileRead{}

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

func UpdatePerson(w http.ResponseWriter, r *http.Request) {
	
	parameter := mux.Vars(r)

    var person map[string]string 

	_ = json.NewDecoder(r.Body).Decode(&person)


	content := db.Update(parameter["id"],person )

	if (string(content)=="Person not found" || string(content)=="url id and body id not same"){
		log.Println(string (content))
		json.NewEncoder(w).Encode( string (content) )	
	} else{
		var tmp Person
		log.Println(string(content) )
		json.Unmarshal(content,&tmp)
	
		json.NewEncoder(w).Encode(tmp)	
	
	} 
}



func main() {
	router := mux.NewRouter() 
	router.HandleFunc("/person/{id}",GetPerson ).Methods("GET")
	router.HandleFunc("/person/{id}",UpdatePerson ).Methods("POST")
	router.HandleFunc("/person",CreatePerson ).Methods("POST")
	router.HandleFunc("/person",GetPeople ).Methods("GET")
	log.Println("Listening...")
	http.ListenAndServe(":8080", router)
}