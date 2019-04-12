/*
实现计算机单位的枚举
*/
/*
package main

import(
	"fmt"
)

const (
	// iota 是枚举
	B float64 = 1 << (iota *10)
	KB
	MB
)

func main() {
	fmt.Println(B)
	fmt.Println(KB)
	fmt.Printf("MB:%v", MB)
}
*/

/*
package main
import "fmt"

func main() {
	a := 1024
	var p *int = &a
	fmt.Println(*p)
	fmt.Println(p)
}
*/

/*
// switch
package main

import "fmt"

func main() {
	switch a := 1; {
	case a >= 0:
		fmt.Println("a>=0")
		fallthrough //继续检索后面的条件
	case a >= 1:
		fmt.Println("a>=1")
	default:
		fmt.Println("None")
	}
}
*/
/*
package main

import (
	"fmt"
)

func main() {
LABEL:
	for i := 0; i < 3; i++ {
		for {
			fmt.Println(i)
			break LABEL
		}
	}
	fmt.Println("ok")
}
*/
/*
package main

import (
	"fmt"
)

func main() {
	// 	a := [...]int{5:6}
	// 	var count = len(a)
	// 	for i:=0;i<count;i++{
	// 		fmt.Println(a[i])
	// 	}
	// }
	p := new([5]int)
	fmt.Println(p)
}
*/
/*
package main
import "fmt"
func main() {
	a := [2][3] int {{1,2,3},{2,3,4}}
	var b [2][2] int

	fmt.Println(a[0][0],b)

}
*/

package main

import "fmt"

func main() {
	// 冒泡排序
	a := [...]int{10, 4, 7, 8, 32, 9}
	count := len(a)
	for i := 0; i < count; i++ {
		for j := i + 1; j < count; j++ {
			fmt.Printf("this is a[i]: %v\n", a[i])
			fmt.Printf("this is a[j]: %v\n", a[j])
			if a[i] < a[j] {
				fmt.Printf("a[i] is %v, a[j] is %v\n", a[i], a[j])
				a[i], a[j] = a[j], a[i]
			}
			// fmt.Println(a)
		}
	}
	fmt.Println(a)
}

/*
package main
import(
	"fmt"
)

// slice
func main() {
	// s1 := make([]int,3,10)
	// fmt.Println(len(s1), cap(s1))
	// fmt.Println(s1)
	a := []int{1, 2, 3, 4, 5}
	s1 := a[2:5]
	s2 := a[1:3]
	s1[0] = 8
	fmt.Println(s1,s2)
	fmt.Printf("%p,%p",a,s1)
}
*/
/*
package main
import "fmt"

// slice copy append
func main() {
	// define slice
	s1 := []int{1,2,4}
	s2 := []int{4,5,6}
	// append
	s1 = append(s1,6)
	s1 = append(s1,s2...)
	s3 := s1[:]
	// copy
	copy(s2,s1[5:6])
	fmt.Println(s1[5:6])
	fmt.Println(s1,s2)
	fmt.Printf("%p,%p",s1,s3)
}
*/
/*
package main
import "fmt"
// map
func main() {
	// define map
	// var m map[int]string
	// var m map[int]string = make(map[int]string)
	m := make(map[int]string)
	m[1] = "OK"
	// delete key
	delete(m, 1)
	fmt.Println(m)
}
*/
/*
package main

import "fmt"
// map ex
func main() {
	m1 := map[int]string{1: "a", 2: "b", 3: "c"}
	// m2 := map[string]int{}
	m2 := make(map[string]int)
	for k, v := range m1 {
		m2[v] = k
	}
	fmt.Printf("m1: %v, m2: %v\n", m1, m2)
	fmt.Println(m2["a"],m2["b"],m2["c"])
}
*/

/*
package main
import "fmt"

func main(){
	test()
}
func test(){
	fmt.Println("I'm function test")
}
*/
/*
package main

import "fmt"

func main() {
	var fs = [4]func(){}

	for i := 0; i < 4; i++ {
		defer fmt.Println("defer i = ", i)
		defer func() {
			fmt.Println("defer_closure i = ", i)
		}()
		fs[i] = func() {
			fmt.Println("closure i = ", i)
		}
	}
	for _, f := range fs {
		f()
	}
}
*/
