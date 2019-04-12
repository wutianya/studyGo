/*
package main
import "fmt"

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
