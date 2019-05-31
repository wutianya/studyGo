/*
package main

import (
	"errors"
	"fmt"
)

func main() {
	a := func() {fmt.Println("1")}
	fmt.Println(&a)
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}
*/
/*
package main

import "fmt"

var (
	num          int = 10
	numx2, numx3 int
)

func main() {
	numx2, numx3 = getX2AndX3(num)
	PrintValues()
	numx2, numx3 = getX2AndX3_2(num)
	PrintValues()
	fmt.Println(numx2,numx3)
}
func PrintValues() {
	fmt.Printf("num = %d, 2x num = %d, 3x num = %d\n", num, numx2, numx3)
}
func getX2AndX3(input int) (int, int) {
	return 2 * input, 3 * input
}
func getX2AndX3_2(input int) (x2, x3 int) {
	x2 = 2 * input
	x3 = 3 * input
	return
}
*/
/*
package main

import "fmt"

func main() {
	sum, sub, mul := multRet(10, 5)
	fmt.Println(sum, sub, mul)
}
func multRet(a, b int) (sum, sub, mul int) {
	return a + b, a - b, a * b
	// sum = a + b
	// sub = a - b
	// mul = a * b
	// return
}
*/
/*
// 计算一个 float64 类型浮点数的平方根，如果参数是一个负数的话将返回一个错误。编写两个版本，一个是非命名返回值，一个是命名返回值
package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	fmt.Print("First example with -1: ")
	ret1, err1 := MySqrt(-1)
	if err1 != nil {
		fmt.Println("Error! Return values are: ", ret1, err1)
	} else {
		fmt.Println("It's ok! Return values are: ", ret1, err1)
	}

	fmt.Print("Second example with 5: ")
	//you could also write it like this
	if ret2, err2 := MySqrt(5); err2 != nil {
		fmt.Println("Error! Return values are: ", ret2, err2)
	} else {
		fmt.Println("It's ok! Return values are: ", ret2, err2)
	}
	// named return variables:
	fmt.Println(MySqrt2(5))
}
func MySqrt(f float64) (float64, error) {
	if f < 0 {
		return float64(math.NaN()), errors.New("I won't be able to do a sqrt of negative number!")
	}
	return math.Sqrt(f), nil
}
func MySqrt2(f float64) (ret float64, err error) {
	if f < 0 {
		ret = float64(math.NaN())
		err = errors.New("I won't be able to do a sqrt of negative number!")
	}
	ret = math.Sqrt(f)
	return
}
*/
/*
// 改变外部变量,传递指针给函数
package main

import (
	"fmt"
)

func main() {
	n := 0
	reply := &n
	Multiply(10, 5, reply)
	fmt.Printf("Multiply: %d", *reply)
}
func Multiply(a, b int, reply *int) {
	*reply = a * b
}
*/

// 传递变长参数
package main

import (
	"fmt"
)

func main() {
	// x := min(1, 3, 2, 0)
	// fmt.Printf("The mininum is: %d\n", x)
	slice := []int{7,9,3,5,1}
	x := min(slice...)
	fmt.Printf("The mininum is: %d\n", x)
}
func min(s ...int) int {
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		fmt.Printf("v = %d min = %d\n", v, min)
		if v < min {
			fmt.Printf("v = %d min = %d\n", v, min)
			min = v
		}
	}
	return min
}

/*
package main

import (
	"fmt"
)

func main() {
	s := []string{"jreey", "balance", "stack"}
	foo(s...)
}
func foo(s ...string) {
	if len(s) == 0 {
		return
	}
	for _, v := range s {
		fmt.Println(v)
	}
}
*/
/*
package main
import (
	"fmt"
)
func main() {
	l :=[]int{1,2,4,3}
	L :=[...]int{1,2,4,3}
	bar(l)
	fmt.Printf("l type is: %T\n",l)
	fmt.Printf("l type is: %T",L)
}
func bar(s []int)  {
	fmt.Println(s)
}
*/