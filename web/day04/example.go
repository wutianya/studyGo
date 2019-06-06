/*
	模板操作
		1. 条件动作 if
		2. 迭代动作	range
		3. 设置动作 with
		4. 包含动作
		5. 定义动作 define
	参数 变量和管道
*/
/*
// 在处理器函数中触发模板引擎
package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html")
	//		|
	//		|
	//		V
	// t := template.New("tmpl.html")
	// t, _ = t.ParseFiles("tmpl.html")

	// t := template.Must(template.ParseFiles("tmpl.html"))
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
/*
// 在处理器里面生成一个随机数
package main

import (
	"html/template"
	"net/http"
	// "math/rand"
	// "time"
)

func process(w http.ResponseWriter, r *http.Request) {
	// t, _ := template.ParseFiles("tmpl.html")
	// 条件动作
	// rand.Seed(time.Now().Unix())
	// t.Execute(w, rand.Intn(10) > 5)

	// 迭代动作
	// daysOfWeek := []string{}
	// daysOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	// t.Execute(w, daysOfWeek)

	// 设置动作
	// t.Execute(w, "hello")

	// 包含动作
	t, _ := template.ParseFiles("t1.html","t2.html")
	t.Execute(w ,"hello,world!")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
*/

/*
// 函数
package main

import (
	"net/http"
	"html/template"
	"time"
)

func FormatDate(t time.Time) string{
	layout := "2006-01-02"
	return t.Format(layout)
}

func process(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{"fdate": FormatDate}
	t := template.New("test.html").Funcs(funcMap)
	t, _ = t.ParseFiles("test.html")
	t.Execute(w, time.Now())
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
*/

// 上下文感知
package main

import (
	"net/http"
	"html/template"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("test.html")
	t.ExecuteTemplate(w, "layout", "")
}

func main() {
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}
	http.HandleFunc("/process",process)
	server.ListenAndServe()
}
