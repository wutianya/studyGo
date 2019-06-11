package main

import (
	"fmt"
	"os"
	"time"

	"github.com/jlaffaye/ftp"
)

func main() {
	var buf = make([]byte, 1024)
	c, err := ftp.Dial("172.17.10.4:21", ftp.DialWithTimeout(5*time.Second))

	if err != nil {
		fmt.Println("connect", err)
	}
	err = c.Login("simpletour", "mev5Je8IeSha")
	if err != nil {
		fmt.Println("login", err)
	}
	f, _ := os.OpenFile("C:\\Users\\wutianya\\Desktop\\meituan.log", os.O_CREATE|os.O_WRONLY, 0666)
	defer f.Close()
	r, _ := c.Retr("meituan.log")

	for {
		n, _ := r.Read(buf)
		fmt.Println(n)
		if n == 0 {
			break
		}
		f.Write(buf[:n])
	}
	r.Close()
	c.Quit()
}
