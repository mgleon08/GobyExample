# Go by Example: Time Formatting / Parsing

* [Go by Example: Time Formatting / Parsing](https://gobyexample.com/time-formatting-parsing)
* [time](https://golang.org/pkg/time/)

```go
// Go supports time formatting and parsing via
// pattern-based layouts.

package main

import "fmt"
import "time"

func main() {
  p := fmt.Println

  // Here's a basic example of formatting a time
  // according to RFC3339, using the corresponding layout
  // constant.
  t := time.Now()
  p("1.", t.Format(time.RFC3339))

  // Time parsing uses the same layout values as `Format`.
  t1, e := time.Parse(
    time.RFC3339,
    "2012-11-01T22:08:41+00:00")
  p("2.", t1)

  // `Format` and `Parse` use example-based layouts. Usually
  // you'll use a constant from `time` for these layouts, but
  // you can also supply custom layouts. Layouts must use the
  // reference time `Mon Jan 2 15:04:05 MST 2006` to show the
  // pattern with which to format/parse a given time/string.
  // The example time must be exactly as shown: the year 2006,
  // 15 for the hour, Monday for the day of the week, etc.
  // 給定一個範本，根據範本輸出時間的格式
  p("3.", t.Format("3:04PM"))
  p("4.", t.Format("Mon Jan _2 15:04:05 2006"))
  p("5.", t.Format("2006-01-02T15:04:05.999999-07:00"))
  form := "3 41 PM"
  t2, e := time.Parse(form, "8 41 PM")
  p("6.", t2, "/", e)

  // For purely numeric representations you can also
  // use standard string formatting with the extracted
  // components of the time value.
  fmt.Printf("7. %d-%02d-%02dT%02d:%02d:%02d-00:00\n",
    t.Year(), t.Month(), t.Day(),
    t.Hour(), t.Minute(), t.Second())

  // `Parse` will return an error on malformed input
  // explaining the parsing problem.
  ansic := "Mon Jan _2 15:04:05 2006"
  _, e = time.Parse(ansic, "8:41PM")
  p("8.", e)
}
```