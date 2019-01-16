package TextFileRead

import (
	"io/ioutil" 
	"log"
	)

type TextFileRead struct{

	
}

func (T *TextFileRead) ReadFile(parameter map[string] string) []byte {
	content, err := ioutil.ReadFile(parameter["id"] )
	if err != nil {
		log.Fatal(err)
	}	

	return content
} 