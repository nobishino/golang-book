/*
プログラミング言語Goの練習問題の答えを書いていくパッケージ
*/
package solutions

import (
	"fmt"
	"testing"
)

func ExampleNaiveJoin() {
	result := NaiveJoin([]string{"hello", "world"})
	fmt.Println(result)
	//Output:
	//hello world
}
func TestNaiveJoin(t *testing.T) {
	result := NaiveJoin([]string{"hello", "world"})
	if result != "hello world" {
		t.Error(`NaiveJoin([]string{"hello", "world" != "hello world"`)
	}
}

func BenchmarkNaiveJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NaiveJoin(massiveWords)
	}
}

func ExampleEfficientJoin() {
	result := EfficientJoin([]string{"hello", "world"})
	fmt.Println(result)
	//Output:
	//hello world
}
func TestEfficientJoin(t *testing.T) {
	result := EfficientJoin([]string{"hello", "world"})
	if result != "hello world" {
		t.Error(`NaiveJoin([]string{"hello", "world" != "hello world"`)
	}
}

func BenchmarkEfficientJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EfficientJoin(massiveWords)
	}
}

var massiveWords = []string{
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
	"word",
}
