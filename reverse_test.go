package linq

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	var items = []int{1, 2, 3, 4, 5}
	var result = Reverse(items)
	if !reflect.DeepEqual(result, []int{5, 4, 3, 2, 1}) {
		t.Fail()
	}
}
