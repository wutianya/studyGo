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