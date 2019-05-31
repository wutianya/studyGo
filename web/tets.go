package main

import(
	"crypto/rand"
	"crypto/sha1"
	// "database/sql"
	"fmt"
	"log"
	"time"
)

func main() {
	
	a := createUUID()
	b := Encrypt("123456")
	fmt.Println(a)
	fmt.Println(b)
	c := time.Now()
	fmt.Println(c.Format("Jan 2, 2006 at 3:04"))
}
// creat a random UUID with from RFC 4122
func createUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}
	u[8] =(u[8] | 0x40) & 0x7F
	u[6] = (u[6] & 0xF) |(0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

// hash plaintext with SHA-1
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x",sha1.Sum([]byte(plaintext)))
	return
}