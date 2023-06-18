package slice

import (
	"github.com/gocrud/linq"
	"testing"
)

func TestWhere(t *testing.T) {
	var items = []int{1, 2, 3, 4, 5}
	var result = linq.Where(items, func(i int) bool { return i > 3 })
	t.Log(result) // [4 5]
}
