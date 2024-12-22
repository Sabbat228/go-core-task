package main

import (
	"sync"
)

type SemaphoreWaitGroup struct {
	count int
	sem   chan struct{}
	mu    sync.Mutex
}

func NewSemaphoreWaitGroup(max int) *SemaphoreWaitGroup {
	return &SemaphoreWaitGroup{
		sem: make(chan struct{}, max),
	}
}

func (wg *SemaphoreWaitGroup) Add(delta int) {
	wg.mu.Lock()
	defer wg.mu.Unlock()

	wg.count += delta
	for i := 0; i < delta; i++ {
		wg.sem <- struct{}{}
	}
}

func (wg *SemaphoreWaitGroup) Done() {
	wg.mu.Lock()
	defer wg.mu.Unlock()

	if wg.count > 0 {
		wg.count--
		<-wg.sem
	}
}

func (wg *SemaphoreWaitGroup) Wait() {
	for i := 0; i < wg.count; i++ {
		<-wg.sem
	}
}
