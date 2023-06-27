package test

import (
	"github.com/gocrud/linq/slice"
	"testing"
)

func TestMax(t *testing.T) {
	var items = []int{1, 2, 3, 4, 5}
	var result = slice.Max(items, func(i int) int { return i })
	if *result != 5 {
		t.Fail()
	}
}
