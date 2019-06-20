package main

import (
	"encoding/json"
	"net/http"
	"path"
	"strconv"
)

type Post struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	server := http.Server{
		Addr: "0.0.0.0:8080",
	}
	http.HandleFunc("/post/", handleRequest)
	server.ListenAndServe()
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
		err = handlePost(w, r)
	case "PUT":
		err = handlePut(w, r)
	case "DELETE":
		err = handleDelete(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
// call  curl -i -X GET http://127.0.0.1:8080/post/1
func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	// r.URL.Path == 1 -->/post/1
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	// 从数据库里获取数据,并将其填充到post结构中
	post, err := retrieve(id)
	if err != nil {
		return
	}
	// MarshalIndent将结构封装为由字节切片组成的JSON数据
	// 把post结构封装为json字符串
	output, err := json.MarshalIndent(&post, "", "\t")
	if err != nil {
		return
	}
	// 把json数据写入ResponseWriter中
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}
// 创建帖子的函数
// 测试: curl -i -X POST -H "Content-Type: application/json" -d 
//   '{"content": "My frist post", "author": "skyler"}' http://192.168.0.181:8080/post/
func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	len := r.ContentLength
	// create slice
	body := make([]byte, len)
	// 读取请求主体,并将其存储在字节切片中
	r.Body.Read(body)
	var post Post
	// 把切片存储的数据解封至post结构
	json.Unmarshal(body, &post)
	// 创建数据库记录
	err = post.create()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}
// update post
// call curl -i -X PUT -H "Content-Type: application/json" -d '{"content":
//  "Updated post", "author": "skyler"}' http://192.168.0.181:8080/post/1
func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}

	post, err := retrieve(id)
	if err != nil {
		return
	}
	len := r.ContentLength
	body := make([]byte, len)
	// 从请求主体中读取JSON数据
	r.Body.Read(body)
	// 把JSON数据解封至post结构中
	json.Unmarshal(body, &post)
	// 对数据库进行更新
	err = post.update()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}
// call curl -i -X DELETE http://192.168.0.181:8080/post/1
func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	id ,err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return 
	}
	post, err := retrieve(id)
	if err != nil {
		return
	}
	err = post.delete()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}
