// package main

// import "fmt"
// func main(){
//     // var count,c int   //定义变量不使用也会报错
//     var count int
//     var flag bool
//     count=1
//     //while(count<100) {    //go没有while
//     for count < 100 {
//         count++
//         flag = true;
//         //注意tmp变量  :=
//         for tmp:=2;tmp<count;tmp++ {
//             if count%tmp==0{
//                 flag = false
//             }
//         }

//         // 每一个 if else 都需要加入括号 同时 else 位置不能在新一行
//         if flag == true {
//             fmt.Println(count,"素数")
//         }else{
//             continue
//         }
//     }
// }

// package main

// import "fmt"

// func main() {
// 	var grade string = "B"
// 	var marks int = 41
// 	switch marks {
// 	case 90: grade = "A"
// 	case 80: grade = "B"
// 	case 50,60,70: grade = "C"
// 	default: grade = "F"
// 	}

// 	switch {
// 		case grade == "A":
// 			fmt.Println("优秀!")
// 		case grade == "B", grade == "C":
// 			fmt.Println("lianghao")
// 		case grade == "D":
// 			fmt.Println("jige")
// 		case grade == "F":
// 			fmt.Println("bujige")
// 		default:
// 			fmt.Println("cha")
// 	}
// 	fmt.Printf("ni de deng ji shi %s\n", grade)
// }

// package main

// import "fmt"

// func main() {
// 	var a int = 10
// 	test: for a < 20 {
// 		if a == 15 {
// 			a += 1
// 			goto test
// 		}
	
// 		fmt.Println(a)
// 		a ++
// 	}
// }

// package main

// import "fmt"

// func max(num1, num2 int) int { 
// 	if (num1 > num2) {
// 		return num1
// 	}else {
// 		return num2
// 	}
// }
// func main() {
// 	var a, b int = 100, 200
// 	var ret int
// 	ret = max(a, b)
// 	fmt.Printf("max numbers is: %d\n", ret)
// }

// package main

// import "fmt"

// func swap(x, y string) (string, string) {
// 	return y, x
// }

// func main() {
// 	a, b := swap("hello", "12")
// 	fmt.Println(a, b)
// }

package main
import "fmt"

func main() {
	var a, b int = 100, 200
	swap(&a, &b)
	fmt.Printf("交换后，a 的值 : %d\n", a )
	fmt.Printf("交换后，b 的值 : %d\n", b )
}

func swap(x *int, y *int) {
	var tmp int 
	tmp = *x
	*x = *y
	*y = tmp
}

