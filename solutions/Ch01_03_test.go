package solutions

import (
	"testing"
)

func TestNaiveJoin(t *testing.T) {
	result := NaiveJoin([]string{"hello", "world"})
	if result != "hello world" {
		t.Error(`NaiveJoin([]string{"hello", "world" != "hello world"`)
	}
}
