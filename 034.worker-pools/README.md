# Go by Example: Go by Example: Worker Pools

[Go by Example: Worker Pools](https://gobyexample.com/worker-pools)

* 用 range 或 for 取 channel 的 data 時，必須要告知 channel 已經沒東西了，因此要用 `close(ch)`

```go
// In this example we'll look at how to implement
// a _worker pool_ using goroutines and channels.

package main

import "fmt"
import "time"

// Here's the worker, of which we'll run several
// concurrent instances. These workers will receive
// work on the `jobs` channel and send the corresponding
// results on `results`. We'll sleep a second per job to
// simulate an expensive task.
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	// In order to use our pool of workers we need to send
	// them work and collect their results. We make 2
	// channels for this.
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// This starts up 3 workers, initially blocked
	// because there are no jobs yet.
	// 一次建立 3 個 worker
	// 一開始是空的 channel，因此 jobs 都會被 block 住
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Here we send 5 `jobs` and then `close` that
	// channel to indicate that's all the work we have.
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	// 要記得 close
	close(jobs)

	// Finally we collect all the results of the work.
	for a := 1; a <= 5; a++ {
		<-results
	}
}
```
