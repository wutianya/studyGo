/*
// 读取请求首部
package main

import (
	"fmt"
	"net/http"
)

// use handlefunc
func hello(w http.ResponseWriter, r *http.Request) {
	h := r.Header //type is map
	fmt.Fprintln(w, "Header: ", h)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/",hello)
	server.ListenAndServe()
}
*/
/*
// 读取请求主体中的数据
package main
import (
	"fmt"
	"net/http"
)

func body(w http.ResponseWriter, r *http.Request){
	len := r.ContentLength
	body := make([]byte,len)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))
}
func main(){
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}
	http.HandleFunc("/body",body)
	server.ListenAndServe()
}
*/

// 对表单进行语法分析
package main

import (
	"fmt"
	"net/http"
)
func process(w http.ResponseWriter, r *http.Request) {
	// enctype = application/x-www-form-urlencoded
	r.ParseMultipartForm(1024)
	fmt.Fprintln(w, r.Form)
	fmt.Fprintln(w, r.FormValue("hello"))
	fmt.Fprintln(w, r.PostFormValue("hello"))
	fmt.Fprintln(w, r.PostForm)
	fmt.Fprintln(w, r.MultipartForm)
	// enctype = multipart/form-data
	// r.ParseMultipartForm(1024)
	// fmt.Fprintln(w, r.MultipartForm)
}
func main(){
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}