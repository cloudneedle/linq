package linq

import "testing"

func TestTake(t *testing.T) {
	var items = []int{1, 2, 3, 4, 5}
	var result = Take(items, 2)
	t.Log(result)
}
