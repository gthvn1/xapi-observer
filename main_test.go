package main

import "testing"

func TestSum(t *testing.T) {
	result := Sum(2, 3)
	expected := 5

	if result != expected {
		t.Fatalf("expected %d, got %d", expected, result)
	}
}
