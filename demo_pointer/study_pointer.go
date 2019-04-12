/*
package main
import (
	"fmt"
)

func main() {
	var i1 = 5
	i := i1
	// & Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址
	fmt.Printf("An integer : %d, it's location im memory: %p\n", i1, &i1)
	fmt.Println(&i)
	fmt.Println( &i1 == &i)
}
*/
/*
package main
import (
	"fmt"
)
func main() {
	var p *int
	a := 5
	p = &a
	fmt.Println(&a,p)
	fmt.Printf("a = %d, p = %d\n", a, *p)
}
*/

package main

import (
	"fmt"
)

func main() {
	s := "Good bye"
	var p *string = &s
	*p = "ciao"
	fmt.Printf("here is the pointer p: %p\n", p)
	fmt.Printf("here is the string *p: %s\n", *p)
	fmt.Printf("here is the string s: %s\n", s)
}
