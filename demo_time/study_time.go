package main
import(
	"fmt"
	"time"
)
var week time.Duration
func main() {
	t := time.Now()
	fmt.Println(t)
	fmt.Printf("%02d.%02d.%4d\n", t.Day(),t.Month(),t.Year())
	fmt.Println(t.UTC())
	week = 60 * 60 * 24 * 7 * 1e9
	week_from_new := t.Add(time.Duration(week))
	fmt.Println("week_from_new: ",week_from_new)
	fmt.Println(t.Format(time.RFC822))
	fmt.Println(t.Format(time.ANSIC))
	fmt.Println(t.Format("02 Jan 2006 15:04"))
	fmt.Println(t.Format("2006-01-02 15:04"))
	test := t.Format("20060102")
	fmt.Printf("test type is: %T\n", test)
}