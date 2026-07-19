package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	result := EvenOrOdd(7)
	if result != "Odd" {
		t.Errorf("Expected Odd, but got %s", result)
	}
	result2 := EvenOrOdd(8)
	if result2 != "Even" {
		t.Errorf("Expected Even, but got %s", result2)
	}
}