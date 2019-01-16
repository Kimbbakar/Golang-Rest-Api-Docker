package TextFileRead

import (
	"encoding/json"
	"io/ioutil" 
	"log"
	"os"
	)

type TextFileRead struct{

	
}

func (T *TextFileRead) ReadFile(parameter map[string] string) []byte {
	content, err := ioutil.ReadFile(parameter["id"]+".txt" )
	if err != nil {
		log.Fatal(err)
	}	

	return content
} 


func (T *TextFileRead) WriteFile(content map[string] interface{} ) {

	b,_:=json.Marshal(content)

	a,ok:=content["id"].(string)

	if ok==false{
		log.Fatal("ID not found")
	}

    var _, err = os.Stat(a+".txt")

    // create file if not exists
    if os.IsNotExist(err) {
        _, err = os.Create(a+".txt")
		if err != nil {
			log.Fatal(err.Error())
		} 
    }

	err = ioutil.WriteFile( a+".txt" ,[] byte(b),0)
 

	if err != nil {
		log.Fatal(err)
	} 

}