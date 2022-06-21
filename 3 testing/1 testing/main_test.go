package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(5, 5)
	const expected int = 10
	if total != 10 {
		t.Errorf("Incorrect sum, got %d expected %d", total, expected)
	}

	// Create a table of type slice with all test cases
	tables := []struct {
		a, b, n int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{5, 5, 10},
	}

	for _, value := range tables {
		total := Sum(value.a, value.b)
		if total != value.n {
			t.Errorf("got %d, expected %d", total, value.n)
		}
	}
}

func TestGetMax(t *testing.T) {
	tables := []struct {
		a, b, max int
	}{
		{4, 2, 4},
		{1, 2, 2},
		{100, 10, 100},
	}

	for _, value := range tables {
		max := GetMax(value.a, value.b)
		if max != value.max {
			t.Errorf("GetMax error: got %d, expected %d", max, value.max)
		}
	}

}
