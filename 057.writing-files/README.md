# Go by Example: Writing Files

* [Go by Example: Writing Files](https://gobyexample.com/writing-files)
* [bufio](https://golang.org/pkg/bufio/)
* [bufio/#Writer.Flush](https://golang.org/pkg/bufio/#Writer.Flush)
* [io](https://golang.org/pkg/io/)
* [ioutil](https://golang.org/pkg/io/ioutil/)
* [os](https://golang.org/pkg/os/)
* [os/#File.Sync](https://golang.org/pkg/os/#File.Sync)

```go
// Writing files in Go follows similar patterns to the
// ones we saw earlier for reading.

package main

import (
  "bufio"
  "fmt"
  "io/ioutil"
  "os"
)

func check(e error) {
  if e != nil {
    panic(e)
  }
}

func main() {

  // To start, here's how to dump a string (or just
  // bytes) into a file.
  d1 := []byte("hello\ngo\n")
  err := ioutil.WriteFile("./tmp/dat1", d1, 0644)
  check(err)

  // For more granular writes, open a file for writing.
  f, err := os.Create("./tmp/dat2")
  check(err)

  // It's idiomatic to defer a `Close` immediately
  // after opening a file.
  // 結束後會呼叫來關閉檔案
  defer f.Close()

  // You can `Write` byte slices as you'd expect.
  d2 := []byte{115, 111, 109, 101, 10}
  n2, err := f.Write(d2)
  check(err)
  fmt.Printf("1. wrote %d bytes\n", n2)

  // A `WriteString` is also available.
  n3, err := f.WriteString("writes\n")
  fmt.Printf("2. wrote %d bytes\n", n3)

  // Issue a `Sync` to flush writes to stable storage.

  // Sync commits the current contents of the file to stable storage.
  // Typically, this means flushing the file system's in-memory copy of
  // recently written data to disk.
  f.Sync()

  // `bufio` provides buffered writers in addition
  // to the buffered readers we saw earlier.
  w := bufio.NewWriter(f)
  n4, err := w.WriteString("buffered\n")
  fmt.Printf("3. wrote %d bytes\n", n4)

  // Use `Flush` to ensure all buffered operations have
  // been applied to the underlying writer.
  // Flush writes any buffered data to the underlying io.Writer.
  w.Flush()
}
```