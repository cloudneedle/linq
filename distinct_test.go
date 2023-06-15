package linq

import (
	"reflect"
	"testing"
)

func TestDistinct(t *testing.T) {
	var items = []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	var result = Distinct(items)
	if !reflect.DeepEqual(result, []int{1, 2, 3, 4, 5}) {
		t.Fail()
	}
}
