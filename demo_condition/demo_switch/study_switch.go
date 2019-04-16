/*
package main
import (
	"cmd/go/internal/str"
	"fmt"
)

func main() {
	var num int = 99
	switch num {
	case 98,99:
		fmt.Println("It's equal to 98")
	case 100:
		fmt.Println("It's equal to 100")
	default:
		fmt.Println("It's not equal to 98 or 100")
	}
}
*/
/*
package main 
import (
	"fmt"
)
func main() {
	var num int = 7
	switch  {
	case num <0: 

		fmt.Println("number is negative")
	case num > 0:
		fmt.Println("number ge 0")
		fallthrough
	case num > 5:
		fmt.Println("number ge 5")
	default:
		fmt.Println("number is ")
	}
}
*/
/*
package main
import (
	"fmt"
)
func main() {
	switch num := 5; {
	case num < 0:
		fmt.Println("num lt 5")
	case num == 5:
		fmt.Println("num == 5")
	default:
		fmt.Println("num unkwon")
	}
}
*/
/*
package main
import (
	"fmt"
)
func main() {
	k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}
}
*/
package main
import (
	"fmt"
)
func main() {
	s := Season(5)
	fmt.Println(s)

}
func Season(month int) string {
	switch month {
	case 1,2,3:
		return "Spring"
	case 4,5,6:
		return "Summer"
	case 7,8,9:
		return "Autumn"
	case 10,11,12:
		return "Winter"
	default:
		return "Unknown month"
	}
}