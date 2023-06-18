package slice

import (
	"github.com/gocrud/linq"
	"testing"
)

func TestMin(t *testing.T) {
	var items = []int{1, 2, 3, 4, 5}
	var result = linq.Min(items, func(i int) int { return i })
	if *result != 1 {
		t.Fail()
	}
}
