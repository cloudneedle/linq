package slice

import (
	"github.com/gocrud/linq"
	"testing"
)

func TestGroup(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6, 7, 8}
	result := linq.Group(items, func(i int) int {
		return i % 2
	})
	t.Logf("%+v", result)
}

func TestGroupStringKey(t *testing.T) {
	items := []string{"hello", "world", "hello", "golang", "golang", "golang"}
	result := linq.Group(items, func(i string) string {
		return i
	})
	t.Logf("%+v", result)
}
