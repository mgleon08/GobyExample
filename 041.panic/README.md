# Go by Example: Panic

[Go by Example: Panic](https://gobyexample.com/panic)

* [os package](https://golang.org/pkg/os/)

```go
// A `panic` typically means something went unexpectedly
// wrong. Mostly we use it to fail fast on errors that
// shouldn't occur during normal operation, or that we
// aren't prepared to handle gracefully.

package main

import "os"

func main() {

	// We'll use panic throughout this site to check for
	// unexpected errors. This is the only program on the
	// site designed to panic.
	panic("a problem")

	// A common use of panic is to abort if a function
	// returns an error value that we don't know how to
	// (or want to) handle. Here's an example of
    // `panic`king if we get an unexpected error when creating a new file.
    // 一般 function 都會 return 一個 err，在判斷 err 不是 nil 的話要做哪些處理
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
```
