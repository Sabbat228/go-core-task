package main

import "fmt"

func main() {
	var a []int
	var b []int
	checkSlice(a, b)
}

func checkSlice(a []int, b []int) (bool, []int) {
	var newSlice []int
	for i, _ := range a {
		var number int = 0
		for j, _ := range b {
			if a[i] == b[j] {
				number++
			}
		}
		if number != 0 {
			newSlice = append(newSlice, a[i])
		}
	}
	fmt.Println(newSlice)
	if len(newSlice) > 0 {
		return true, newSlice
	}
	fmt.Println(newSlice)
	return false, newSlice
}
