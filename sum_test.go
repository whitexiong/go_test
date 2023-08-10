package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	number := []int{1, 2, 3, 4, 5}
	t.Run("saying hello to perhaps", func(t *testing.T) {
		got := Sum(number)
		want := 15
		if want != got {
			t.Errorf("expected %d but got %d", want, got)
		}
	})

	t.Run("递归测试", func(t *testing.T) {
		got := RecursionSum(number)
		want := 15
		if want != got {
			t.Errorf("expected %d but got %d", want, got)
		}
	})
}

func RecursionSum(numbers []int) int {
	return func(numbers []int) int {
		if len(numbers) == 0 {
			return 0
		}
		return numbers[0] + Sum(numbers[1:])
	}(numbers)
}

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
