package main

import (
	"testing"
)

func generateNaturals(n int) <-chan int {
	naturals := make(chan int)
	go func() {
		for x := 0; x <= n; x++ {
			naturals <- x
		}
		close(naturals)
	}()
	return naturals
}

func cubeNaturals(naturals <-chan int) <-chan float64 {
	cube := make(chan float64)
	go func() {
		for x := range naturals {
			cube <- float64(x * x * x)
		}
		close(cube)
	}()
	return cube
}

func TestGenerateNaturals(t *testing.T) {
	naturals := generateNaturals(10)

	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	i := 0
	for x := range naturals {
		if x != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], x)
		}
		i++
	}
}

func TestCubeNaturals(t *testing.T) {
	naturals := generateNaturals(10)
	cubes := cubeNaturals(naturals)

	expectedCubes := []float64{0, 1, 8, 27, 64, 125, 216, 343, 512, 729, 1000}
	i := 0
	for x := range cubes {
		if x != expectedCubes[i] {
			t.Errorf("Expected %f, got %f", expectedCubes[i], x)
		}
		i++
	}
}
