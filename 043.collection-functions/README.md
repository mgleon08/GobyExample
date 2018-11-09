# Go by Example: Collection Functions

[Go by Example: Collection Functions](https://gobyexample.com/collection-functions)

這篇很有趣，自己寫一些 `Collection Functions`，像是 ruby 內建就有的 `include` `any` `all` `map`

* [string](https://golang.org/pkg/strings)

```go
// We often need our programs to perform operations on
// collections of data, like selecting all items that
// satisfy a given predicate or mapping all items to a new
// collection with a custom function.

// In some languages it's idiomatic to use [generic](http://en.wikipedia.org/wiki/Generic_programming)
// data structures and algorithms. Go does not support
// generics; in Go it's common to provide collection
// functions if and when they are specifically needed for
// your program and data types.

// Here are some example collection functions for slices
// of `strings`. You can use these examples to build your
// own functions. Note that in some cases it may be
// clearest to just inline the collection-manipulating
// code directly, instead of creating and calling a
// helper function.

package main

import "strings"
import "fmt"

// Index returns the first index of the target string `t`, or
// -1 if no match is found.
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// Include returns `true` if the target string t is in the
// slice.
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// Any returns `true` if one of the strings in the slice
// satisfies the predicate `f`.
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// All returns `true` if all of the strings in the slice
// satisfy the predicate `f`.
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// Filter returns a new slice containing all strings in the
// slice that satisfy the predicate `f`.
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// Map returns a new slice containing the results of applying
// the function `f` to each string in the original slice.
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func main() {

	// Here we try out our various collection functions.
	var strs = []string{"peach", "apple", "pear", "plum"}

	fmt.Println("1.", Index(strs, "pear"))

	fmt.Println("2.", Include(strs, "grape"))

	fmt.Println("3.", Any(strs, func(v string) bool {
		// 第一個字是 p 就 return true
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println("4.", All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println("5.", Filter(strs, func(v string) bool {
		// 只要包含 e 就 return true
		return strings.Contains(v, "e")
	}))

	// The above examples all used anonymous functions,
	// but you can also use named functions of the correct
	// type.
  // string 轉大寫
	fmt.Println("6.", Map(strs, strings.ToUpper))

}
```
