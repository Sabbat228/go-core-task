package main

import (
	"fmt"
	"math/rand"
)

func main() {
	intCh := make(chan int)

	go randomGeneration(intCh)

	for {
		num := <-intCh
		fmt.Println("Случайное число:", num)
	}
}

func randomGeneration(intCh chan<- int) {
	for {
		num := rand.Intn(100)
		intCh <- num
	}
}
