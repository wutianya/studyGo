package main

import (
	"os"
	"fmt"
	"runtime"
	"path/filepath"
	
	"github.com/jlaffaye/ftp"
)
type Conn struct {
	Server		string
	User 		string
	Pass		string
	BasePath	string
}

func (c *Conn)login()(cli *ftp.ServerConn, err error){
	cli ,err = ftp.Dial(c.Server)
	if err != nil {
		return
	}
	err = cli.Login(c.User, c.Pass)
	if err != nil {
		return
	}
	return
}
func (c *Conn)echo() (filename []string, err error) {
	cli, err := c.login()
	if err != nil {
		return 
	}
	defer cli.Quit()
	filename, err = cli.NameList(c.BasePath)
	return
}
func (c *Conn)download(filename, dir string) (err error) {
	var buf = make([]byte, 1024)
	sep := getVar()
	if dir == "" {
		dir = "."
	}
	srcFile := filepath.Base(filename)
	dstFile := filepath.Join(dir + sep + srcFile)
	
	cli, err := c.login()
	if err != nil {
		return 
	}
	defer cli.Quit()
	f, err := os.OpenFile(dstFile, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Printf("The file will be download to %s\n", dstFile)
	r, _ := cli.Retr(filename)
	defer r.Close()
	for {
			n, _ := r.Read(buf)
			if n == 0 {
				break
			}
			f.Write(buf[:n])
		}
	return
}
func (c *Conn)upload(filename, dir string) (err error) {
	if dir == "" {
		dir = c.BasePath
	}else {
		dir = dir + c.BasePath
	}
	cli, err := c.login()
	if err != nil {
		return 
	}
	defer cli.Quit()
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	filename = filepath.Base(filename)

	fmt.Printf("The file will be upload to %s%s\n", dir,filename)
	err = cli.Stor(dir + filename, f)
	if err != nil {
		return 
	}
	return
}
func getVar()(sep string) {
	switch runtime.GOOS {
	case "linux":
		 sep =  "/"

	case "windows":
		sep =  "\\"
	
	}
	return 
}