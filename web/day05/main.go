// 存储数据

/*
package main

import (
	"fmt"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

var PostById map[int]*Post
var PostByAuthor map[string][]*Post

func strore(post Post) {
	PostById[post.Id] = &post
	PostByAuthor[post.Author] = append(PostByAuthor[post.Author], &post)
}

func main() {
	PostById = make(map[int]*Post)
	PostByAuthor = make(map[string][]*Post)

	post1 := Post{Id: 1, Content: "post1", Author: "Susan"}
	post2 := Post{Id: 2, Content: "post2", Author: "peter"}
	post3 := Post{Id: 3, Content: "post3", Author: "peter"}
	post4 := Post{Id: 4, Content: "post4", Author: "Susan"}

	strore(post1)
	strore(post2)
	strore(post3)
	strore(post4)

	// fmt.Println(PostById[1])
	// fmt.Println(PostByAuthor)
	for k, post := range PostByAuthor["Susan"] {
		fmt.Println(post)
		fmt.Println(k)
	}
	for _, post := range PostByAuthor["peter"] {
		fmt.Println(post)
	}
}
*/

// 文件读写
package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data := []byte("hello world\n")
	err := ioutil.WriteFile("data1", data, 0644)
	if err != nil {
		panic(err)

	}
	r1, _ := ioutil.ReadFile("data1")
	fmt.Println(string(r1))

	f1, _ := os.Create("data2")
	defer f1.Close()

	bytes, _ := f1.Write(data)
	fmt.Printf("Wrote %d bytes to file\n", bytes)

	f2, _ := os.Open("data2")
	defer f2.Close()

	r2 := make([]byte, len(data))
	bytes, _ = f2.Read(r2)
	fmt.Printf("Read %d bytes from file\n", bytes)
	fmt.Println(string(r2))
}
