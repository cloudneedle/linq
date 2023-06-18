package linq

import "testing"

func TestSkip(t *testing.T) {
	var items = []int{1, 2, 3, 4, 5}
	var result = Skip(items, 2)
	t.Log(result)
}
