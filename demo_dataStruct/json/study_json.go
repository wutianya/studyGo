// json.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}
type VCard struct {
	FirstName string
	LastName  string
	Address   []*Address
	Remark    string
}

func main() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "belgium"}
	vc := VCard{"Jan", "kersschot", []*Address{pa, wa}, "none"}
	js, _ := json.Marshal(vc)
	fmt.Printf("JSON format: %s\n", js)

	b := []byte(`{"Name": "Wednesday", "Age": 6, "Parents": ["Gomez", "Morticia"]}`)
	var f interface{}
	e := json.Unmarshal(b, &f)
	if e == nil {
		fmt.Println("f is: ", f.(map[string]interface{}))
	}

	file, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := json.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		log.Println("Error in encoding json")
	}

}
