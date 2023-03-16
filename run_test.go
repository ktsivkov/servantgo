package servantgo

import (
	"sync"
	"testing"
	"time"
)

type taskMock struct {
}

func (t *taskMock) Hash() Hash {
	return "test"
}

func (t *taskMock) Exec() {
}

type loggedTaskMock struct {
	hash   Hash
	result int64
}

func (t *loggedTaskMock) Hash() Hash {
	return t.hash
}

func (t *loggedTaskMock) Exec() {
	time.Sleep(time.Millisecond * 50)
	t.result = time.Now().UnixNano()
}

func Test_Run(t *testing.T) {
	t.Run("Run two simultaneous requests", func(t *testing.T) {
		t.Run("Same task hash", func(t *testing.T) {
			wg := sync.WaitGroup{}
			wg.Add(2)
			var r1, r2 int64

			go func() {
				defer wg.Done()
				r1 = Run(&loggedTaskMock{hash: "test-hash"}).result
			}()
			go func() {
				defer wg.Done()
				r2 = Run(&loggedTaskMock{hash: "test-hash"}).result
			}()

			wg.Wait()
			if r1 != r2 {
				t.Errorf("Should have been executed only once.")
			}
		})
		t.Run("Different task hash", func(t *testing.T) {
			wg := sync.WaitGroup{}
			wg.Add(2)
			var r1, r2 int64

			go func() {
				defer wg.Done()
				r1 = Run(&loggedTaskMock{hash: "test-hash-1"}).result
			}()
			go func() {
				defer wg.Done()
				r2 = Run(&loggedTaskMock{hash: "test-hash-2"}).result
			}()

			wg.Wait()
			if r1 == r2 {
				t.Errorf("Should have been two different tasks.")
			}
		})
	})

	t.Run("Run two sequential requests", func(t *testing.T) {
		t.Run("Same task hash", func(t *testing.T) {
			var r1, r2 = Run(&loggedTaskMock{hash: "test-hash"}).result,
				Run(&loggedTaskMock{hash: "test-hash"}).result
			if r1 == r2 {
				t.Errorf("Should have been executed twice.")
			}
		})
		t.Run("Different task hash", func(t *testing.T) {
			var r1, r2 = Run(&loggedTaskMock{hash: "test-hash-1"}).result,
				Run(&loggedTaskMock{hash: "test-hash-2"}).result
			if r1 == r2 {
				t.Errorf("Should have been executed twice.")
			}
		})
	})
}

func Benchmark_Run(b *testing.B) {
	wg := &sync.WaitGroup{}
	task := &taskMock{}
	for i := 0; i < b.N; i++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			Run(task)
		}()
		go func() {
			defer wg.Done()
			Run(task)
		}()
		wg.Wait()
	}
}
