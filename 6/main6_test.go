package main

import (
	"testing"
	"time"
)

func TestRandomGeneration(t *testing.T) {
	intCh := make(chan int)
	go randomGeneration(intCh)

	timeout := time.After(1 * time.Second)

	for {
		select {
		case num := <-intCh:
			if num < 0 || num >= 100 {
				t.Errorf("Сгенерированное число %d вне диапазона [0, 100)", num)
			}
		case <-timeout:
			return
		}
	}
}
