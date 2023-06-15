package linq

import "testing"

func TestMax(t *testing.T) {
	var items = []int{1, 2, 3, 4, 5}
	var result = Max(items, func(i int) int { return i })
	if *result != 5 {
		t.Fail()
	}
}
