package slice

import (
	"github.com/gocrud/linq"
	"testing"
)

func TestSelect(t *testing.T) {
	var items = []int{1, 2, 3, 4, 5}
	var result = linq.Map(items, func(i int) int { return i * 2 })
	t.Log(result)
}
