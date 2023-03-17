# ServantGo
Task execution and merging library

**ServantGo** provides a simple and idiomatic way to merge tasks of the same type that happens to run simultaneously.

## Installation + Example

Run `go get github.com/ktsivkov/servantgo`

Create a struct for your task, and ensure that it implements the `Task` interface.

```go
package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/ktsivkov/servantgo"
)

type taskMock struct {
	result time.Time
}

func (t *taskMock) Hash() servantgo.Hash {
	return "test"
}

func (t *taskMock) Exec() {
	t.result = time.Now()
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	var r1, r2 *taskMock

	go func() {
		defer wg.Done()
		r1 = servantgo.Run(&taskMock{})
	}()
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 2)
		r2 = servantgo.Run(&taskMock{})
	}()

	wg.Wait()
	fmt.Println(r1.result.String(), r2.result.String())
}
```

In this example we request two different instances of the same task to be run.

The task will take ~5 seconds to execute, and the second task will be sent for execution ~2 seconds after the first one.

Both tasks produce the same `Hash`, and therefore they are both considered identical by the library. As a result only the first one will execute, but its result will be returned to both.

Because of this the total execution time will be ~5 seconds.

**NOTE:** The result returned to all _subscribed_ parties is the same one (same pointer).