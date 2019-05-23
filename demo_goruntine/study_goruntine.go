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
	time.Sleep(6 * 1e9)
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
/*
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
*/
/*
// goroutines2.go
package main

import (
	"fmt"
)

func numGen(start, count int, out chan<- int) {
	for i := 0; i < count; i++ {
		// use channel sent variable start for type int
		out <- start
		start += count
	}
	close(out)
}
func numEchoRangein(in <-chan int, done chan<- bool) {
	for num := range in {
		fmt.Printf("%d\n", num)
	}
	done <- true
}
func main() {
	numChan := make(chan int)
	done := make(chan bool)
	go numGen(0, 10, numChan)
	go numEchoRangein(numChan, done)

	<-done
}
*/
/*
package main

import (
	"fmt"
	"time"
)

func main() {
	stream := pump()
	go suck(stream)
	time.Sleep(1e9)
}
func pump() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}
func suck(ch chan int) {
	// for {
	// 	fmt.Println(<-ch)
	// }
	for v := range ch {
		fmt.Println(v)
	}
}
*/
/*
package main

import "fmt"

func generate(ch chan int) {
	for i := 2; ; i++ {
		ch <- i // send i to channel ch
	}
}

func filter(in, out chan int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

func main() {
	ch := make(chan int)
	go generate(ch)
	for c := 0; c < 3; c++ {
		prime := <-ch
		fmt.Println("prime: ", prime)
		ch1 := make(chan int)
		a := <-ch
		fmt.Println("ch: ", a)
		go filter(ch, ch1, prime) // 2 3 2
		b := <-ch1
		fmt.Println("ch1: ", b)
		ch = ch1 // ch = 3
	}
}
*/
/*
// 显式关闭通道
package main

import "fmt"
// import "time"

func main() {
	ch := make(chan string)
	go sendData(ch) // sendData 是协程,和main函数不在同个线程中
	getData(ch) // getData 是和main函数在同个线程中, 只有使用了go关键字的才是协程
	// time.Sleep(1e9)
}

func sendData(ch chan string) {
	ch <- "Go"
	ch <- "Pyhon"
	ch <- "Shell"
	ch <- "Java"
	close(ch) // 通道仅为发送者才能使用显示关闭通道
}

func getData(ch chan string) {
	for {
		input, open := <-ch // the "open" type is bool
		if !open {
			break
		}
		fmt.Printf("input: %s\n", input)
	}
}
*/

/*
// 使用 select 切换协程
// goroutine_select.go

package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go pump1(ch1)
	go pump2(ch2)
	go suck(ch1, ch2)
	time.Sleep(1e9)
}
func pump1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 2
	}
}
func pump2(ch chan int) {
	for i := 0; ; i++ {
		ch <- i + 5
	}
}
func suck(ch1, ch2 chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Printf("Received on channel 1: %d\n", v)
		case v := <-ch2:
			fmt.Printf("Received on channel 2: %d\n", v)
		}
	}
}
*/
// /*
package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	var i int
	var ok bool = true
	go tel(ch)

	for ok {
		if i, ok = <-ch; ok {
			fmt.Println(i)
		}
	}

}

func tel(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch) // if this is ommitted: panic: all goroutines are asleep - deadlock!
}

// */
