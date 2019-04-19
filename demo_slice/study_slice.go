//  slice
/*
切片是一个 长度可变的数组
切片是 引用
make 创建切片
	slice1 := make([]type, len)
	s2 := make([]int, 10)
	var slice1 []int = make([]int, 10)
*/
/*
package main
import (
	"fmt"
)
func main() {
	var arr = [5]int{18,10,32,91,99}
	s := arr[:]
	m := Max(s)
	fmt.Println(m)
}
// 获取slice中的最大值
func Max(arr []int) (ret int){
	ret = arr[0]
	for v,_ := range arr{
		if arr[v] > ret {
			ret = arr[v]
		}
	}
	return
}
*/
/*
package main

import "fmt"

func main() {
	var arr1 [6]int
	slice1 := arr1[1:5]
	fmt.Println(cap(arr1))
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}
	fmt.Printf("arr1: %v\nslice1: %v\n", arr1, slice1)
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	fmt.Printf("The length of arr1 is %d\n", len(arr1))
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))
}
*/
/*
package main
import "fmt"

func main() {
	b:= []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	fmt.Println(b[1:4])
	fmt.Println(b[:2])
	fmt.Println(b[2:])
	fmt.Println(b[:])
}
*/
/*
// For-range 结构
package main
import "fmt"

func main() {
	s :=[]int{1,5,3,4}
	for ix,v:=range s{
		fmt.Printf("ix is: %d, v is: %d\n", ix, v)
	}
}
*/
/*
package main
import "fmt"
func main() {
	items := [...]int{10, 20, 30, 40, 50}
	// for _,item := range items {
	// 	item *= 2
	// }
	for i := range items {
		items[i] *= 2
	}
	fmt.Println(items)
}
*/
/*
package main
import "fmt"

func main() {
	var ar = [10]int{0,1,2,3,4,5,6,7,8,9}
	var a = ar[5:7]
	fmt.Println(a)
	a = a[0:4]
	a[0] = 99

	fmt.Println(a[:])
}
*/
/*
// 切片的复制与追加
package main

import "fmt"

func main() {
	sl_from := []int{1, 2, 3}
	sl_to := make([]int, 10)
	n := copy(sl_to, sl_from)
	fmt.Printf("copied %d elements\n", n)
	fmt.Println(sl_to)

	arr := []int{99,98,97}
	sl3 := []int{3, 2, 1}
	sl3 = append(sl3, 1, 2, 3)
	sl3 = append(sl3,arr...)
	fmt.Println(sl3,len(sl3),cap(sl3))
}
*/
/*
package main
import "fmt"

func main() {
	s:=[]byte{11,24,5,2}
	fmt.Println(AppendByte(s,98,24))
}
func AppendByte(slice []byte, data ...byte) [] byte {

	m := len(slice)
	n := m + len(data)
	if n > cap(slice){
		newSlice := make([]byte,(n+1)*2)
		copy(newSlice,slice)
		slice = newSlice
		fmt.Println(slice)
	}
	slice = slice[0:n]
	fmt.Println(slice)
	copy(slice[m:n],data)
	return slice
}
*/
package main

import "fmt"

func main() {
	s := "hello"
	c := []byte(s)
	c[0] = 'c'
	s2 := string(c)
	fmt.Println(s2)
}
