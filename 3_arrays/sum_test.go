package arrays

import "testing"
import "slices"

func TestSum(t *testing.T) {
	numbers := []int{1,2,3,4,5}

	got := Sum(numbers)
	want := 15
	
	assertValuesAreEqual(t, got, want, numbers)
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}
	
	assertSlicesEqual(t, got, want)
}

func TestSumAllTails(t *testing.T) {
	t.Run("non-empty slice", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		assertSlicesEqual(t, got, want)
	})

	t.Run("empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		assertSlicesEqual(t, got, want)
	})
}

func assertSlicesEqual(t testing.TB, got, want []int) {
	t.Helper()

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertValuesAreEqual(t testing.TB, got, expected int, numbers []int ) {
	t.Helper()

	if got != expected {
		t.Errorf("got %d expected %d given %v", got, expected, numbers)
	}
}
