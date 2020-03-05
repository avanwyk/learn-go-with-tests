package sum

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("should sum slice", func(t *testing.T) {
		ns := []int{1, 2, 3, 4, 5, 6, 7}
		actual := Sum(ns)
		expected := 28

		if actual != expected {
			t.Errorf("Expected sum: %d, actual sum: %d, given %v", expected, actual, ns)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("should sum all slices", func(t *testing.T) {
		actual := SumAll([]int{1, 2}, []int{2, 3})
		expected := []int{3, 5}
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Expected %v, actual %v", expected, actual)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("should sum all tails", func(t *testing.T) {
		actual := SumAllTails([]int{1, 2}, []int{2, 3, 4}, []int{4, 5, 6, 7})
		expected := []int{2, 7, 18}
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Expected %v, actual %v", expected, actual)
		}
	})

	t.Run("should safely sum all tails", func(t *testing.T) {
		actual := SumAllTails([]int{1, 2}, []int{}, []int{4, 5, 6, 7})
		expected := []int{2, 0, 18}
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Expected %v, actual %v", expected, actual)
		}
	})
}
