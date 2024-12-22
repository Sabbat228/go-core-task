package main

import (
	"testing"
	"time"
)

func TestMergeChannels(t *testing.T) {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	go func() {
		for i := 5; i < 10; i++ {
			ch2 <- i
		}
		close(ch2)
	}()

	merged := mergeChannels(ch1, ch2)

	var result []int
	timeout := time.After(1 * time.Second)

	for {
		select {
		case val, ok := <-merged:
			if !ok {
				return
			}
			result = append(result, val)
		case <-timeout:
			t.Error("Тест превышает время ожидания")
			return
		}
	}

	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	if len(result) != len(expected) {
		t.Errorf("Ожидалось %d значений, получено %d", len(expected), len(result))
	}

	for i, v := range expected {
		if result[i] != v {
			t.Errorf("Ожидалось значение %d на позиции %d, получено %d", v, i, result[i])
		}
	}
}
