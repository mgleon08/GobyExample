# Go by Example: String Formatting

[Go by Example: String Formatting](https://gobyexample.com/string-formatting)

* [string](https://golang.org/pkg/strings)
* [fmt](https://golang.org/pkg/fmt/)

```go
// Go offers excellent support for string formatting in
// the `printf` tradition. Here are some examples of
// common string formatting tasks.

package main

import "fmt"
import "os"

type point struct {
  x, y int
}

func main() {

  // Go offers several printing "verbs" designed to
  // format general Go values. For example, this prints
  // an instance of our `point` struct.
  p := point{1, 2}
  fmt.Printf("1. %v\n", p)

  // If the value is a struct, the `%+v` variant will
  // include the struct's field names.
  fmt.Printf("2. %+v\n", p)

  // The `%#v` variant prints a Go syntax representation
  // of the value, i.e. the source code snippet that
  // would produce that value.
  fmt.Printf("3. %#v\n", p)

  // To print the type of a value, use `%T`.
  fmt.Printf("4. %T\n", p)

  // Formatting booleans is straight-forward.
  fmt.Printf("5. %t\n", true)

  // There are many options for formatting integers.
  // Use `%d` for standard, base-10 formatting.
  fmt.Printf("6. %d\n", 123)

  // This prints a binary representation.
  fmt.Printf("7. %b\n", 14)

  // This prints the character corresponding to the
  // given integer.
  fmt.Printf("8. %c\n", 33)

  // `%x` provides hex encoding.
  fmt.Printf("9. %x\n", 456)

  // There are also several formatting options for
  // floats. For basic decimal formatting use `%f`.
  fmt.Printf("10. %f\n", 78.9)

  // `%e` and `%E` format the float in (slightly
  // different versions of) scientific notation.
  fmt.Printf("11. %e\n", 123400000.0)
  fmt.Printf("12. %E\n", 123400000.0)

  // For basic string printing use `%s`.
  fmt.Printf("13. %s\n", "\"string\"")

  // To double-quote strings as in Go source, use `%q`.
  fmt.Printf("14. %q\n", "\"string\"")

  // As with integers seen earlier, `%x` renders
  // the string in base-16, with two output characters
  // per byte of input.
  fmt.Printf("15. %x\n", "hex this")

  // To print a representation of a pointer, use `%p`.
  fmt.Printf("16. %p\n", &p)

  // When formatting numbers you will often want to
  // control the width and precision of the resulting
  // figure. To specify the width of an integer, use a
  // number after the `%` in the verb. By default the
  // result will be right-justified and padded with
  // spaces.
  fmt.Printf("17. |%6d|%6d|\n", 12, 345)

  // You can also specify the width of printed floats,
  // though usually you'll also want to restrict the
  // decimal precision at the same time with the
  // width.precision syntax.
  fmt.Printf("18. |%6.2f|%6.2f|\n", 1.2, 3.45)

  // To left-justify, use the `-` flag.
  fmt.Printf("19. |%-6.2f|%-6.2f|\n", 1.2, 3.45)

  // You may also want to control width when formatting
  // strings, especially to ensure that they align in
  // table-like output. For basic right-justified width.
  fmt.Printf("20. |%6s|%6s|\n", "foo", "b")

  // To left-justify use the `-` flag as with numbers.
  fmt.Printf("21. |%-6s|%-6s|\n", "foo", "b")

  // So far we've seen `Printf`, which prints the
  // formatted string to `os.Stdout`. `Sprintf` formats
  // and returns a string without printing it anywhere.
  s := fmt.Sprintf("22. a %s", "string")
  fmt.Println(s)

  // You can format+print to `io.Writers` other than
  // `os.Stdout` using `Fprintf`.
  fmt.Fprintf(os.Stderr, "23. an %s\n", "error")
}
```