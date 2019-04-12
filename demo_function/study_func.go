package main

import "fmt"

func main() {
	a := func() {fmt.Println("1")}
	fmt.Println(&a)
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i)
		}()
	}
}
