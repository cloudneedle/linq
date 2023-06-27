package test

import (
	"github.com/gocrud/linq/slice"
	"testing"
)

func TestSelect(t *testing.T) {
	var items = []int{1, 2, 3, 4, 5}
	var result = slice.Map(items, func(i int) int { return i * 2 })
	t.Log(result)
}
