/*
// 使用等待组
package main

import (
	"fmt"
	"time"
	"sync"
)

func printNumbers2(wg *sync.WaitGroup) {
	for i := 0; i < 10 ; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%d ", i)
	}
	wg.Done() // 对计数器执行减一操作
}
func printLetters2(wg *sync.WaitGroup) {
	for i := 'A'; i < 'A'+10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%c ", i)
	}
	wg.Done()
}
func main() {
	var wg sync.WaitGroup // 声明一个等待组
	wg.Add(2)	// 为计数器设置值
	go printNumbers2(&wg)
	go printLetters2(&wg)
	wg.Wait() // 阻塞到计数器的值为0
}
*/
/*
// 使用通道同步goroutine
package main

import (
	"fmt"
	"time"
)

func printNumbers2(w chan bool) {
	for i := 0 ; i < 10; i ++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%d ", i)
	}
	// 把一个布尔值放入通道,以便解除主程序的阻塞状态
	w <- true
}
func printLetters2(w chan bool) {
	for i := 'A'; i < 'A'+10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%c ", i)
	}
	w <- true
}
func main() {
	w1, w2 := make(chan bool), make(chan bool)
	go printNumbers2(w1)
	go printLetters2(w2)
	// 主程序将一直阻塞,直到通道里面出现可弹出的值为止
	<-w2
	<-w1	
}
*/
/*
// 使用通道实现消息传递
package main
import (
	"fmt"
	"time"
)

func thrower(c chan int) {
	for i := 0; i < 5; i++ {
		c <- i // 把值传入channel c中
		fmt.Println("Threw >> ", i)
	}
}
func catcher(c chan int) {
	for i := 0; i < 5; i++ {
		num := <-c // 将channel中的值取出赋值给num
		fmt.Println("Caught << ", num)
	}
}
func main() {
	c := make(chan int)  // 无缓冲通道
	// c := make(chan int, 3) //有缓冲通道
	go thrower(c)
	go catcher(c)
	time.Sleep(200 * time.Millisecond)
}
*/
/*
// 从多个通道中选择
package main

import (
	"fmt"
	"time"
)

func callerA(c chan string) {
	c <- "hello world"
}
func callerB(c chan string) {
	c <- "hello mundo"
}

func main() {
	a ,b := make(chan string), make(chan string)
	go callerA(a)
	go callerB(b)
	for i := 0 ; i < 5 ; i++ {
		time.Sleep(1 * time.Microsecond) //使通道能够正常接收goroutine发生给它们的值
		select {
		case msg := <- a:
			fmt.Printf("%s from A\n", msg)
		case msg := <- b:
			fmt.Printf("%s from B\n", msg)
		default:
			fmt.Printf("Default\n")
		}
	}
}
*/

// close channel
package main

import (
	"fmt"
)

func callerA(c chan string) {
	c <- "hello world"
	close(c)
}
func callerB(c chan string) {
	c <- "hola mundo"
	close(c)
}
func main(){
	a, b := make(chan string), make(chan string)
	ok1, ok2 := true, true
	go callerA(a)
	go callerB(b)
	var msg string
	for ok1 || ok2 {
		select {
		// 多值格式, 从通道a里面取出的值被赋值给变量msg,
		// 变量ok1则会被设置为用于表示通道是否仍然处于打开状态的布尔值. 如果通道已被关闭,那么ok1的值将被设置为false
		case msg, ok1 = <-a:
			if ok1 {
				fmt.Printf("%s from A\n", msg)
			}
			fmt.Println("ok1: ", ok1)
		case msg, ok2 = <-b:
			if ok2 {
				fmt.Printf("%s from B\n", msg)
			}
			fmt.Println("ok2: ", ok2)
		}
	}
}