package main

import (
	"sync"
	"testing"
	"time"
)

func TestSemaphoreWaitGroup(t *testing.T) {
	const max = 3
	wg := NewSemaphoreWaitGroup(max)

	wg.Add(5)
	if wg.count != 5 {
		t.Errorf("expected count to be 5, got %d", wg.count)
	}

	wg.Done()
	if wg.count != 4 {
		t.Errorf("expected count to be 4, got %d", wg.count)
	}

	go func() {
		time.Sleep(100 * time.Millisecond)
		wg.Done()
	}()

	start := time.Now()
	wg.Wait()
	duration := time.Since(start)

	if duration > 200*time.Millisecond {
		t.Errorf("Wait took too long: %v", duration)
	}
}

func TestSemaphoreWaitGroup_ConcurrentAccess(t *testing.T) {
	const max = 3
	wg := NewSemaphoreWaitGroup(max)
	var wgSync sync.WaitGroup

	numGoroutines := 10
	wgSync.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wgSync.Done()
			wg.Add(1)
			time.Sleep(10 * time.Millisecond)
			wg.Done()
		}()
	}

	wgSync.Wait()
	wg.Wait()

	if wg.count != 0 {
		t.Errorf("expected count to be 0 after all Done calls, got %d", wg.count)
	}
}
