package main

import (
	"testing"
)

func TestCheckSlice(t *testing.T) {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}
	expected := []string{"apple", "cherry", "43", "lead", "gno1"}

	result := checkSlice(slice1, slice2)

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("check slice failed, expected %v, got %v", expected, result)
		}
	}
}
