package main

import "fmt"

func main() {
	var slice1 []string
	var slice2 []string

	checkSlice(slice1, slice2)
}

func checkSlice(slice1 []string, slice2 []string) []string {
	var newSlice []string
	for i, _ := range slice1 {
		var number int = 0
		for j, _ := range slice2 {
			if slice1[i] == slice2[j] {
				number++
			}
		}
		if number == 0 {
			newSlice = append(newSlice, slice1[i])
		}
	}
	fmt.Println(newSlice)
	return newSlice
}
