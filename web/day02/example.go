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
