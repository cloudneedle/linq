package slice

import (
	"github.com/gocrud/linq"
	"testing"
)

func TestTake(t *testing.T) {
	var items = []int{1, 2, 3, 4, 5}
	var result = linq.Take(items, 2)
	t.Log(result)
}
