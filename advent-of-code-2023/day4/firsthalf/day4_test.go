package main

import (
	"testing"
)

func TestSmallInput(t *testing.T) {
	path := "../test_input.txt"
	result := resultIs(path)
	expected := 13

	if result != expected {
		t.Errorf("resultIs(path) returned %d, expected %d", result, expected)
	}
}
