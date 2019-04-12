// package main
// import "fmt"

// func main() {
// 	var balance [5] float32
// 	var a [5][4] string 
//  	balance[0], balance[1] = 12, 11.21
// 	a[1][0] = "99"
// 	fmt.Println(balance[1])
// 	fmt.Println(a)
// }

// package main
// import "fmt"

// func main() {
// 	var balance = []int {2, 4, 6, 8, 10}
// 	var avg float32
// 	avg = getAverage(balance, 5)
// 	fmt.Printf( "平均值为: %f ", avg )
// }

// func getAverage(arr [] int, size int) float32 {
// 	var i int 
// 	var avg, sum float32
// 	for i = 0; i < size; i++ {
// 		sum += float32(arr[i])
// 	}
// 	avg = float32(sum) / float32(size)
// 	return avg
// }

// package main
// import "fmt" 

// func main() {
// 	var a int = 10
// 	fmt.Printf("var is address: %x\n", &a)
// }

// package main
// import (
// 	"fmt"
// 	"strconv"
// )

// const (
// 	a, b  = 1, 3.14
// 	str = len("hello")
// )
// func main() {
// 	x, _ := strconv.Atoi("12")
// 	fmt.Printf("%v %v\n",a, b)
// 	println(x, str)
// }

// package main 
// import (
// 	"errors"
// 	"log"
// )

// func check(x int) error {
// 	if x <= 0 {
// 		return errors.New("x <= 0")
// 	}
// 	return nil
// }
// func main() {
// 	x := -1
// 	if err := check(x); err == nil {
// 		x ++
// 		println (x)
// 	} else {
// 		log.Fatalln(err)
// 	}
// }
// package main
// import "fmt"

// func test(x *int) {
// 	fmt.Printf("pointer: %p, target: %v\n", &x, x)
// }

// func main() {
// 	a := 0x100
// 	p := &a
// 	fmt.Printf("pointer: %p, target: %v\n", &p, p)
// 	test(p)
// }

// package main 

// import "fmt"

// func test(s string, a ...int){
// 	fmt.Printf("%s, %T, %v\n",s, a, a[0])
// }

// func main() {
// 	test("str", 1,2,3,4)
// }

// package main
// import "fmt"

// func test(a ...int) {
// 	fmt.Println(a)
// }

// func main() {
// 	// var a = [3]int{10, 15, 25}
// 	a := []int{10, 20, 30}
// 	test(a[:]...)
// }
// package main
// import (
// 	"fmt"
// 	"errors"
// )

// func div(x, y int) (int, error) {
// 	if y == 0 {
// 		return 0, errors.New("division by zero")
// 	}
// 	return x / y, nil
// }
// func log(x int, err error) {
// 	fmt.Println(x, err)
// }
// func test() (int, error) {
// 	return div(5, 0)
// }
// func main() {
// 	log(test())
// 	fmt.Printf("%v\n",add(1,3))
// }
// func add(x, y int) (z int) {
// 	{
// 		z := x + y
// 		return z
// 	} 
// 	return z 
// }

// 匿名函数
// package main
// import "fmt"

// func test(f func()) {
// 	f()
// }

// func main() {
// 	test (func() {
// 		fmt.Println("hello,world")
// 	})
// }

// 闭包
// package main
// import "fmt"

// func test(x int) func() {
// 	// fmt.Println(x)
// 	// x = 100
// 	return func() {
// 		fmt.Println(x)
// 	}
// }
// func main() {
// 	f := test(123)
// 	f()
// }
// package main
// import "fmt"

// func test(x int) func() {
// 	fmt.Println(x)
// 	return func() {
// 		fmt.Println(&x, x)
// 	}
// }
// func main() {
// 	f := test(0x100)
// 	f()
// }

// 延迟调用
// package main
// import (
// 	"fmt"
// 	"math"
// )
// /*
// func main() {
// 	var x, y int = 1, 2
// 	defer func(a int) {
// 		fmt.Printf("defer x, y = %d, %d\n", a, y)
// 	}(y)
// 	x += 100
// 	y += 200
// 	fmt.Println(x, y)
// }
// */
// func test() (z int) {
// 	defer func() {
// 		fmt.Printf("defer:%d\n", z)
// 		z += 100
// 	}()
// 	return 100
// }
// func main() {
// 	fmt.Printf("test:%d", test())
// 	fmt.Println(math.MaxInt16)
// }

// package main
// import (
// 	"fmt"
// )
// func main() {
// 	var a int = 66
// 	b := string(a)
// 	fmt.Println(b)
// }

package main
import (
	"fmt"
	"strconv"
)
func main() {
	var a int = 65
	b := strconv.Itoa(a)
	a, _ = strconv.Atoi(b)
	fmt.Println(a)
}