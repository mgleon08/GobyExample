# Go by Example: Number Parsing

* [Go by Example: Number Parsing](https://gobyexample.com/number-parsing)
* [strconv](https://golang.org/pkg/strconv/)

```go
// Parsing numbers from strings is a basic but common task
// in many programs; here's how to do it in Go.

package main

// The built-in package `strconv` provides the number
// parsing.
import "strconv"
import "fmt"

func main() {

  // With `ParseFloat`, this `64` tells how many bits of
  // precision to parse.
  f, _ := strconv.ParseFloat("1.234", 64)
  fmt.Println("1.", f)

  // For `ParseInt`, the `0` means infer the base from
  // the string. `64` requires that the result fit in 64
  // bits.
  i, _ := strconv.ParseInt("123", 0, 64)
  fmt.Println("2.", i)

  // `ParseInt` will recognize hex-formatted numbers.
  d, _ := strconv.ParseInt("0x1c8", 0, 64)
  fmt.Println("3.", d)

  // A `ParseUint` is also available.
  // ParseUint is like ParseInt but for unsigned numbers.
  u, _ := strconv.ParseUint("789", 0, 64)
  fmt.Println("4.", u)

  // `Atoi` is a convenience function for basic base-10
  // `int` parsing.
  // Atoi returns the result of ParseInt(s, 10, 0) converted to type int.
  k, _ := strconv.Atoi("135")
  fmt.Println("5.", k)

  // Parse functions return an error on bad input.
  _, e := strconv.Atoi("wat")
  fmt.Println("6.", e)
}
```
