package arrayslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{4, 3, 2, 1}
		got := Sum(numbers)
		want := 10
		assertSliceSum(t, got, want)
	})
	t.Run("collection of any size", func(t *testing.T) {
		number := []int{1, 2, 3, 4, 5}
		got := Sum(number)
		want := 15
		assertSliceSum(t, got, want)
	})

}

func assertSliceSum(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}

}

func TestSumAll(t *testing.T) {
	got := sumAll([]int{1, 2, 3}, []int{0, 9})
	want := []int{6, 9}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
