package slices

import (
	"reflect"
	"testing"
)

func Test_Foldl(t *testing.T) {

	res0 := foldl(func(a, b int) int { return a + b }, 0, []int{})
	if res0 != 0 {
		t.Fail()
	}

	res1 := foldl(func(a, b int) int { return a + b }, 0, []int{1, 2, 3, 4, 5})
	if res1 != 15 {
		t.Fail()
	}

	res2 := foldl(func(a, b int) int { return a * b }, 1, []int{1, 2, 3, 4, 5})
	if res2 != 120 {
		t.Fail()
	}

}

func Test_Foldr(t *testing.T) {

	res0 := foldr(func(a, b int) int { return a + b }, 0, []int{})
	if res0 != 0 {
		t.Fail()
	}

	res1 := foldr(func(a, b int) int { return a + b }, 0, []int{1, 2, 3, 4, 5})
	if res1 != 15 {
		t.Fail()
	}

	res2 := foldr(func(a, b int) int { return a * b }, 1, []int{1, 2, 3, 4, 5})
	if res2 != 120 {
		t.Fail()
	}

}

func TestFilter(t *testing.T) {

	res0 := Filter[int](func(a int) bool { return a%2 == 0 }, []int{1, 2, 3, 4, 5, 6, 7, 8, 0})
	if !reflect.DeepEqual(res0, []int{2, 4, 6, 8, 0}) {
		t.Fail()
	}

}
