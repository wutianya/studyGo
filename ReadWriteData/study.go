/*
// 读取用户的输入
// Scanln 和 Sscanf的区别
package main
import (
	"context"
	"golang.org/x/tools/internal/lsp/source"
	"cmd/go/internal/str"
	"fmt"
)

var (
	firstName, lastName, s string
	i int
	f float32
	input = "56.12 / 2 / Go"
	format = "%f / %d / %s"
)

func main() {
	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&firstName, &lastName)
	fmt.Printf("Hi %s %s!\n", firstName, lastName)
	fmt.Sscanf(input, format ,&f,&i,&s)
	fmt.Println("From the string we read: ", f, i, s)
}
*/
/*
package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	inputReader *bufio.Reader
	input       string
	err         error
)

func main() {
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")
	input, err = inputReader.ReadString('\n')
	if err == nil {
		fmt.Printf("The input was: %s\n", input)
	}
}
*/
/*
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name: ")
	input, err := inputReader.ReadString('\n')

	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return
	}
	fmt.Printf("Your name is %s", input)
	// For Unix: test with delimiter "\n", for Windows: test with "\r\n"
	// version 1
	switch input {
	case "Philip\r\n":
		fmt.Println("Welcome Philip!")
	case "Chris\r\n":
		fmt.Println("Welcome Chris")
	case "SB\r\n":
		fmt.Println("Fuck you SB!")
	default:
		fmt.Println("You are not welcome here! Goodbye!")
	}
	// version 2
	switch input {
	case "Philip\r\n":
		fallthrough
	case "Ivo\r\n":
		fallthrough
	case "Chirs\r\n":
		fmt.Printf("Welcome %s\n", input)
	default:
		fmt.Println("You are not welcome here! Goodbye!")
	}
	// version 3
	switch input {
	case "Philip\r\n", "Ivo\r\n":
		fmt.Printf("Welcome %s\n", input)
	default:
		fmt.Println("You are not welcome here! Goodbye!")
	}
}
*/
/*
package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

var nrchars, nrwords, nrlines int

func main() {
	nrchars, nrwords, nrlines = 0, 0, 0
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input, type S to stop: ")
	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("An error occurred: %s\n", err)
		}
		if input == "S\r\n" {
			fmt.Println("Here are the counts:")
			fmt.Printf("Number of characters: %d\n",nrchars)
			fmt.Printf("Number of words: %d\n", nrwords)
			fmt.Printf("Number of lines: %d\n", nrlines)
			os.Exit(0)
		}
		Counters(input)
	}
}

func Counters(input string) {
	nrchars += len(input) -2 // -2 for \r\n
	nrwords += len(strings.Fields(input))
	nrlines++
}
*/
/*
// 读文件
// fileinput.go
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, inputError := os.Open("input.txt")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got access to it?\n")
		return // exit the function on error
	}
	// defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadBytes('\n')
		// inputString, readerError := inputReader.ReadString('\n')
		fmt.Printf("The input was: %s", inputString)
		if readerError == io.EOF {
			return
		}
	}
}
*/
/*
// read_write_file.go
package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	inputFile := "F:\\devops\\github\\go\\ReadWriteData\\a.txt"
	outputFile := "F:\\devops\\github\\go\\ReadWriteData\\b.txt"
	buf, err := ioutil.ReadFile(inputFile)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%s\n", string(buf))
	err = ioutil.WriteFile(outputFile, buf, 0644) // oct, not hex
	if err != nil {
		panic(err.Error())
	}
}
*/

