/*
	处理器:
		1. 定义struct
		2. 创建ServeHTTP方法: {w http.ResponseWriter, r *http.Request}
		3. 实例化struct
		4. 调用实例化对象: http.Handle("/hello", &hello) // hello 是实例化struct的对象
	处理器函数:
		1. 定义个拥有HandlerFunc函数类型的函数(interface: ResponseWriter, sturct: Request)
		2. 调用: http.HandlerFunc("/hello", hello) // 这里的hello是函数
	串联多个处理器:
		@处理器
		1. 创建函数
			func log(h http.Handler) http.Handler {
				return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					......
					h(w, r)
				})
			}
	串联多个处理器函数:
		@处理器函数
		1. 创建函数
			func log(h http.HandlerFunc) http.HandlerFunc {
				return func(w http.ResponseWriter, r *http.Request) {
					......
					h(w, r)
				}
			}
	ServeMux(HTTP请求多路复用器):
		负责接收http请求并根据请求中的url将请求重定向到正确的处理器.
		是一个结构类型,也是处理器
	DefaultServeMux:
		是ServeMux的一个实例
*/

/*
package main

import (
	"fmt"
	"net/http"
)

type MyHeadler struct {

}
func (h *MyHeadler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world~")
}

func main() {
	// http.ListenAndServe("", nil)

	handler := MyHeadler{}
	server := http.Server{
		Addr: "127.0.0.1:8888",
		Handler: &handler,
	}
	// server.ListenAndServeTLS("./simpletour.com.pem","simpletour.com.key")
	server.ListenAndServe()
}
*/
/*
// 使用多个处理器
package main

import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}
type WorldHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}
func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World")
}

func main() {
	hello := HelloHandler{}
	world := WorldHandler{}
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.Handle("/hello", &hello)
	http.Handle("/world", &world)
	server.ListenAndServe()
}
*/
/*
// 处理器函数

package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "hello")
}

func world(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "world")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", hello)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)

	server.ListenAndServe()
}
*/
/*
// 串联多个处理器函数
package main

import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

func hello(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "hello")
}

// this log function
func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called - " + name)
		h(w,r)
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/hello", log(hello))
	server.ListenAndServe()
}
*/
/*
// 串联多个处理器
package main
import (
	"fmt"
	"net/http"
)

type HelloHandler struct{}
func (h HelloHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello")
}

func log(h http.Handler) http.Handler {
	return http.HandlerFunc (func(w http.ResponseWriter, r *http.Request){
		fmt.Printf("Handler called - %T\n", h)
		h.ServeHTTP(w, r)
	}) 
}

func protect(h http.Handler) http.Handler {
	return http.HandlerFunc (func(w http.ResponseWriter, r *http.Request){
		fmt.Printf("Protect called - %T\n", h)
		h.ServeHTTP(w, r)
	})
}
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	hello := HelloHandler{}
	http.Handle("/hello",protect(log(hello)))
	server.ListenAndServe()

}
*/

// use httprouter
package main
import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", p.ByName("name"))
}

func main() {
	mux := httprouter.New()
	mux.GET("/hello/:name", hello)

	server := http.Server{
		Addr: "127.0.0.1:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}