package servantgo

import "sync"

type scheduledTask[T Task] struct {
	wg   *sync.WaitGroup
	task T
}
