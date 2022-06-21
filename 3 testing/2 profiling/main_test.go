package main

import "testing"

func TestFibonacci(t *testing.T) {
	tables := []struct {
		n, result int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}

	for _, item := range tables {
		fib := Fibonacci(item.n)
		if fib != item.result {
			t.Errorf("Wrong result in Fibonacci, got %d, expected %d", fib, item.result)
		}
	}
}
