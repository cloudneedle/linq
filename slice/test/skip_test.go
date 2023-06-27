package test

import (
	"github.com/gocrud/linq/slice"
	"testing"
)

func TestSkip(t *testing.T) {
	var items = []int{1, 2, 3, 4, 5}
	var result = slice.Skip(items, 2)
	t.Log(result)
}
