package slice

import (
	"github.com/gocrud/linq"
	"testing"
)

func TestMax(t *testing.T) {
	var items = []int{1, 2, 3, 4, 5}
	var result = linq.Max(items, func(i int) int { return i })
	if *result != 5 {
		t.Fail()
	}
}
