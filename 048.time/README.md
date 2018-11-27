# Go by Example: Time

[Go by Example: Time](https://gobyexample.com/time)

* [time](https://golang.org/pkg/time/)

```go
// Go offers extensive support for times and durations;
// here are some examples.

package main

import "fmt"
import "time"

func main() {
  p := fmt.Println

  // We'll start by getting the current time.
  now := time.Now()
  p("1.", now)

  // You can build a `time` struct by providing the
  // year, month, day, etc. Times are always associated
  // with a `Location`, i.e. time zone.
  // yeay, month, day, hour, min, sec, ns, timezone
  then := time.Date(
    2017, 11, 27, 20, 34, 58, 651387237, time.UTC)
  p("2.", then)

  // You can extract the various components of the time
  // value as expected.
  p("3.", then.Year())
  p("4.", then.Month())
  p("5.", then.Day())
  p("6.", then.Hour())
  p("7.", then.Minute())
  p("8.", then.Second())
  p("9.", then.Nanosecond())
  p("10.", then.Location())

  // The Monday-Sunday `Weekday` is also available.
  p("11.", then.Weekday())

  // These methods compare two times, testing if the
  // first occurs before, after, or at the same time
  // as the second, respectively.
  p("12.", then.Before(now))
  p("13.", then.After(now))
  p("14.", then.Equal(now))

  // The `Sub` methods returns a `Duration` representing
  // the interval between two times.
  diff := now.Sub(then)
  p("15.", diff)

  // We can compute the length of the duration in
  // various units.
  p("16.", diff.Hours())
  p("17.", diff.Minutes())
  p("18.", diff.Seconds())
  p("19.", diff.Nanoseconds())

  // You can use `Add` to advance a time by a given
  // duration, or with a `-` to move backwards by a
  // duration.
  p("20.", then.Add(diff))
  p("21.", then.Add(-diff))
}
```