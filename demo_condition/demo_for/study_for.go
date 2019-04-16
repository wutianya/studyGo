/*
package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}
*/
/*
package main

import "fmt"

func main() {
	// for i := 0; i < 15; i++ {
	// 	fmt.Printf("The counter is at %d\n", i)
	// }
	i := 0
START:
	fmt.Printf("The counter is at %d\n", i)
	i++
	if i < 15 {
		goto START
	}
}
*/
/*
package main

import (
	// "fmt"
	// "strings"
)

func main() {
	// s := "G"
	// for i := 1; i <= 25; i++ {
	// 	fmt.Println(strings.Repeat(s, i))
	// }
	for i := 1; i <= 25; i++ {
		for j := 1; j <= i; j++ {
			print("G")
		}
		println()
	}
}
*/
/*
//从 1 打印到 100 的程序，但是每当遇到 3 的倍数时，不打印相应的数字，但打印一次 "Fizz"。遇到 5 的倍数时，打印 Buzz 而不是相应的数字。对于同时为 3 和 5 的倍数的数，打印 FizzBuzz
package main

import "fmt"

const (
	FIZZ     int = 3
	BUZZ     int = 5
	FIZZBUZZ int = 15
)

func main() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%FIZZBUZZ == 0:
			fmt.Println("FIZZBUZZ")
		case i%FIZZ == 0:
			fmt.Println("Fizz")
		case i%BUZZ == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}
*/

package main

import "fmt"

func main() {
	for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j,
		s = i+1, j+1, s+"a" {
		fmt.Println("Value of i, j, s:", i, j, s)
	}
}
