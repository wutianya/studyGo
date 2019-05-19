// flag 模块学习
/*
package main

import(
	"fmt"
	"flag"
)

func main() {
	num := flag.Int("n", 10, "number")
	str := flag.String("name", "zs", "print name")
	flag.Parse()
	fmt.Println(*num)
	fmt.Println(*str)
}
*/

package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	h    bool
	v, t bool
)

func init() {
	flag.BoolVar(&h, "h", false, "this help")
	flag.BoolVar(&v, "v", false, "show version and exit")
	flag.BoolVar(&t, "t", false, "test configuration and exit")
	flag.Usage = usage
}

func main() {
	flag.Parse()
	if h {
		flag.Usage()
	}
	if t {
		fmt.Println("test configure")
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, `nginx version: nginx/1.10.0
Usage: nginx [-hvVtTq] [-s signal] [-c filename] [-p prefix] [-g directives]
Options:
`)
	flag.PrintDefaults()
}
