package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	result := EvenOrOdd(7)
	if result != "Odd" {
		t.Errorf("Expected Odd, but got %s", result)
	}
}