/*
package main

import (
	"fmt"
)

type A struct {
	Name string
}
type B struct {
	Name string
}

func (getA *A) echo() {
	fmt.Println("A")
	getA.Name = "Ailis"
}
func (getB B) echo() {
	fmt.Println("B")
	getB.Name = "zs"
}
func main() {
	a := A{}
	a.echo()	
	fmt.Println(a.Name)
	b := B{}
	b.echo()
	fmt.Println(b.Name)
}
*/

package main
import (
	"fmt"
)

type TZ int
type Add struct {
	Num int
}

func (a *Add) add(num int) {
	a.Num += num
}
func (add *TZ) Increase(num int) {
	*add += TZ(num)
}



func main() {
	 var a TZ
	 a.Increase(100)
	 fmt.Println(a)
	 b := Add{}
	 b.add(50)
	 fmt.Println(b.Num)
}