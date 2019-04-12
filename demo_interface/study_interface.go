package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("cpu number is: ", runtime.NumCPU())
}