package test

import (
	"github.com/gocrud/linq/slice"
	"reflect"
	"testing"
)

func TestDistinct(t *testing.T) {
	var items = []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5}
	var result = slice.Distinct(items, func(item int) int {
		return item
	})
	if !reflect.DeepEqual(result, []int{1, 2, 3, 4, 5}) {
		t.Fail()
	}
}

func TestDistinctObject(t *testing.T) {
	type Item struct {
		Name  string
		Value int
	}
	var items = []Item{
		{"A", 1},
		{"B", 2},
		{"C", 3},
		{"D", 4},
		{"E", 5},
		{"A", 1},
		{"B", 2},
		{"C", 3},
		{"D", 4},
		{"E", 5},
	}
	var result = slice.Distinct(items, func(item Item) string {
		return item.Name
	})
	if !reflect.DeepEqual(result, []Item{
		{"A", 1},
		{"B", 2},
		{"C", 3},
		{"D", 4},
		{"E", 5},
	}) {
		t.Fail()
	}
	t.Logf("%+v", result)
}
