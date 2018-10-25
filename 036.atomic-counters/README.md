# Go by Example: Atomic Counters

[Go by Example: Atomic Counters](https://gobyexample.com/atomic-counters)

```go
// The primary mechanism for managing state in Go is
// communication over channels. We saw this for example
// with [worker pools](worker-pools). There are a few other
// options for managing state though. Here we'll
// look at using the `sync/atomic` package for _atomic
// counters_ accessed by multiple goroutines.

package main

import "fmt"
import "time"
import "sync/atomic"

func main() {

	// We'll use an unsigned integer to represent our
	// (always-positive) counter.
	var ops uint64

	// To simulate concurrent updates, we'll start 50
	// goroutines that each increment the counter about
	// once a millisecond.
	// 同時開 50 個去新增 count
	for i := 0; i < 50; i++ {
		go func() {
			for {
				// To atomically increment the counter we
				// use `AddUint64`, giving it the memory
				// address of our `ops` counter with the
				// `&` syntax.
				atomic.AddUint64(&ops, 1)

				// Wait a bit between increments.
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// Wait a second to allow some ops to accumulate.
	time.Sleep(time.Second)

	// In order to safely use the counter while it's still
	// being updated by other goroutines, we extract a
	// copy of the current value into `opsFinal` via
	// `LoadUint64`. As above we need to give this
	// function the memory address `&ops` from which to
	// fetch the value.

	//取當前的數字，在 sleep 1 秒數字會更大
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
```

另外稍微改一下，用一般的 `+1` 去比較 	`atomic.AddUint64(&ops, 1)` 就會發現，一般的 `+1` 當多個 goroutine，執行時就會發生，數字錯誤的現象，是因為可能造成同時去取得資料的狀況，但用 `atomic.AddUint64(&ops, 1)` 就會有互斥的效果

[atomic.AddUint64](https://golang.org/pkg/sync/atomic/#AddUint64)

```go
package main

import "fmt"
import "time"
import "sync/atomic"

func main() {

	var ops uint64
	total := 0
	for i := 0; i < 100000; i++ {
		go func() {
			atomic.AddUint64(&ops, 1)
			total++
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
	fmt.Println("total:", total)
}
```
