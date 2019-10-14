package main

import (
	"fmt"
	"testing"
)

func IntMin(num1, num2 int) int {
	if num1 < num2 {
		return num1
	} else {
		return num2
	}
}

func TestIntMin(t *testing.T) {
	result := IntMin(2, -2)

	if result != -2 {
		t.Errorf("IntMin(2, -2) = %d; want -2", result)
	}
}

func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		num1, num2, want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
	}

	for _, test := range tests {
		name := fmt.Sprintf("%d, %d", test.num1, test.num2)
		t.Run(name, func (t *testing.T) {
			result := IntMin(test.num1, test.num2)
			if result != test.want {
				t.Errorf("IntMin(%d, %d) = %d; want %d", test.num1, test.num2, result, test.want)
			}
		})
	}
}
