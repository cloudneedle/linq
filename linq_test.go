package linq

import (
	"fmt"
	"testing"
)

type TestObject struct {
	Name string
	Age  int
}

func TestLinq(t *testing.T) {
	var testObjects = []TestObject{
		{"Alice", 18},
		{"Bob", 20},
		{"Cindy", 22},
		{"Nancy", 44},
		{"David", 24},
	}

	result := Slice(testObjects)

	sum := result.Where(func(item TestObject) bool {
		return item.Age >= 24
	}).Sum(func(item TestObject) int {
		return item.Age
	})
	t.Logf("%v", sum)
}

func TestGroup(t *testing.T) {
	var testObjects = []TestObject{
		{"Alice", 18},
		{"Bob", 20},
		{"Alice", 22},
		{"Alice", 22},
		{"Nancy", 44},
	}

	result := Slice(testObjects).Group(func(item TestObject) any {
		return item.Name
	})

	t.Logf("%v", result)
}

func TestDistinct(t *testing.T) {
	var testObjects = []TestObject{
		{"Alice", 18},
		{"Alice", 22},
		{"Alice", 22},
		{"Alice", 22},
		{"Nancy", 44},
	}

	result := Slice(testObjects).Distinct(func(item TestObject) any {
		return fmt.Sprintf("%s-%d", item.Name, item.Age)
	}).ForEach(func(to *TestObject) {
		to.Age++
	}).ToSlice()

	t.Logf("%v", result)
}

func TestWhere(t *testing.T) {
	var testObjects = []TestObject{
		{"Alice", 18},
		{"Bob", 20},
		{"Cindy", 22},
		{"Nancy", 44},
		{"David", 24},
	}

	result := Slice(testObjects).Where(func(item TestObject) bool {
		return item.Age > 30
	}).ToSlice()

	if len(result) != 1 {
		t.Errorf("expected: %d, actual: %d", 1, len(result))
	}
}

func TestCount(t *testing.T) {
	var testObjects = []TestObject{
		{"Alice", 18},
		{"Bob", 20},
		{"Cindy", 22},
		{"Nancy", 44},
		{"David", 24},
	}
	result := Slice(testObjects).Count(func(item TestObject) bool {
		return item.Age > 30
	})

	if result != 1 {
		t.Errorf("expected: %d, actual: %d", 8, result)
	}
}

func TestSingle(t *testing.T) {
	var testObjects = []TestObject{
		{"Alice", 18},
		{"Bob", 20},
		{"Cindy", 22},
		{"Nancy", 44},
		{"David", 24},
	}
	result, err := Slice(testObjects).Single(func(item TestObject) bool {
		return item.Age == 18
	})

	if err != nil {
		t.Error(err)
	}

	if result == nil {
		t.Errorf("expected: %v, actual: %v", &testObjects[0], result)
	}
}

func TestSingleOrDefault(t *testing.T) {
	var testObjects = []TestObject{
		{"Alice", 18},
		{"Bob", 20},
		{"Cindy", 22},
		{"Nancy", 44},
		{"David", 24},
	}
	result, err := Slice(testObjects).SingleOrDefault(func(item TestObject) bool {
		return item.Age == 18
	})

	if err != nil {
		t.Error(err)
	}

	if result == nil {
		t.Errorf("expected: %v, actual: %v", &testObjects[0], result)
	}
}

func TestFirst(t *testing.T) {
	var testObjects = []TestObject{
		{"Alice", 18},
		{"Bob", 20},
		{"Cindy", 22},
		{"Nancy", 44},
		{"David", 24},
	}
	result, err := Slice(testObjects).First(func(item TestObject) bool {
		return item.Age > 30
	})

	if err != nil {
		t.Error(err)
	}

	if result == nil {
		t.Errorf("expected: %v, actual: %v", &testObjects[7], result)
	}
}

