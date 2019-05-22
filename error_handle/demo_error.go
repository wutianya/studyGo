/*
package main

import (
	"errors"
	"fmt"
)
var errNotFound error = errors.New("Not found error")

func main() {
	fmt.Printf("error: %v", errNotFound)
}
*/
/*
// panic
package main

import ("fmt";"os";"log")

func main() {
	fmt.Println("starting teh program")
	fmt.Println(os.Getenv("USER"))
	defer fmt.Println("ending the program")
	// panic("a server error occurred: stopping the program")
	protect(test)
	
}
func test(){
	fmt.Println("test")
}
func protect(g func()) {
	defer func() {
		log.Println("done")
		if err := recover(); err != nil {
			log.Printf("run time panic: %v",err)
		}
	}()
	log.Println("start")
	g()
}
*/
/*
package main

import (
	"fmt"
)

func badcall() {
	panic("bad end")
}

func test() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicing %s\r\n", e)
		}
	}()
	badcall()
	fmt.Printf("After bad call\r\n")
}

func main() {
	fmt.Printf("Calling test\r\n")
	test()
	fmt.Printf("Test completed\r\n")
}
*/

// exec.go
package main

import (
	"fmt"
	// "os/exec"
	"os"
)

func main() {
	env := os.Environ()
	procAttr := &os.ProcAttr{
		Env: env,
		Files: []*os.File{
			os.Stdin,
			os.Stdout,
			os.Stderr,
		},
	}
	// go fmt.Println("hello world")
	pid, err := os.StartProcess("/bin/ls", []string{"ls", "-l"}, procAttr)
	if err != nil {
		fmt.Printf("Error %v starting process!", err)
		os.Exit(1)
	}
	fmt.Printf("The process id is %v", pid)
}