/*
package main
import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filename, err := filepath.Abs("test.txt")
	if err != nil {
		panic(err)
	}
	file,err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var col1, col2, col3 []string
	for {
		var v1, v2, v3 string
		_, err := fmt.Fscanln(file, &v1, &v2, &v3)
		if err != nil {
			break
		}
		col1 =append(col1, v1)
		col2 =append(col2, v2)
		col3 =append(col3, v3)
	}
	fmt.Println(col1)
	fmt.Println(col2)
	fmt.Println(col3)
}
*/
/*
package main

import (
	"fmt"
	"bufio"
	"os"
	"compress/gzip"
)

func main() {
	fName := "F:\\devops\\github\\go\\ReadWriteData\\test.tgz"
	var r *bufio.Reader
	f, err := os.Open(fName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v, Can't open %s: error: %s\n", os.Args[0], fName, err)
		os.Exit(1)
	}
	defer f.Close()

	fz, err := gzip.NewReader(f)
	if err != nil {
		r = bufio.NewReader(f)
	} else {
		r = bufio.NewReader(fz)
	}

	for {
		line, err := r.ReadString('\n')
		// fmt.Println(err)
		if err != nil {

			fmt.Println("Done reading file")
			os.Exit(0)
		}
		fmt.Println(line)
	}
}
*/
/*
// 使用缓存写入文件,如果不使用缓存,可以直接使用f.WriteString
package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	f, err := os.OpenFile("F:\\devops\\github\\go\\ReadWriteData\\input.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("An error occurrd with opening or creation")
		return
	}
	defer f.Close()
	fWriter := bufio.NewWriter(f)
	fWriter.WriteString("test")
	fWriter.Flush()
}
*/
/*
// initFile.go
package main

import (
	"fmt"
	"time"
	"os"
	"strings"
	"bufio"
)

type Info struct {
	author 		string
	date		string
	filename	string
}

func (self *Info) instance() {
	t := time.Now()
	self.date = t.Format("2006-01-02 15:04")
	self.filename = os.Args[1]
	if len(os.Args) == 2 {
		self.author = "wutianya2018@gmail.com"
	}else if len(os.Args) == 3{
		self.author = os.Args[2]
	}
}

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Error! Please specific filename")
		os.Exit(1)
	}

	i := new(Info)
	i.instance()

	_, err := os.Stat(i.filename)
	if err == nil {
		fmt.Println("The file already exists, not initialization!")
		os.Exit(0)
	}

	str := fmt.Sprintf(Content(), i.date, i.author)

	f, err := os.OpenFile(i.filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("An error occurrd with opening or creation")
		return
	}
	defer f.Close()
	fWriter := bufio.NewWriter(f)
	f.WriteString(FristLine(i.filename) + "\n")
	f.WriteString(str)
	fWriter.Flush()

	if strings.HasSuffix(i.filename, ".sh") {
		os.Chmod(i.filename, 0755)
	}
	fmt.Println("Initialization file succeeded")
}

func Content() string {
	str := `
# CreateTime: %s
# Author: %s

`
	return str
}
func FristLine(filename string) (line string){

	// 	bash  #!/usr/bin/bash
	// 	python #!/usr/bin/env python
	// 	other #

	if strings.HasSuffix(filename, ".sh") {
		line = "#!/usr/bin/bash"
	}else if strings.HasSuffix(filename, ".py") {
		line = "#!/usr/bin/env python"
	}else {
		line = "# " + filename
	}
	return
}
*/
/*
// wiki_part.go
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (self Page) save() {
	f, err := os.OpenFile(self.Title, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("open file error")
		os.Exit(0)
	}

	// file colsed
	defer f.Close()

	fWriter := bufio.NewWriter(f)
	fWriter.WriteString(string(self.Body))
	fWriter.Flush()
}

func load(title string) {
	buf, err := ioutil.ReadFile(title)

	if err != nil {
		fmt.Fprintf(os.Stderr, "File error: %s\n", err)
	}
	fmt.Printf("%s\n", string(buf))
}
func main() {
	p := new(Page)
	p.Title = "F:\\devops\\github\\go\\ReadWriteData\\input.txt"
	p.Body = []byte{56, 57, }
	p.save()
	load(p.Title)
}
*/
/*
// wiki_part1.go
package main

import (
	"fmt"
	"io/ioutil"
)

type Page struct {
	Title string
	Body []byte
}

func (this *Page) save() (err error) {
	return ioutil.WriteFile(this.Title, this.Body, 0666)
}

func (this *Page) load(title string) (err error) {
	this.Title = title
	this.Body, err = ioutil.ReadFile(this.Title)
	return err
}

func main() {
	// page := Page{
	// 	"Page.md",
	// 	[]byte("# Page\n## Section1\nThis is section1."),
	// }
	page := new(Page)
	page.Title = "F:\\devops\\github\\go\\ReadWriteData\\Page.md"
	page.Body = []byte("# Page\n## Section1\nThis is section1.")
	page.save()

	// load
	var p Page
	p.load("F:\\devops\\github\\go\\ReadWriteData\\Page.md")
	fmt.Println(string(p.Body))

}
*/
/*
// sline example
package main

import "fmt"

func main() {
	var a = []byte{1,23}
	b:= []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	b = []byte("This is section1.")
	fmt.Println(a)
	fmt.Println(string(b))
}
*/
/*
// 文件拷贝
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	CopyFile("F:\\devops\\github\\go\\ReadWriteData\\target", "F:\\devops\\github\\go\\ReadWriteData\\Page.md")
	fmt.Println("Copy file done")
}
func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()
	return io.Copy(dst, src)
}
*/
/*
// 从命令行读取参数
package main

import (
	"fmt"
	"os"
	"strings"
)

func main(){
	who := "Alice "
	if len(os.Args) > 1 {
		who += strings.Join(os.Args[1:], ",")
	}
	fmt.Println(os.Args[1:])
	fmt.Println("Hello ",who)
}
*/
/*
package main
import (
	"flag"
	"os"
)

var NewLine =flag.Bool("n", false, "print newline")

const (
	Space = " "
	Newline = "\n"
)

func main() {

	flag.PrintDefaults()
	flag.Parse()
	var s string = ""
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += " "
			if *NewLine {
				s += Newline
			}
		}
		s += flag.Arg(i)
	}
	os.Stdout.WriteString(s)
}
*/
/*
// 用 buffer 读取文件
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func cat(r *bufio.Reader) {
	for {
		buf, err := r.ReadBytes('\n')
		fmt.Fprintf(os.Stdout, "%s", buf)
		if err == io.EOF {
			break
		}
		// fmt.Fprintf(os.Stdout, "%s", buf)
	}
	return
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}
	for i:=0; i<flag.NArg(); i++ {
		f,err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s:error reading from %s: %s\n", os.Args[0], flag.Arg(i), err.Error())
		}
		cat(bufio.NewReader(f))
	}
}
*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var numberFlag = flag.Bool("n", false, "number each line")

func cat(r *bufio.Reader) {
	i := 1
	for {
		buf, err := r.ReadBytes('\n')
		if *numberFlag {
			fmt.Fprintf(os.Stdout, "%d %s", i, buf)
			// fmt.Fprintf(os.Stdout,"%s", buf)
			i ++
		}else {
			fmt.Fprintf(os.Stdout, "%s", buf)
		}
		if err == io.EOF {
			break
		}
	}
}

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s:error reading from %s: %s\n", os.Args[0], flag.Arg(i), err.Error())
			continue
		}
		cat(bufio.NewReader(f))
	}
}