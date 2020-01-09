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

func BenchmarkNaiveJoinWith10Words(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NaiveJoin(Words10)
	}
}
func BenchmarkNaiveJoinWith100Words(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NaiveJoin(Words100)
	}
}
func BenchmarkNaiveJoinWith1000Words(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NaiveJoin(Words1000)
	}
}

func BenchmarkNaiveJoinWith10000Words(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NaiveJoin(Words10000)
	}
}

func BenchmarkEfficientJoinWith10Words(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EfficientJoin(Words10)
	}
}
func BenchmarkEfficientJoinWith100Words(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EfficientJoin(Words100)
	}
}
func BenchmarkEfficientJoinWith1000Words(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EfficientJoin(Words1000)
	}
}
func BenchmarkEfficientJoinWith10000Words(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EfficientJoin(Words10000)
	}
}

var Words10 = make([]string, 10)
var Words100 = make([]string, 100)
var Words1000 = make([]string, 1000)
var Words10000 = make([]string, 10000)

func init() {
	word := "word"
	for i := 0; i < 10; i++ {
		Words10[i] = word
	}
	for i := 0; i < 100; i++ {
		Words100[i] = word
	}
	for i := 0; i < 1000; i++ {
		Words1000[i] = word
	}
}
