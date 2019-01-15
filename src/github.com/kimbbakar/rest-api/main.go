package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"io/ioutil"
	)

type Person struct{
	FirstName string	`json: "firstname: omitempty"`
	LastName string		`json: "lastname: omitempty"`
	ID 		string 		`json: "id"`
}

var People []Person

func GetPerson(w http.ResponseWriter, r *http.Request) {
	
	parameter := mux.Vars(r)
	
//	found := false 

	content, err := ioutil.ReadFile(parameter["id"] )
	if err != nil {
		log.Fatal(err)
	}

	var person Person

	err = json.Unmarshal(content,&person)

	if err != nil {
		log.Fatal(err)
	}


	json.NewEncoder(w).Encode(person)	
	// for _,item:=range People{
	// 	if item.ID == parameter["id"]{ 
	// 		json.NewEncoder(w).Encode(item)
	// 		found = true
	// 		break
	// 	}
	// }

	// if found==false{
				
	// }
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	log.Println("Creating person")
	var person Person
    _ = json.NewDecoder(r.Body).Decode(&person)

	a,_ := json.Marshal(person)

	err := ioutil.WriteFile( person.ID , a,0)

	

//	People = append(People,person)

	if err != nil {
		log.Fatal(err)
	}

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