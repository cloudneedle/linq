package slice

import (
	"github.com/gocrud/linq"
	"testing"
)

func TestSortAsc(t *testing.T) {
	var items = []int{5, 4, 3, 2, 1}
	var result = linq.Sort(items, func(i, j int) bool {
		return i < j
	})
	t.Logf("%v", result) // [1 2 3 4 5]
}

func TestSortDesc(t *testing.T) {
	var items = []int{1, 2, 3, 4, 5}
	var result = linq.Sort(items, func(i, j int) bool {
		return i > j
	})
	t.Logf("%v", result) // [5 4 3 2 1]
}
