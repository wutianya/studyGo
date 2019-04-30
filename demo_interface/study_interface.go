/*
定义接口
type Namer interface {
	Method1(param_list) return_type
	Method2(param_list) return_type
}
Namer 是接口类型
*/

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
