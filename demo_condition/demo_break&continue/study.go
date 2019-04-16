/*
// break
package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 10; j++ {
			if j > 5 {
				break
			}
			fmt.Printf("%d",j)
		}
		fmt.Println()
	}
}
*/

// continue
package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Printf("%d", i)
	}
	// fmt.Println()
}
