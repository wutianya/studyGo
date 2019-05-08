/*
定义接口
type Namer interface {
	Method1(param_list) return_type
	Method2(param_list) return_type
}
Namer 是接口类型
*/
/*
package main

import "fmt"

// interface
type Shaper interface {
	Area() float32
}

// struct
type Square struct {
	side float32
}
type Rectangle struct {
	length, width float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}
func (r Rectangle) Area() float32 {
	return r.length * r.width
}

func main() {
	q := &Square{5}

	r := Rectangle{5, 3}
	shapes := []Shaper{r, q}
	fmt.Println("Looping through shapes for area ...")
	for n, _ := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}
}
*/
/*
package main

import "fmt"

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	length, width float32
}
func (r Rectangle) Area() float32{
	return r.length * r.width
}
func main() {
	sq1 := new(Square)
	sq1.side = 5
	r := Rectangle{3,4}
	shapes := []Shaper{r, sq1}
	for n,_ :=range shapes {
		fmt.Println(shapes[n])
		fmt.Println(shapes[n].Area())
	}
	var ai Shaper
	ai = sq1
	fmt.Printf("The square has area: %f\n",ai.Area())
}
*/
/*
package main

import "fmt"

type stockPosition struct {
	ticker 		string
	sharePrice 	float32
	count 		float32
}
func (s stockPosition) getValue() float32 {
	return s.sharePrice * s.count
}
type car struct {
	make 	string
	model 	string
	price 	float32
}
func (c car) getValue() float32 {
	return c.price
}
type valuable interface {
	getValue() float32
}
func showValue(asset valuable) {
	fmt.Printf("Value of the asset is %f\n",asset.getValue())
}
func main() {
	var o valuable = stockPosition{"GOOG",577.20,4}
	showValue(o)
	c := car{"BMW","M3",66500}
	showValue(c)
}
*/
/*
package main
import "fmt"
type Simpler interface{
	Get() int
	Set(s int)
}
type Simple struct {
	v int
}
func (s *Simple) Set(value int) {
	s.v = value
}
func (s Simple) Get() int {
	return s.v
}

func bar(i Simpler,v int) int {
	i.Set(v)
	return i.Get()
}
func main() {
	s := new(Simple)
	a := bar(s, 10)
	fmt.Println(a)
	si := s
	si.Set(5)
	fmt.Println(si.Get())
}
*/

/*
// 接口嵌套接口
package main

type Buffer byte
type ReadWrite interface {
	Read(b Buffer) bool
	Write(b Buffer) bool
}
type Lock interface {
	Lock()
	Unlock()
}
type File interface {
	ReadWrite
	Lock
	Close()
}
func main() {
}
*/
/*
// 检测和转换接口变量的类型
package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float32
}
type Circle struct {
	radius float32
}
type Shaper interface {
	Area() float32
}
func (sq *Square) Area() float32 {
	return sq.side * sq.side
}
func (ci *Circle) Area() float32 {
	return ci.radius * ci.radius * math.Pi
}

func main() {
	var areaIntf Shaper
	sq1 := new(Square)
	sq1.side = 5

	areaIntf = sq1
	ci := Circle{3.0}
	s := Shaper(&ci)
	fmt.Println(s.Area())

	if t,ok := areaIntf.(*Square); ok {
		fmt.Println(ok)
		fmt.Printf("The type of areaIntf is: %T\n", t)
	}
	if u,ok := s.(*Circle); ok {
		fmt.Printf("The type of areaIntf is: %T\n", u)
	}else {
		fmt.Println("areaIntf does not contain a variable of type Circle")
	}
}
*/
/*
//  使用方法集与接口

package main

import (
	"fmt"
)

type List []int

func (l List) Len() int {
	return len(l)
}
func (l *List) Append(val int) {
	*l = append(*l, val)
}

type Appender interface {
	Append(int)
}

func CountInfo(a Appender, start, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

type Lener interface {
	Len() int
}

func LongEnough(l Lener) bool {
	return l.Len()*10 > 42
}
func main() {
	var lst List
	// cannot use lst (type List) as type Appender in argument to CountInfo:
	// 		List does not implement Appender (Append method has pointer receiver)
	// CountInfo(lst, 1, 10)
	if LongEnough(lst) {
		fmt.Printf("- lst is long enough\n")
	}

	plst := new(List)
	CountInfo(plst, 5, 10)
	if LongEnough(plst) {
		fmt.Printf("- plst is long enough\n")
	}
}
*/

package main

import (
	"./sort"
	"fmt"
)

func ints() {
	data := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	a := sort.IntArray(data) //conversion to type IntArray
	sort.Sort(a)
	if !sort.IsSorted(a) {
		panic("fails")
	}
	fmt.Printf("The sorted array is: %v\n", a)
}

func strings() {
	data := []string{"monday", "friday", "tuesday", "wednesday", "sunday", "thursday", "", "saturday"}
	a := sort.StringArray(data)
	sort.Sort(a)
	if !sort.IsSorted(a) {
		panic("fail")
	}
	fmt.Printf("The sorted array is: %v\n", a)
}

type day struct {
	num       int
	shortName string
	longName  string
}

type dayArray struct {
	data []*day
}

func (p *dayArray) Len() int           { return len(p.data) }
func (p *dayArray) Less(i, j int) bool { return p.data[i].num < p.data[j].num }
func (p *dayArray) Swap(i, j int)      { p.data[i], p.data[j] = p.data[j], p.data[i] }

func days() {
	Sunday    := day{0, "SUN", "Sunday"}
	Monday    := day{1, "MON", "Monday"}
	Tuesday   := day{2, "TUE", "Tuesday"}
	Wednesday := day{3, "WED", "Wednesday"}
	Thursday  := day{4, "THU", "Thursday"}
	Friday    := day{5, "FRI", "Friday"}
	Saturday  := day{6, "SAT", "Saturday"}
	data := []*day{&Tuesday, &Thursday, &Wednesday, &Sunday, &Monday, &Friday, &Saturday}
	a := dayArray{data}
	sort.Sort(&a)
	if !sort.IsSorted(&a) {
		panic("fail")
	}
	for _, d := range data {
		fmt.Printf("%s ", d.longName)
	}
	fmt.Printf("\n")
}

func main() {
	ints()
	strings()
	days()
}