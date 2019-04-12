// package main
// import "fmt"

// func main() {
// 	fmt.Println("Hello, 世界")
// }

// -----------

// package main
// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	var s, sep string
// 	// := 短变量声明
// 	for i := 1; i < len(os.Args); i++ {
// 		sep = " "
// 		s += sep + os.Args[i]
// 	}
// 	fmt.Println(s)
// }

// ---------------

// package main
// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	s, sep := "", ""
// 	for _, arg := range os.Args[1:] {
// 		s += sep + arg
// 		sep = " "
// 	}
// 	fmt.Println(s)
// }
// ---------------
// package main
// import "fmt"

// func main() {
// 	const LENGTH int = 10
// 	const WIDTH int = 5
// 	var area int 
// 	const a, b, c = 1, false, "str"

// 	area = LENGTH * WIDTH
// 	fmt.Printf("面积为: %d\n", area)
// 	// fmt.Println()
// 	fmt.Println(a, b, c)
// }
// ---------------
// package main

// import "fmt"
// const (
//     i=1<<iota
// 	// j=3<<iota
// 	j=iota
//     k
//     l
// )

// func main() {
//     fmt.Println("i=",i)
//     fmt.Println("j=",j)
//     fmt.Println("k=",k)
//     fmt.Println("l=",l)
// }
// ---------------
package main

import "fmt"

func main() {
	var a uint = 3
	var c uint = 0
	c = a << 1
	fmt.Println(a, c)
}