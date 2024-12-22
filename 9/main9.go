package main

import (
	"fmt"
)

func main() {
	naturals := make(chan uint8)
	cube := make(chan float64)

	go func() {
		for x := 0; x <= 10; x++ {
			naturals <- uint8(x)
		}
		close(naturals)
	}()

	go func() {
		for x := range naturals {
			cube <- float64(x * x * x)
		}
		close(cube)
	}()

	for x := range cube {
		fmt.Println(x)
	}
}
