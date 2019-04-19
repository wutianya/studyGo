/*
map 是引用类型
var map1 map[keytype]valuetype
var map1 map[string]int
make 是用来初始化的 如: var map1 = make(map[string]int)
*/
/*
package main
import "fmt"

func main() {
	var mapLit map[string]int
	//var mapCreated map[string]float32
	var mapAssigned map[string]int

	mapLit = map[string]int{"one": 1, "two": 2}
	mapCreated := make(map[string]float32)
	mapAssigned = mapLit

	mapCreated["key1"] = 4.5
	mapCreated["key2"] = 3.14159
	mapAssigned["two"] = 3

	fmt.Printf("Map literal at \"one\" is: %d\n", mapLit["one"])
	fmt.Printf("Map created at \"key2\" is: %f\n", mapCreated["key2"])
	fmt.Printf("Map assigned at \"two\" is: %d\n", mapLit["two"])
	fmt.Printf("Map literal at \"ten\" is: %d\n", mapLit["ten"])
}
*/
/*
// 判断key是否存在
package main

import "fmt"

func main() {
	var d = make(map[string][]int)
	s := []int{1, 3, 1}
	d["age"] = s
	if _, ok := d["age"]; ok {
		fmt.Println(d["age"])
	}
	fmt.Println(d)
}
*/

// for range
package main

import "fmt"

func main() {
	capitals := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}
	for key := range capitals {
		fmt.Println("Map item: Capital of", key, "is", capitals[key])
	}
}
