// defer
/*
package main

import (
	"go/constant"
	"time"
	"fmt"
)

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
/*
package main
import "fmt"
func main() {
	// 匿名函数
	func(i int) {
		fmt.Println(i)
	}(10)
}
*/
/*
package main
import "fmt"
func main() {
	fv := func(){
		fmt.Println("hello,world")
	}
	fv()
}
*/
/*
package main
import "fmt"
func f()(ret int) {
	defer func(){
		ret++
	}()
	return 1
}
func main() {
	fmt.Println(f())
}
*/
/*
// 应用闭包：将函数作为返回值
package main
import "fmt"
func main() {
	p2:=Add2()
	fmt.Printf("Call Add2 for 3 gives: %v\n",p2(3))
	TwoAdder := Adder(2)
	fmt.Printf("The result is: %v\n",TwoAdder(3))
}
func Add2() func(b int) int{
	return func(b int) int {
		return b + 2
	}
}
func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}
*/
/*
package main

import "fmt"

func Adder() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}
func main() {
	var f = Adder()
	fmt.Println(f(300))
	fmt.Println(f(300))
	fmt.Println(f(300))
}
*/
/*
// 工厂函数
package main
import (
	"fmt"
	"strings"
)

func main() {
	addJpg :=MakeAddSuffix(".jpg")
	addBmp :=MakeAddSuffix(".bmp")
	fmt.Println(addBmp("df20190418"))
	fmt.Println(addJpg("xz1022193"))

}
func MakeAddSuffix(suffix string) func(string) string{
	return func(name string) string{
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
*/
/*
// 利用闭包实现斐波那契数列
package main

import (
	"fmt"
	"time"
)

func main() {
	f := fib()
	start := time.Now()
	for i := 0; i < 25; i++ {
		fmt.Println(f())
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)

}
func fib() func() int {
	var c,a,b int = 0,1,1
	return func() int{
		switch c {
		case 0,1:
			c ++
		default:
			a,b =b,b+a
		}
		return b
	}
}
*/
/*
// 计算函数执行时间
package main

import (
	"fmt"
	"time"
)

func main() {
	result := 0
	start := time.Now()
	for i := 0; i < 25; i++ {
		result = fib(i)
		fmt.Printf("fib(%d) is: %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}
func fib(n int) int {
	if n <= 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}
*/

// 通过内存缓存来提升性能
package main
import (
	"fmt"
	"time"
)
const LIM = 41
var fibs [LIM]uint64
func main() {
	var result uint64 = 0
	start := time.Now()
	for i := 0; i < LIM; i++ {
		result = fib(i)
		fmt.Printf("fib(%d) is: %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}
func fib(n int) (res uint64) {
	if fibs[n] != 0{
		res = fibs[n]
		return
	}
	if n<= 1{
		res = 1
	}else {
		res = fib(n-1) + fib(n-2)
	}
	fibs[n] = res
	return 
}