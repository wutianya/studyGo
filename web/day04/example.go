/*
	模板操作
		1. 条件动作
		2. 迭代动作
		3. 设置动作
		4. 包含动作
*/
/*
// 在处理器函数中触发模板引擎
package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	// t, _ := template.ParseFiles("tmpl.html")
	//		|
	//		|
	//		V
	// t := template.New("tmpl.html")
	// t, _ = t.ParseFiles("tmpl.html")
	t := template.Must(template.ParseFiles("tmpl.html"))
	t.Execute(w, "hello world")
}

func main() {
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
*/

// 