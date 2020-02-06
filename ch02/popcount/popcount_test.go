package popcount

import (
	"testing"
)

func BenchmarkPopCount(b *testing.B) {
	for i := uint64(0); i < 1<<20; i++ {
		PopCount(i)
	}
}
func BenchmarkPopCount2_3(b *testing.B) {
	for i := uint64(0); i < 1<<20; i++ {
		PopCount2_3(i)
	}
}

func BenchmarkPopCount2_4(b *testing.B) {
	for i := uint64(0); i < 1<<20; i++ {
		PopCount2_4(i)
	}
}
func BenchmarkPopCount2_5(b *testing.B) {
	for i := uint64(0); i < 1<<20; i++ {
		PopCount2_5(i)
	}
}

func TestPopCount2_3(t *testing.T) {
	for i := uint64(0); i < 1<<20; i++ {
		if PopCount(i) != PopCount2_3(i) {
			t.Fatalf("Popcount() returns %v, while PopCount2_3 returns %v", PopCount(i), PopCount2_3(i))
		}
	}
}
func TestPopCount2_4(t *testing.T) {
	for i := uint64(0); i < 1<<20; i++ {
		if PopCount(i) != PopCount2_4(i) {
			t.Fatalf("Popcount() returns %v, while PopCount2_4 returns %v", PopCount(i), PopCount2_4(i))
		}
	}
}

func TestPopCount2_5(t *testing.T) {
	for i := uint64(0); i < 1<<20; i++ {
		if PopCount(i) != PopCount2_5(i) {
			t.Fatalf("Popcount() returns %v, while PopCount2_5 returns %v", PopCount(i), PopCount2_5(i))
		}
	}
}
