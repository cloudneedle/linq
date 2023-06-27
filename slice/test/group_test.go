package test

import (
	"github.com/gocrud/linq/slice"
	"testing"
)

func TestGroup(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8}
	result := slice.Group(items, func(i int) int {
		return i % 2
	})
	t.Logf("%+v", result)
}

func TestGroupStringKey(t *testing.T) {
	items := []string{"hello", "world", "hello", "golang", "golang", "golang"}
	result := slice.Group(items, func(i string) string {
		return i
	})
	t.Logf("%+v", result)
}
