package linq

import "testing"

func TestMin(t *testing.T) {
	var items = []int{1, 2, 3, 4, 5}
	var result = Min(items, func(i int) int { return i })
	if *result != 1 {
		t.Fail()
	}
}
