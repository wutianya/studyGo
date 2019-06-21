/*
package main

import (
	"testing"
	"time"
)

func TestDecode(t *testing.T) {
	post, err := decode("post.json")
	if err != nil {
		t.Error(err)
	}

	if post.Id != 1{
		t.Error("Wrong id, was expecting 1 but got", post.Id)
	}
	if post.Content != "Yes is me" {
		t.Error("Wrong content", post.Content)
	}
}
func TestLongRunningTest(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping long-running test in short mode")
	}
	time.Sleep(10 * time.Second)
}
func TestEnCode(t *testing.T) {
	t.Skip("Skipping encoding for now")
}
*/
/*
// 并行测试
package main

import (
	"testing"
	"time"
)
func TestParallel_1(t *testing.T) {
	t.Parallel()
	time.Sleep(1 * time.Second)
}

func TestParallel_2(t *testing.T) {
	t.Parallel()
	time.Sleep(2 * time.Second)
}

func TestParallel_3(t *testing.T) {
	t.Parallel()
	time.Sleep(3 * time.Second)
}
*/

// 基准测试
// go test -v -cover -short -bench .
package main 

import (
	"testing"
)

func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		decode("post.json")
	}
}

func BenchmarkUnmarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		unmarshal("post.json")
	}
}