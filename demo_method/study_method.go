/*
package main

import (
	"time"
	"cmd/go/internal/str"
	"fmt"
)

type A struct {
	Name string
}
type B struct {
	Name string
}

func (getA *A) echo() {
	fmt.Println("A")
	getA.Name = "Ailis"
}
func (getB B) echo() {
	fmt.Println("B")
	getB.Name = "zs"
}
func main() {
	a := A{}
	a.echo()
	fmt.Println(a.Name)
	b := B{}
	b.echo()
	fmt.Println(b.Name)
}
*/
/*
package main
import (
	"fmt"
)

type TZ int
type Add struct {
	Num int
}

func (a *Add) add(num int) {
	a.Num += num
}
func (add *TZ) Increase(num int) {
	*add += TZ(num)
}



func main() {
	 var a TZ
	 a.Increase(100)
	 fmt.Println(a)
	 b := Add{}
	 b.add(50)
	 fmt.Println(b.Num)
}
*/
/*
// 结构体类型上方法
package main
import (
	"fmt"
)

type TwoInts struct {
	a int
	b int
}
func main() {
	var t1 TwoInts  // --> t1 := new(TwoInts)
	t1.a = 10
	t1.b = 20
	fmt.Println(t1.AddThem())
	fmt.Println(t1.AddToParam(50))

}

func (tn *TwoInts) AddThem() int {
	return tn.a + tn.b
}
// 有参数的method
func (tn *TwoInts) AddToParam(param int) int {
	return tn.a + tn.b + param
}
*/
/*
// 非结构体类型上方法
package main
import "fmt"
type IntVrctor []int

func (self IntVrctor) Sum() (s int) {
	for _,x := range self {
		 s += x
	}
	return
}

func main() {
	var i = IntVrctor{10,9,8,7}
	fmt.Println(i.Sum()) // call --> fmt.Println(IntVrctor{10,9,8,7}.sum)
}
*/
/*
// example 定义结构体 employee，它有一个 salary 字段，给这个结构体定义一个方法 giveRaise 来按照指定的百分比增加薪水
package main
import (
	"fmt"
)

type Employee struct {
	salary float64
}

func (self *Employee) giveRaise(custom float64) float64 {
	return self.salary + (self.salary * custom)
}
func main() {
	e := new(Employee)
	e.salary = 10000.0
	fmt.Println(e.giveRaise(0.2))
}
*/
/*
// method_on_time
package main
import (
	"fmt"
	"time"
)

type myTime struct {
	time.Time
}

func (t myTime) first3Chars() string {
	return t.Time.String()[0:4]
}
func main() {
	m := myTime{time.Now()}
	fmt.Println("Full time now:", m.String())
	fmt.Println(m.first3Chars())
}
*/
/*
package main

import "fmt"

type B struct {
	thing int
}

func (b *B) change() {
	b.thing = 1
}
func (b B) write() string {
	return fmt.Sprint(b)
}

func main() {
	var b1 B
	b1.change()
	fmt.Println(b1.write())
	b2 := new(B)
	b2.change()
	fmt.Println(b2.thing)
	var b3 B
	fmt.Println(b3.thing)
}
*/
/*
package main
import (
	"fmt"
)
type List []int

func (l List) Len() int { return len(l) }
func (l *List) Append(val int) { *l=append(*l, val) }

func main()  {
	var lst List
	lst.Append(2)
	fmt.Printf("%v (len: %d)\n",lst, lst.Len())

	plst :=new(List)
	plst.Append(10)
	fmt.Printf("%v (len: %d)\n",plst, plst.Len())

}
*/
/*
package main
import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}
type NamedPoint struct {
	Point
	name string
}
func (n *NamedPoint) Abs() float64 {
	return n.Point.Abs() * 100
}
func (p *Point) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}
func main() {
	n := &NamedPoint{Point{3,4}, "python"}
	fmt.Println(n.Abs())
}
*/
/*
package main

type Car struct {
	wheelCount int
}
type Mercedes struct {
	Car
}

func (n Car) numberOfWheels() {

}
func (m Mercedes) sayHiMerkel() {

}
func main ()  {
	n := new(Mercedes)
	n.numberOfWheels()
	n.sayHiMerkel()
}
*/
/*
package main

import "fmt"

type Base struct{}
type Voodoo struct {
	Base
}

func (Base) Magic() {
	fmt.Println("base magic")
}
func (self Base) MoreMagic() {
	self.Magic()
	self.Magic()
}
func (Voodoo) Magic() {
	fmt.Println("voodoo magic")
}
func main() {
	v := new(Voodoo)
	v.Magic()
	v.MoreMagic()
}
*/
/*
package main 

import (
	"fmt"
	"strconv"
)

type TwoInts struct {
	a, b int
}

func main() {
	t := new(TwoInts)
	t.a = 12
	t.b = 10
	fmt.Printf("t is: %v\n", t)
	// %#v 会给出实例的完整输出，包括它的字段
	fmt.Printf("t is: %#v\n", t)

}
func (t *TwoInts) String() string {
	return strconv.Itoa(t.a) + " " + strconv.Itoa(t.b)
}
*/
/*
package main
import "fmt"

type Celsius float64

func (c Celsius) String(s Celsius) {
	fmt.Printf("%.1fC",s)
}
func main() {
	c := new(Celsius)
	c.String(10.0)
}
*/

package main

import (
	"fmt"
	"runtime"
)

func main() {
	// var m runtime.MemStats
	m := new(runtime.MemStats)
	runtime.ReadMemStats(m)
	fmt.Printf("%d Kb\n", m.Alloc / 1024)
}