package servantgo

import "sync"

type scheduledTask[T Task] struct {
	wg   *sync.WaitGroup
	mu   *sync.RWMutex
	task T
}
