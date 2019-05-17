/*
// strings.Join
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := [...]string{"foo", "bar", "baz"}
	s1 := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s1, ", "))
	fmt.Printf("s is %v\t s1 is %v", s, s1)
}
*/
/*
// 判断字符串字首和字尾
package main

import (
	"fmt"
	"strings"
)

// func main() {
// 	var s string="saas"
// 	fmt.Println(strings.HasPrefix(s, "s"))
// 	fmt.Println(strings.HasSuffix(s, "s"))
// }
func main() {
	var str string = "This is an example of a string"
	fmt.Printf("T/F? Does the string \"%s\" have prefix %s? ", str, "Th")
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
}
*/
/*
// 字符串包含
package main
import (
	"fmt"
	"strings"
)
func main() {
	if s := "like use cacht"; strings.Contains(s, "catch") {
		fmt.Println("condition is true")
	}else {
		fmt.Println("condition is false")
	}
}
*/
/*
//  判断子字符串或字符在父字符串中出现的位置（索引）
package main
import (
	"fmt"
	"strings"
)
func main() {
	var str string = "Hi, I'm Marc, Hi."
	fmt.Printf("The position of \"Marc\" is : %d\n",strings.Index(str, "Marc"))
	fmt.Printf("The position of the first instance of \"Hi\" is: %d\n", strings.Index(str, "Hi"))
	fmt.Printf("The position of the first instance of \"Hi\" is: %d\n", strings.LastIndex(str, "Hi"))
	fmt.Printf("The position of the first instance of \"Burger\" is: %d\n", strings.Index(str, "Burger"))
}
*/
/*
// 字符串替换&统计
package main
import (
	"fmt"
	"strings"
)
func main() {
	var s string = "l use python. python is simple"
	fmt.Println(strings.Replace(s, "python", "golang", 2))
	fmt.Printf("\"python\" total count is: %d\n", strings.Count(s, "python"))
}
*/
/*
// 重复字符串
package main
import (
	"fmt"
	"strings"
)
func main() {
	var s string = "Hi there! "
	var newS string
	newS = strings.Repeat(s, 3)
	fmt.Println(newS)
}
*/
/*
// 修改字符串大小写&修剪
package main
import (
	"fmt"
	"strings"
)
func main() {
	s := " Love Wife 520"
	fmt.Println(strings.ToLower(s))
	fmt.Println(strings.ToUpper(s))
	// trim 剔除字符串开头和结尾的空白符号
	fmt.Println(strings.TrimSpace(s))
	fmt.Println(strings.TrimRight(s, " 520"))
}
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "the quick brown foc jumps over the lazy dog\n 112"
	sl := strings.Fields(str)
	fmt.Println(sl)
	for _, val := range sl {
		fmt.Println(val)
	}
	str2 := "GO1|The ABC of Go|25"
	sl2 := strings.Split(str2, "|")
	fmt.Println("sl2: ", sl2)
	// join
	str3 := strings.Join(sl2, "---")
	fmt.Println(str3)
}

/*
// strconv
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var orig string = "666"
	var an int
	var newS string 
	var err error

	fmt.Printf("The size of ints is : %d\n", strconv.IntSize)
	// string ---> int 
	an, err = strconv.Atoi(orig)
	fmt.Println("err: ", err)
	fmt.Printf("The integer is: %d\n", an)
	an += 5
	// int ---> string
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)
}
*/