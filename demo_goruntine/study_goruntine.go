/*
// goruntine
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("In main()")
	// longWait()
	// shortWait()
	go longWait()
	go shortWait()
	fmt.Println("About to sleep in main()")
	time.Sleep(10 * 1e9)
	fmt.Println("At the end of main()")
}

func longWait() {
	fmt.Println("Beginning longWait()")
	time.Sleep(5 * 1e9)
	fmt.Println("End of LongWait()")
}

func shortWait() {
	fmt.Println("Beginning shortWait()")
	time.Sleep(2 * 1e9)
	fmt.Println("End of shortWait()")
}
*/

// channel
/*
define channel:
	var identifier chan datatype
	// create channel
	var ch1 chan string
	// instance channel
	ch1 = make(chan string)

	通信操作符 <-
	ch <- int1 表示：用通道 ch 发送变量 int1（双目运算符，中缀 = 发送
	int2 = <- ch 表示：变量 int2 从通道 ch（一元运算的前缀操作符，前缀 = 接收）接收数据（获取新值）；假设 int2 已经声明过了，如果没有的话可以写成：int2 := <- c
*/
/*
package main

import (
	"fmt"
	"time"
)

func main() {
	// var ch chan int
	ch := make(chan string)
	go sendData(ch)
	go getData(ch)

	time.Sleep(1e9)
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "hello"
	ch <- "world"
}

func getData(ch chan string) {
	var input string
	// time.Sleep(1e9)
	for {
		input = <- ch
		fmt.Println("input: ", input)
	}
}
*/
/*
// 通道阻塞
package main

import ("fmt";"time")

func main() {
	ch := make(chan int)
	go pump(ch)
	go suck(ch)
	time.Sleep(1e9)
}

func pump(ch chan int) {
	for i := 1; ; i++ {
		ch <- i
	}
}

func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}
*/
/*
package main
import ("fmt"; "time")

func main() {
	c := make(chan int)

	go func() {
		time.Sleep(5 * 1e9)
		// variable x received channel c
		// x := <- c
		fmt.Println("received: ", <-c)
	}()

	fmt.Println("sending: ", 10)
	// channel sent variable 10
	c <- 10
	fmt.Println("sent",10)
}
*/
/*
// channel deadlock
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	// go f1(ch)
	// channel sent
	go func() {
		ch <- 2
	}()
	go func() {
		fmt.Println(<-ch)
	}()

	time.Sleep(0.1 * 1e9)
}
func f1(ch chan int) {
	// time.Sleep(2 * 1e9)
	fmt.Println(<-ch)
}
*/
/*
// 使用带缓冲的通道
package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 2)
	go func() {
		time.Sleep(5 * 1e9)
		fmt.Println("received: ", <-c)
	}()
	fmt.Println("sending: ", 10)
	c <- 10
	fmt.Println("sent: ", 10)
	
}
*/

// go sum.go

package main
import "fmt"

func sum(x, y int, c chan int) {
	c <- x + y
}

func main() {
	c := make(chan int)
	go sum(12,6,c)
	fmt.Println(<-c)
}