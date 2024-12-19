package main

import (
	"reflect"
	"testing"
)

func TestCheckSlice(t *testing.T) {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}
	expected := []int{3, 64}

	resultBool, resultSlice := checkSlice(a, b)

	if resultBool == false {
		t.Errorf("check slice failed, expected %v, got %v", expected, resultBool)
	}

	if !reflect.DeepEqual(resultSlice, expected) {
		t.Errorf("check slice failed, expected %v, got %v", expected, resultSlice)
	}
}
