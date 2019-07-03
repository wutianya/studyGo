package main

import (
	"os"
	"fmt"
	"flag"
)


var h,l bool
var d,get, put string
func init() {
	flag.BoolVar(&h, "h", false, "show help")
	flag.BoolVar(&l, "l", false, "show ftp file list")
	flag.StringVar(&get, "get","","download file from ftp")
	flag.StringVar(&put, "put", "", "upload fiel from ftp")
	flag.StringVar(&d, "d", "", "specify action dir")
	flag.Usage = Usage
}
func Usage() {
	fmt.Fprintf(os.Stdout, `goftp version: v0.0.1
Usage: goftp [options] filename targetdir.
Options:
`)
	flag.PrintDefaults()
}
func main() {
	c := Conn{"172.17.10.4:21","simpletour","mev5Je8IeSha","/"}
	var fileList []string
	var err error
	flag.Parse()
	if h || len(os.Args[1:]) == 0 {
		flag.Usage()
		os.Exit(0)
	}
	switch  {
	case l:
		fileList,err = c.echo()
		if err != nil {
			panic(err)
		}
		for i := 0; i < len(fileList); i++ {
			fmt.Println(fileList[i])
		}
		os.Exit(0)
	case get != "":
		err = c.download(get,d)
		if err != nil {
			panic(err)
		}
		os.Exit(0)
	case put != "":
		err = c.upload(put,d)
		if err != nil {
			panic(err)
		}
		os.Exit(0)
	}

}

