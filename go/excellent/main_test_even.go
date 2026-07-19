package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	result := EvenOrOdd(8)
	if result != "Even" {
		t.Errorf("Expected Even, but got %s", result)
	}
}