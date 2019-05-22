/*
// gob1.go
package main

import (
	"cmd/go/internal/str"
	"bytes"
	"fmt"
	"encoding/gob"
	"log"
)

type P struct {
	X, Y, Z int
	Name	string
}

type Q struct {
	X, Z	*int32
	Name	string
}

func main() {
	var netwrok bytes.Buffer
	enc := gob.NewEncoder(&netwrok)
	dec := gob.NewDecoder(&netwrok)

	err := enc.Encode(P{3,4,5,"Sam"})
	if err != nil {
		log.Fatal("encode error: ", err)
	}

	var q Q
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error: ", err)
	}
	fmt.Printf("%q: {%d,%d}\n",q.Name, *q.X, *q.Z)
}
*/
/*
// 使用gob编码
package main

import (
	"encoding/gob"
	"log"
	"os"
)

type Address struct {
	Type	string
	City	string
	Country	string
}

type VCard struct {
	FirstName	string
	LastName	string
	Address		[]*Address
	Remark		string
}

var content	string

func main() {
	pa := &Address{"private", "HangZhou", "China"}
	wa := &Address{"work", "ChengDu", "China"}
	vc := VCard{"Zhang", "SanFeng", []*Address{pa,wa}, "none"}

	file, _ := os.OpenFile("vcard.gob", os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	enc := gob.NewEncoder(file)
	err := enc.Encode(vc)

	if err != nil {
		log.Println("Error in encoing gob")
	}
}

*/
package main

import (
	// "bufio"
	"encoding/gob"
	"fmt"
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
	Addresses []*Address
	Remark    string
}

var vc VCard

func main() {
	f, err := os.Open("vcard.gob")
	if err != nil {
		fmt.Println("open file error")
	}
	defer f.Close()
	// inReader := bufio.NewReader(f)
	dec := gob.NewDecoder(f)
	err = dec.Decode(&vc)
	if err != nil {
		fmt.Println("file parse error")
	}
	fmt.Println(vc)

}
