package linq

import "testing"

func TestSum(t *testing.T) {
	type Item struct {
		Value int
	}
	var items = []Item{
		{1},
		{2},
		{3},
		{4},
		{5},
	}
	var result = Sum(items, func(i Item) int { return i.Value })
	if result != 15 {
		t.Fail()
	}
}
