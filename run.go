package servantgo

import (
	"sync"
)

var queue = &sync.Map{}

func Run[T Task](t T) T {
	hash := t.Hash()

	task := &scheduledTask[T]{
		wg:   &sync.WaitGroup{},
		task: t,
	}
	task.wg.Add(1)

	var runningTask, ok = queue.LoadOrStore(hash, task)
	rt := runningTask.(*scheduledTask[T])

	if !ok {
		rt.task.Exec()
		rt.wg.Done()
		queue.Delete(hash)
	}
	rt.wg.Wait()

	return rt.task
}
