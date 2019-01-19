package Mongodb

import (
	"github.com/globalsign/mgo" 
	"encoding/json"
	"gopkg.in/mgo.v2/bson"
	"log"
)

const (
	url = "localhost"
)

type Person struct {
	FirstName string	`bson: "firstname: omitempty"`
	LastName string		`bson: "lastname: omitempty"`
	ID 		string 		`bson: "id"`
}

type Mongodb struct { 
	session *mgo.Session
	db 		*mgo.Database
}

func (T *Mongodb)  Init()   {
	T.session, _ = mgo.Dial(url)

	T.db = T.session.DB("People")
	
	if T.db == nil {
		log.Println("db People not found, exiting ")
		return
	}
}

func (T *Mongodb)  Close()   {
	T.session.Close()
}

func (T *Mongodb) DatabaseName () string {
	return "Mongo DB"
}




func (T *Mongodb) ReadFile(parameter map[string] string) []byte {
	c := T.db.C("info")
	var result Person
	count, _ := c.Find(bson.M{"id": parameter["id"] }).Count()
	if count==0{
		return []byte ("Person not found")
	}

	c.Find(bson.M{"id": parameter["id"] }).One(&result)

	content,_ := json.Marshal(result)

	return content
} 

func (T *Mongodb) WriteFile(content map[string] interface{} ) {
	c := T.db.C("info")
	var result Person


	b,_:=json.Marshal(content)
	_=json.Unmarshal(b,&result)
	
	c.Insert(result)
}

func (T *Mongodb) GetPeople() []byte{
 

	c := T.db.C("info")

	count, _ := c.Find(nil).Count()
	if count==0{
		return []byte ("Person not found")
	}

	var results []map[string] interface{}
	_ = c.Find(nil).All(&results)

	content,_ := json.Marshal(results)

	return content

} 

func (T *Mongodb) Update(id string, person map[string]string ) []byte {
	c := T.db.C("info")

	var content = "Person not found"
	
  
	if id != person["id"]{
		return []byte ("url id and body id not same")
	}

	b,_:=json.Marshal(person) 

	var tmp Person
	_ = json.Unmarshal(b,&tmp)

	count, _ := c.Find(bson.M{"id": id}).Count()
	if count==0{
		return []byte (content)
	}

	c.Update(bson.M{"id": id}, bson.M{"$set": tmp } )
 
	return b
} 