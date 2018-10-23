# Go by Example: Range over Channels

[Go by Example: Range over Channels](https://gobyexample.com/range-over-channels)

* 用 range 或 for 取 channel 的 data 時，必須要告知 channel 已經沒東西了，因此要用 `close(ch)`

```go
// In a [previous](range) example we saw how `for` and
// `range` provide iteration over basic data structures.
// We can also use this syntax to iterate over
// values received from a channel.

package main

import "fmt"

func main() {

	// We'll iterate over 2 values in the `queue` channel.
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// This `range` iterates over each element as it's
	// received from `queue`. Because we `close`d the
	// channel above, the iteration terminates after
	// receiving the 2 elements.
	for elem := range queue {
		fmt.Println(elem)
	}
}
```
