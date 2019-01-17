package InMemoryfile

import (
	"encoding/json" 
	"log" 
	)

type Person struct {
	FirstName string	`json: "firstname: omitempty"`
	LastName string		`json: "lastname: omitempty"`
	ID 		string 		`json: "id"`
}


type InMemoryfile struct{
	People []Person		
}

func (T *InMemoryfile) ReadFile(parameter map[string] string) []byte {
	
	for _,val:= range T.People{
		if val.ID==parameter["id"]{
			b,_:=json.Marshal(val)
			return []byte(b)
		}
	}
	
	return []byte("Found nothing")
		
} 


func (T *InMemoryfile) WriteFile(content map[string] interface{} ) {

	b,_:=json.Marshal(content)

	var p Person
	err :=json.Unmarshal(b,&p)
	T.People=append(T.People,p)

	if err!=nil{
		log.Fatal(err)
	}
}

func (T *InMemoryfile) GetPeople() []byte{
 
	var content = ""
	for _,val:= range T.People{
		b,_:=json.Marshal(val)
		content+=string(b)
	}     
		
	if len(content)==0{
		content = "Person not found"
	}

	return []byte(content)
}