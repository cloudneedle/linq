package slice

import (
	"github.com/gocrud/linq"
	"testing"
)

func TestReduce(t *testing.T) {
	var items = []int{1, 2, 3, 4, 5}
	var result = linq.Reduce(items, 0, func(i int, j int) int { return i + j })
	if result != 15 {
		t.Fail()
	}
}
