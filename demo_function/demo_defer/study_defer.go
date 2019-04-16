// defer
/*
package main

import "fmt"

func main() {
	// f1()
	// a()
	f3()
}
func f1() {
	fmt.Println("f1() start!")
	defer f2()
	fmt.Println("f1() end!")
}
func f2() {
	fmt.Println("f2()")
}
func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}
func f3() {
	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}
}
*/
/*
package main
import "fmt"
func main() {
	doDBOperations()
}
func connectToDB() {
	fmt.Println("ok, connected to db")
}
func disconnectFromDB() {
	fmt.Println("ok, disconnectd from db")
}
func doDBOperations() {
	connectToDB()
	fmt.Println("Defering the database disconnect.")
	defer disconnectFromDB()
	fmt.Println("Doing some DB operations ...")
	fmt.Println("Oops! some crash or network error ...")
	fmt.Println("Returning from function here!")
	return
}
*/
/*
// 使用 defer 语句来记录函数的参数与返回值
package main

import (
	"io"
	"log"
)

func main() {
	func1("Go")
}
func func1(s string) (n int, err error) {
	defer func() {
		log.Printf("func(%q) = %d, %v", s, n, err)
	}()
	return 7, io.EOF
}
*/
/*
// 递归函数
package main

import "fmt"

func main() {
	res := 0
	for i := 0; i <= 10; i++ {
		res = fib(i)
		fmt.Printf("fib(%d) is: %d\n", i, res)
	}
}
func fib(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fib(n-1) + fib(n-2)
	}
	return
}
*/
/*
package main

import (
	"fmt"
)
func main() {
	// fmt.Printf("%d is even: is %t\n", 16, even(16)) // 16 is even: is true
	// fmt.Printf("%d is odd: is %t\n", 17, odd(17))
	// 17 is odd: is true
	// fmt.Printf("%d is odd: is %t\n", 18, odd(18))
	// 18 is odd: is false
	test(10)
}
func even(nr int) bool {
	if nr == 0 {
		return true
	}
	return odd(RevSign(nr) - 1)
}

func odd(nr int) bool {
	if nr == 0 {
		return false
	}
	return even(RevSign(nr) - 1)
}

func RevSign(nr int) int {
	if nr < 0 {
		return -nr
	}
	fmt.Println(nr)
	return nr
}

func test(n int) {
	if n < 0 {
		return
	}
	fmt.Println(n)
	n--
	test(n)
}
*/
/*
// 函数回调
package main

import (
	"fmt"
)

func main() {
	callback(1, Add)
	fmt.Printf("%T",callback)
}
func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}
func callback(y int, f func(int, int)) {
	f(y, 2)
}
*/
// 闭包 
package main
import "fmt"
func main() {
	// 匿名函数
	f :=func(i int) {
		fmt.Println(i)
	}
	f(1)
	
}