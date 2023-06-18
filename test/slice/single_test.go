package slice

import (
	"github.com/gocrud/linq"
	"testing"
)

func TestSingle(t *testing.T) {
	var items = []int{1, 2, 3, 4, 5}
	var result, err = linq.Single(items, func(i int) bool { return i == 3 })
	if err != nil {
		t.Fail()
	}
	t.Log(*result) // 3
}

func TestSingleErrNotItemFound(t *testing.T) {
	var items = []int{1, 2, 3, 4, 5}
	var _, err = linq.Single(items, func(i int) bool { return i == 6 })
	if err == nil {
		t.Fail()
	}
}

func TestSingleErrMultipleItemsFound(t *testing.T) {
	var items = []int{1, 2, 3, 4, 5}
	var _, err = linq.Single(items, func(i int) bool { return i > 3 })
	if err == nil {
		t.Fail()
	}
}
