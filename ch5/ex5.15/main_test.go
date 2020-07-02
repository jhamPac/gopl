package main

import "testing"

func TestMin(t *testing.T) {
	expected := -1
	result := min2(3, 9, -1, 2, 6)

	if result != expected {
		t.Errorf("got %d does not equal %d", result, expected)
	}
}

func TestMax(t *testing.T) {
	expected := 7
	result := max2(-1, 3, 5, 7)

	if result != expected {
		t.Errorf("result %d does not equal expected %d", result, expected)
	}
}
