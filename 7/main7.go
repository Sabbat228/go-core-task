package main

import (
	"fmt"
	"time"
)

func mergeChannels(channels ...chan int) chan int {
	out := make(chan int)

	go func() {
		for _, ch := range channels {
			for val := range ch {
				out <- val
			}
		}
		close(out)
	}()

	return out
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch1 <- i
			time.Sleep(100 * time.Millisecond)
		}
		close(ch1)
	}()

	go func() {
		for i := 5; i < 10; i++ {
			ch2 <- i
			time.Sleep(100 * time.Millisecond)
		}
		close(ch2)
	}()

	merged := mergeChannels(ch1, ch2)

	for val := range merged {
		fmt.Println("Получено:", val)
	}
}
