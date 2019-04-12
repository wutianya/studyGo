/*
package main
import "fmt"
func main() {
	a := f1()
	fmt.Println(a)
}
func f1() string {
	if condition := true ; condition {
		return "123"
	}else {
		return "435"
	}
}
*/

// 测试多返回值函数的错误
package main
import(
	"fmt"
	"strconv"
)
func main() {
	var orig string = "abc"
	var newS string
	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
	an, err := strconv.Atoi(orig)
	if err != nil {
		fmt.Printf("orig %s is not integger - exiting with error \n", orig)

		return 
	}
	fmt.Printf("The integer is %d\n", an)
	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)
}