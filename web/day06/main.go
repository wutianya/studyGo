/*
// 对XML进行分析
package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

// 结构标签 struct tag
type Post struct {
	XMLName xml.Name	`xml:"post"`
	Id		string		`xml:"id,attr"`
	Content	string		`xml:"content"`
	Author	string		`xml:"author"`
	Xml		string		`xml:",innerxml"`
}

type  Author struct {
	Id		string	`xml:"id,attr"`
	Name	string	`xml:",chardata"`
}

func main() {
	xmlFile, err := os.Open("post1.xml")
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()
	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error reading XML data:", err)
		return
	}
	var post Post
	xml.Unmarshal(xmlData, &post)
	fmt.Println(post)
}
*/
/*
// json分析程序-01
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Post struct {
	Id       int        `json:"id"`
	Content  string     `json:"content"`
	Author   Author     `json:"author"`
	Comments []Comments `json:"comments"`
}

type Author struct {
	Id		int		`json:"id"`
	Name	string	`json:"name"`
}

type Comments struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	jsonFile, err := os.Open("post1.json")
	if err != nil {
		fmt.Printf("Error opening JSON file: %s\n", err)
		return
	}
	defer jsonFile.Close()
	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Printf("Error reading JSON file: %s\n", err)
		return
	}
	var post Post
	json.Unmarshal(jsonData, &post)
	fmt.Println(post)
}
*/

/*
// json分析程序-02
// 使用Decoder对JSON进行语言分析
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Post struct {
	Id       int        `json:"id"`
	Content  string     `json:"content"`
	Author   Author     `json:"author"`
	Comments []Comments `json:"comments"`
}

type Author struct {
	Id		int		`json:"id"`
	Name	string	`json:"name"`
}

type Comments struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	jsonFile, err := os.Open("post1.json")
	if err != nil {
		fmt.Printf("Error opening JSON file: %s\n", err)
		return
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	for {
		var post Post
		err := decoder.Decode(&post)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error deconding JSON: ", err)
			return
		}
		fmt.Println(post)
	}
}
*/
/*
// 使用MarshalIndent创建json文件
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Post struct {
	Id       int        `json:"id"`
	Content  string     `json:"content"`
	Author   Author     `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id		int		`json:"id"`
	Name	string	`json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	post := Post{
		Id:	1,
		Content: "Yes is me",
		Author: Author{
			Id: 2,
			Name: "Jreey",
		},
		Comments: []Comment{
			Comment{
				Id: 3,
				Content: "Hava a good day!",
				Author: "Adam",
			},
			Comment{
				Id: 4,
				Content: "How are you today?",
				Author: "Betty",
			},
			Comment{
				Id: 5,
				Content: "Who am i",
				Author: "Dawen",
			},
		},
	}
	// MarshalIndent函数将结构封装为由字节切片组成的JSON数据
	output, err := json.MarshalIndent(&post, "", "\t\t")
	if err != nil {
		fmt.Println("Error marshalling to JSON: ", err)
		return
	}
	err = ioutil.WriteFile("post2.json", output, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file: ", err)
		return
	}
}
*/

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Post struct {
	Id       int        `json:"id"`
	Content  string     `json:"content"`
	Author   Author     `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id		int		`json:"id"`
	Name	string	`json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	post := Post{
		Id:	1,
		Content: "Yes is me",
		Author: Author{
			Id: 2,
			Name: "Jreey",
		},
		Comments: []Comment{
			Comment{
				Id: 3,
				Content: "Hava a good day!",
				Author: "Adam",
			},
			Comment{
				Id: 4,
				Content: "How are you today?",
				Author: "Betty",
			},
			Comment{
				Id: 5,
				Content: "Who am i",
				Author: "Dawen",
			},
		},
	}
	jsonFile, err := os.Create("post.json")
	if err != nil {
		fmt.Println("Error marshalling to JSON: ", err)
		return
	}
	encoder := json.NewEncoder(jsonFile)
	err = encoder.Encode(&post)
	if err != nil {
		fmt.Println("Error encoding JSON to file: ", err)
		return
	}
}