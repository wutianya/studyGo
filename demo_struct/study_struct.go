/*
package main
import (
	"github.com/derekparker/delve/pkg/proc/native"
	"debug/pe"
	"fmt"
)

type person struct {
	Name string
	Age	int
}

func main() {
	a := &person{
		Name: "bob",
		Age: 12,
	}
	a.Name = "sb"
	fmt.Println(a)
	A(a)
	fmt.Println(a)
}
func A(per *person) { // 地址调用
	per.Age = 18
	fmt.Println("A", per)
}
*/
/*
// 匿名结构
package main
import "fmt"
func main() {
	a := &struct {
		Name string
		Age int
	}{
		Name: "joe",
		Age: 11,
	}
	fmt.Println(a)
}
*/
/*
// 匿名结构拓展
package main

import "fmt"
type person struct{
	Name string
	Age int
	Contact struct {
		Phone, City string
	}
}
func main()  {
	a := &person{Name: "job", Age: 18}
	a.Contact.City = "Hangzhou"
	a.Contact.Phone = "18202381345"
	fmt.Println(a)
}
*/
/*
// 匿名字段
package main
import "fmt"
type person struct {
	string
	int
}
func main() {
	// a := &person{"job", 19}
	a := &person{}
	a.string = "Job"
	a.int = 18
	fmt.Println(a.string)
}
*/
/*
// 组合类似于继承
package main

import "fmt"

type human struct {
	Sex int
	Age int
}
type teacher struct {
	human
	Name string
	Age  int
}
type student struct {
	human
	Name string
	Age  int
}

func main() {
	a := teacher{Name: "joe", Age: 19, human: human{Sex: 0}}
	b := student{Name: "bob", Age: 18, human: human{Sex: 1}}
	b.human.Age = 88
	fmt.Println(a, b)
	fmt.Println(b.Age)
	fmt.Println(b.human.Age)
}
*/
/*
// 2019-04-19 重新学习struct
package main
import (
	"fmt"
)

func main() {
	// 初始化struct
	i := test{Name:"Jack",Gender: "M"}
	i1 := &test{Name:"Jack",Gender: "M"}
	i.Name = "Bob"
	fmt.Println(i)
	fmt.Println(i1)

}
// define struct
type test struct {
	Name string
	Age int
	Gender string
}
*/
/*
package main

import (
	"fmt"
	"strings"
)

type Person struct {
	firstName string
	lastName  string
}

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}
func main() {
	// 三种初始化struct方法pers1 pers2 pers3
	var pers1 Person
	pers1.firstName = "huang"
	pers1.lastName = "di"
	upPerson(&pers1)
	fmt.Printf("The name of the person is %s %s\n", pers1.firstName, pers1.lastName)

	pers2 := new(Person)
	pers2.firstName = "huang"
	pers2.lastName = "di" // --> (*pers2).lastName = "di"
	upPerson(pers2)
	fmt.Printf("The name of the person is %s %s\n", pers2.firstName, pers2.lastName)

	pers3 := &Person{firstName:"huang",lastName:"di"}
	upPerson(pers3)
	fmt.Printf("The name of the person is %s %s\n", pers3.firstName, pers3.lastName)
}
*/
/*
// 结构体转换
package main

import "fmt"

type number struct {
	f float64
}
type nr number

func main() {
	a := number{5.0}
	a.f = 88
	b := nr{5.0}
	fmt.Println(a, b)
	var c = number(b)
	fmt.Println(a, b, c)
}
*/

// 匿名struct 

package main 
import "fmt"
 
type Person struct {
	Name string
	Gender string
	Human
}

type Human struct {
	age int 

	 
}
func main() {
	Jack := Person{Name:"Jack",Gender:"F",Human: Human{age: 18}}
	fmt.Println(Jack.age)
}
