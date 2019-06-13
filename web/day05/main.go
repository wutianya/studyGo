package main

import (
	"fmt"
)

type Post struct {
	Id	int
	Content	string
	Author	string
}

var PostById map[int]*Post
var	PostByAuthor	map[string][]*Post

func strore(post Post) {
	PostById[post.Id] = &post
	PostByAuthor[post.Author] = append(PostByAuthor[post.Author],&post)
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

	fmt.Println(PostById[1])
	// fmt.Println(PostByAuthor)
	for _, post := range PostByAuthor["Susan"] {
		fmt.Println(post)
	}
	for _, post := range PostByAuthor["peter"] {
		fmt.Println(post)
	}
}