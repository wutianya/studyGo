// 数组的定义
// var arr [5]int --> var arr = [int]5{} --> arr := [5]int{}
// var arr = new([5]int)

package main

import "fmt"

func main() {
	arr := [5]int{1, 5, 3, 4}
	var arr1 = new([5]int)
	a := [...]string{"a", "b", "c", "d"}
	fmt.Println(arr1)
	fmt.Println(len(arr), cap(a))
	for i, _ := range arr {
		fmt.Println(arr[i])
	}
	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}
}

/*
package main

import "fmt"

func f(a [3]int) {
	fmt.Println(a)
}
func fp(a *[3]int) {
	fmt.Println(a)
}
func main() {
	var ar [3]int
	f(ar)
	fp(&ar)
}
*/
/*
package main
import "fmt"

func main() {
	var a = [5]int{2,4,1,6,9}
	b := &a
	b[2] = 8
	fmt.Println(a)
	fmt.Println(b) 
}
*/
/*
package main
import "fmt"

func main() {
	// var arrAge = [5]int{18, 20, 15, 22, 16}
	// var arrLazy = [...]int{5, 6, 7, 8, 22}
	var arrLazy = []int{5, 6, 7, 8, 22}
	var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}
	// var arrKeyValue = []string{3: "Chris", 4: "Ron"}
	fmt.Println(append(arrLazy,88))
	for i:=0; i < len(arrKeyValue); i++ {
		fmt.Printf("Person at %d is %s\n", i, arrKeyValue[i])
	}
}
*/

/*
把一个大数组传递给函数会消耗很多内存。有两种方法可以避免这种现象：

传递数组的指针
使用数组的切片
*/
/*
package main
import "fmt"
func main() {
	arr := [3]float64{1.1,2,3}
	x := Sum(&arr)
	fmt.Println(x)
}
func Sum(a *[3]float64) (sum float64){
	for _,v :=range a {
		sum += v
	}
	return 
}
*/