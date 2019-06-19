// 存储数据

/*
create table  "test".posts (
	id serial primary key,
	content text,
	author varchar(255)
);

create table "test".comments(
	id serial primary key,
	content text,
	author varchar(255),
	post_id integer references "test".posts(id)
);
*/
/*
package main

import (
	"testing/quick"
	"net/http/internal"
	"cmd/go/internal/web"
	"cmd/go/internal/str"
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
/*
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
*/

/*
// csv 文件读写操作
package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Post struct {
	Id		int
	Content	string
	Author	string
}

func main() {
	csvFile, err := os.Create("posts.csv")
	if err != nil{
		panic(err)
	}
	defer csvFile.Close()

	allPosts := []*Post{
		&Post{Id: 1, Content: "post1", Author: "Susan"},
		&Post{Id: 2, Content: "post2", Author: "peter"},
		&Post{Id: 3, Content: "post3", Author: "peter"},
		&Post{Id: 4, Content: "post4", Author: "Susan"},
	}

	writer := csv.NewWriter(csvFile)
	for _, post := range allPosts {
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()

	// open csv file
	file, err := os.Open("posts.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	fmt.Println(record)
	var posts []Post
	for _, item := range record {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		post := Post{Id: int(id), Content: item[1], Author: item[2]}
		posts = append(posts, post)
	}
	fmt.Println(posts[0].Id)
	fmt.Println(posts[0].Content)
	fmt.Println(posts[0].Author)
}
*/
/*
// gob
package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
)

type Post struct {
	Id 		int
	Content	string
	Author	string
}
// store data
func store(data interface{}, filename string) {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filename, buffer.Bytes(), 0600)
	if err != nil {
		panic(err)
	}
}
// load data
func load(data interface{}, filename string) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {

	}
	buffer := bytes.NewBuffer(raw)
	dec := gob.NewDecoder(buffer)
	err = dec.Decode(data)
	if err != nil {
		panic(err)
	}
}

func main() {
	post := Post{1, "hello world", "zsf"}
	store(post, "post1")

	var postRead Post
	load(&postRead, "post1")
	fmt.Println(postRead.Author)
}
*/

/*
// 对Postgres执行CRUD操作
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Post struct {
	Id		int
	Content	string
	Author	string
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "host=192.168.2.25 user=postgres dbname=gwp password=postgres123456 port=5432 sslmode=disable")
	if err != nil {
		panic(err)
	}
}
// 一次获取多个帖子
func Posts(limit int) (posts []Post, err error) {
	rows, err := Db.Query("select id, content, author from test.posts limit $1", limit)
	if err != nil {
		return
	}
	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.Id, &post.Content, &post.Author)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}
// 单独获取一篇帖子
func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("select * from test.posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}
// 创建一篇新帖子
func (post *Post) Create() (err error) {
	statement := "insert into test.posts (content, author) values ($1, $2) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}
// 更新帖子
func (post *Post) update() (err error) {
	_, err = Db.Exec("update test.posts set content =$2, author = $3 where id = $1", post.Id, post.Content, post.Author)
	return
}
// 删除帖子
func (post *Post) delete() (err error) {
	_, err = Db.Exec("delete from test.posts where id =$1", post.Id)
	return
}

func main() {
	// default id 0
	// post := Post{Content:"hello world!",Author:"wty"}

	// fmt.Println(post)
	// post.Create()


	// readPost, _ := GetPost(post.Id)
	// readPost.Content = "study golang."
	// readPost.Author = "wty"
	// readPost.update()

	posts, _ := Posts(2)
	fmt.Println(posts)

	post := Post{Id: 1, Content:"hello world!",Author:"wty"}
	post.Create()
	// readPost.delete()

}
*/
/*
// 实现一对多以及多对一关系
package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
)

type Comment struct {
	Id		int
	Content	string
	Author	string
	Post	*Post
}
type Post struct {
	Id		int
	Content	string
	Author	string
	Comments	[]Comment
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "host=192.168.2.25 user=postgres dbname=gwp password=postgres123456 port=5432 sslmode=disable")
	if err != nil {
		panic(err)
	}
}

// 创建一条评论
func (comment *Comment) Create() (err error) {
	if  comment.Post == nil {
		err = errors.New("Post not found")
		return
	}
	err = Db.QueryRow("insert into test.comments (content, author, post_id) values ($1, $2, $3) returning id", comment.Content, comment.Author, comment.Post.Id).Scan(&comment.Id)
	return
}
// 单独获取一篇帖子
func GetPost(id int) (post Post, err error) {
	post = Post{}
	post.Comments = []Comment{}
	err = Db.QueryRow("select id, content, author from test.posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)

	rows, err := Db.Query("select id, content, author from test.comments")
	if err != nil {
		return
	}
	for rows.Next() {
		comment := Comment{Post: &post}
		err = rows.Scan(&comment.Id, &comment.Content, &comment.Author)
		if err != nil{
			return
		}
		post.Comments = append(post.Comments, comment)
	}
	rows.Close()
	return
}

func (p *Post) Create() (err error) {
	err = Db.QueryRow("insert into test.posts (content, author) values ($1, $2) returning id", p.Content, p.Author).Scan(&p.Id)
	return
}

func main() {
	post := Post{Content: "hello world", Author: "bob"}
	post.Create()

	comment := Comment{Content: "Good post", Author: "joe", Post: &post}
	comment.Create()
	readPost, _ := GetPost(post.Id)


	fmt.Println(readPost)
	fmt.Println(readPost.Comments)
	fmt.Println(readPost.Comments[0].Post)
}
*/
/*
// 关系映射器
// github.com/jmoiron/sqlx
package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Post struct {
	Id		int
	Content	string
	AuthorName string `db:author`
}

var Db *sqlx.DB

func init() {
	var err error
	Db, err = sqlx.Open("postgres", "host=192.168.2.25 user=postgres dbname=gwp password=postgres123456 port=5432 sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRowx("select id, content, author from test.posts where id = $1", id).StructScan(&post)
	if err != nil {
		return
	}
	return
}

func (post *Post) Create() (err error) {
	err = Db.QueryRow("insert into test.posts (content, author) values ($1, $2) returning id", post.Content, post.AuthorName).Scan(&post.Id)
	return

}

func main() {
	post := Post{Content: "droir", AuthorName: "Tom"}
	post.Create()
	fmt.Println(post)
}
*/

// gorm
package main

import (
	"time"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Post struct {
	Id        int
	Content   string
	Author    string `sql:"not null"`
	Comments  []Comment
	CreatedAt time.Time
}

type Comment struct {
	Id        int
	Comtent   string
	Author    string `sql: "not null"`
	Post      int    `sql:"index"`
	CreatedAt time.Time
}

var Db *gorm.DB

func init() {
	var err error
	Db, err = gorm.Open("postgres", "host=192.168.2.25 user=postgres dbname=gwp password=postgres123456 port=5432 sslmode=disable")

	if err != nil {
		panic(err)
	}
	Db.AutoMigrate(&Post{}, & Comment{})
}

func main() {
	post := Post{Content: "fuck day.", Author:"zs"}
	fmt.Println(post)

	Db.Create(&post)
	fmt.Println(post)

	comment := Comment{Comtent:"fuck post!", Author: "Sam"}
	Db.Model(&post).Association("Comments").Append(comment)

	var readPost Post
	Db.Where("author = $1", "zs").First(&readPost)
	var comments []Comment
	Db.Model(&readPost).Related(&comments)
	fmt.Println(comments)
}