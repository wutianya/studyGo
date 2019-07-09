package main

import (
	"log"
	"os"
	"io"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}
func main() {
	errFile, _ := os.OpenFile("error.log", os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0666)
	var info *log.Logger
	info = log.New(os.Stdout,"Info:",log.Ldate | log.Ltime | log.Lshortfile)
	error := log.New(io.MultiWriter(os.Stderr,errFile),"Error:",log.Ldate | log.Ltime | log.Lshortfile)
	info.Println("info")
	error.Fatal("fatal")
	log.Println("test,test")
	log.Fatal("Fatal")
	log.Panic("panic")
}