package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	originalSlice := make([]int, 10)
	for i := range originalSlice {
		originalSlice[i] = rand.Intn(100)
	}
	fmt.Println(originalSlice)

	sliceExample(originalSlice)

	var number int
	fmt.Fscan(os.Stdin, &number)
	addElements(originalSlice, number)

	copySlice(originalSlice)
	fmt.Println(originalSlice)

	var number2 int
	fmt.Fscan(os.Stdin, &number2)
	removeElement(originalSlice, number2)
}

func sliceExample(originalSlice []int) {
	for i := range originalSlice {
		originalSlice[i] = rand.Intn(100) * 2
	}
	fmt.Println(originalSlice)
}

func addElements(originalSlice []int, number int) {
	originalSlice = append(originalSlice, number)
}

func copySlice(originalSlice []int) {
	copy := originalSlice[:]
	fmt.Println(copy)
}

func removeElement(originalSlice []int, number2 int) []int {
	newSlice := make([]int, 10)
	newSlice = append(originalSlice[:number2], originalSlice[number2+1:]...)
	fmt.Println(newSlice)
	return newSlice
}
