/*
package main

import (
	"fmt"
)

var (
	a int = 1
	b     = 2
)

func main() {
	var a int
	fmt.Println(a, b)
}
*/
/*
package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	var goos string = runtime.GOOS
	fmt.Printf("The operating system is: %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
	h,_ := os.Hostname()
	fmt.Println(h)
}
*/
/*
package  main

var a = "G"

func main() {
	n()
	m()
	n()
}
func n() { print(a) }
func m() {
	a = "O"
	print(a)
}
*/
/*
package main

var a string
func main() {
	a = "G"
	print(a)
	f1()
}
func f1() {
	a := "O"
	print(a)
	{
		print(a)
	}
	f2()
}
func f2() {
	print(a)
}
*/
/*
package main
import "fmt"

func main()  {
	a := 31.1415925
	fmt.Printf("a is: %6.2v",a)
}
*/
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 10; i++ {
		a := rand.Int()
		fmt.Printf("%d / ", a)
	}
}
