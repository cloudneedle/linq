package test

import (
	"github.com/gocrud/linq/slice"
	"testing"
)

func TestReduce(t *testing.T) {
	var items = []int{1, 2, 3, 4, 5}
	var result = slice.Reduce(items, 0, func(i int, j int) int { return i + j })
	if result != 15 {
		t.Fail()
	}
}
