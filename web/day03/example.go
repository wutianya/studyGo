/*
	enctype: 指数据回发到服务器时浏览器使用的编码类型
		multipart/form-data编码通常用于实现文件上传
		application/x-www-form-urlencoded
		text/plain

	知识点:
		1. http.Request结构体
		2. 编码类型
		3. 处理json类型请求主体
		4. cookie的使用
		5. 利用cookie实现消息闪现
*/
/*
// 读取请求首部
package main

import (
	"encoding"
	"go/importer"
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
/*
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
*/
/*
// 通过 MultipartForm 字段接收用户上传的文件
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request){
	// r.ParseMultipartForm(1024)
	// fileHeader := r.MultipartForm.File["uploaded"][0]
	// file, err := fileHeader.Open()
	// if err == nil {
	// 	data, err := ioutil.ReadAll(file)
	// 	if err == nil {
	// 		fmt.Fprintln(w, string(data))
	// 	}
	// }
	file, _, err := r.FormFile("upload")
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.0:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
*/
/*
package main
import (
	"fmt"
	"encoding/json"
	"net/http"
)

type Post struct {
	User	string
	Threads	[]string
}

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
<head><title>go web</title></head>
<body><h1>hello world</h1></body>
</html>
`
	w.Write([]byte(str))
}

func writeHeaderExample(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "No such service, try next door")
}

func headerExample(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Location", "http://www.google.com.hk")
	w.WriteHeader(302)
}

func jsonExample(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	post := &Post{
		User: 		"Tom cat",
		Threads:	[]string{"first","second","third"},
	}
	json, _ := json.Marshal(post)
	w.Write(json)
}

func main(){
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/write", writeExample)
	http.HandleFunc("/writeheader", writeHeaderExample)
	http.HandleFunc("/redirect", headerExample)
	http.HandleFunc("/json", jsonExample)
	server.ListenAndServe()
}
*/
/*
// 向浏览器发送cookie
package main

import (
	"fmt"
	"net/http"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "go web",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    "python web",
		HttpOnly: true,
	}
	// w.Header().Set("Set-Cookie", c1.String())
	// w.Header().Set("Set-cookie", c2.String())
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)

}

func getCookie(w http.ResponseWriter, r *http.Request) {
	h := r.Header["Cookie"]
	fmt.Fprintln(w, h)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/set_cookie", setCookie)
	http.HandleFunc("/get_cookie", getCookie)
	server.ListenAndServe()
}
*/

// 使用cookie实现闪现消息
package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

func setMessage(w http.ResponseWriter, r *http.Request) {
	msg := []byte("hello, world")
	c := http.Cookie{
		Name : "flash",
		Value: base64.URLEncoding.EncodeToString(msg),
	}
	http.SetCookie(w, &c)
}
func showMessage(w http.ResponseWriter, r *http.Request) {
	c,err := r.Cookie("flash")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintln(w, "No message found")
		}
	}else {
		rc := http.Cookie{
			Name: "flash",
			MaxAge: -1,
			Expires: time.Unix(1, 0),
		}
		http.SetCookie(w, &rc)
		val, _ := base64.URLEncoding.DecodeString(c.Value)
		fmt.Fprintln(w, string(val))
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/set_message",setMessage)
	http.HandleFunc("/show_message", showMessage)
	server.ListenAndServe()
}