func TestFirstOrDefault(t *testing.T) {
	var testObjects = []TestObject{
		{"Alice", 18},
		{"Bob", 20},
		{"Cindy", 22},
		{"Nancy", 44},
		{"David", 24},
	}
	result := Slice(testObjects).FirstOrDefault(func(item TestObject) bool {
		return item.Age > 30
	})

	if result == nil {
		t.Errorf("expected: %v, actual: %v", &testObjects[7], result)
	}
}

func TestSkip(t *testing.T) {
	var testObjects = []TestObject{
		{"Alice", 18},
		{"Bob", 20},
		{"Cindy", 22},
		{"Nancy", 44},
		{"David", 24},
	}
	result := Slice(testObjects).Skip(5).ToSlice()

	if len(result) != 0 {
		t.Errorf("expected: %d, actual: %d", 0, len(result))
	}
}

func TestTake(t *testing.T) {
	var testObjects = []TestObject{
		{"Alice", 18},
		{"Bob", 20},
		{"Cindy", 22},
		{"Nancy", 44},
		{"David", 24},
	}
	result := Slice(testObjects).Take(5).ToSlice()

	if len(result) != 5 {
		t.Errorf("expected: %d, actual: %d", 5, len(result))
	}
}

func TestForEach(t *testing.T) {
	var testObjects = []TestObject{
		{"Alice", 18},
		{"Bob", 20},
		{"Cindy", 22},
		{"Nancy", 44},
		{"David", 24},
	}
	Slice(testObjects).ForEach(func(item *TestObject) {
		item.Age++
	})

	if testObjects[0].Age != 19 {
		t.Errorf("expected: %d, actual: %d", 19, testObjects[0].Age)
	}
}

func TestOrderByAsc(t *testing.T) {
	var testObjects = []TestObject{
		{"Alice", 18},
		{"Bob", 20},
		{"Cindy", 22},
		{"Nancy", 44},
		{"David", 24},
	}
	result := Slice(testObjects).Sort(func(a, b TestObject) bool {
		return a.Age < b.Age
	}).ToSlice()

	if result[0].Age != 18 {
		t.Errorf("expected: %d, actual: %d", 18, result[0].Age)
	}

	t.Logf("%v", result)
}

func TestOrderByDesc(t *testing.T) {
	var testObjects = []TestObject{
		{"Alice", 18},
		{"Bob", 20},
		{"Cindy", 22},
		{"Nancy", 44},
		{"David", 24},
	}
	result := Slice(testObjects).Sort(func(a, b TestObject) bool {
		return a.Age > b.Age
	}).ToSlice()

	if result[0].Age != 44 {
		t.Errorf("expected: %d, actual: %d", 46, result[0].Age)
	}

	t.Logf("%v", result)
}

func TestSum(t *testing.T) {
	var testObjects = []TestObject{
		{"Alice", 18},
		{"Bob", 20},
		{"Cindy", 22},
		{"Nancy", 44},
		{"David", 24},
	}
	result := Slice(testObjects).Sum(func(item TestObject) int {
		return item.Age
	})

	if result != 128 {
		t.Errorf("expected: %d, actual: %d", 128, result)
	}

	t.Logf("%v", result)
}

func TestMax(t *testing.T) {
	var testObjects = []TestObject{
		{"Alice", 18},
		{"Bob", 20},
		{"Cindy", 22},
		{"Nancy", 44},
		{"David", 24},
	}
	result := Slice(testObjects).Max(func(item TestObject) int {
		return item.Age
	})

	if result.Age != 44 {
		t.Errorf("expected: %d, actual: %d", 46, result.Age)
	}

	t.Logf("%v", result)
}

func TestMin(t *testing.T) {
	var testObjects = []TestObject{
		{"Alice", 18},
		{"Bob", 20},
		{"Cindy", 22},
		{"Nancy", 44},
		{"David", 24},
	}
	result := Slice(testObjects).Min(func(item TestObject) int {
		return item.Age
	})

	if result.Age != 18 {
		t.Errorf("expected: %d, actual: %d", 18, result.Age)
	}

	t.Logf("%v", result)
}
