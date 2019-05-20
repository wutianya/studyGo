/*
// initFile.go 
// version 1.0
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

// initFile.go 
// version 1.1.0
package main

import (
	"fmt"
	"flag"
	"time"
	"os"
	"strings"
	"bufio"
)

var h bool
var a, f string

func init() {
	flag.BoolVar(&h, "h", false, "show help.")
	flag.StringVar(&a, "a", "wutianya2018@gmail.com", "specify to author")
	flag.StringVar(&f, "f", "", "specify to filename")
	flag.Usage = Usage
}

type Info struct {
	author 		string
	date		string
	filename	string
}

func (self *Info) instance() {
	t := time.Now()
	self.date = t.Format("2006-01-02 15:04")
	self.filename = f
	self.author = a
}

func main() {
	flag.Parse()
	if h || len(os.Args[1:]) == 0 {
		flag.Usage()
		os.Exit(0)
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

func Usage() {
	fmt.Fprintf(os.Stdout, `initfile version: initfile/v1.1.0
Usage: initfile [-f filename] [-a author]
Optios:
`)
	flag.PrintDefaults()
}

func Content() string {
	str := `
# CreateTime: %s
# Author: %s

`
	return str
}
func FristLine(filename string) (line string){
	if strings.HasSuffix(filename, ".sh") {
		line = "#!/usr/bin/bash"
	}else if strings.HasSuffix(filename, ".py") {
		line = "#!/usr/bin/env python"
	}else {
		line = "# " + filename
	}
	return
}