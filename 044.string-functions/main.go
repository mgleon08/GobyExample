// The standard library's `strings` package provides many
// useful string-related functions. Here are some examples
// to give you a sense of the package.

package main

import s "strings"
import "fmt"

// We alias `fmt.Println` to a shorter name as we'll use
// it a lot below.
var p = fmt.Println

func main() {

	// Here's a sample of the functions available in
	// `strings`. Since these are functions from the
	// package, not methods on the string object itself,
	// we need pass the string in question as the first
	// argument to the function. You can find more
	// functions in the [`strings`](http://golang.org/pkg/strings/)
	// package docs.

	// 包含 'es'
	p("Contains:  ", s.Contains("test", "es"))
	// t 的次數
	p("Count:     ", s.Count("test", "t"))
	// 前綴為 te
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	// 後綴為 st
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	// e 是第幾位
	p("Index:     ", s.Index("test", "e"))
	// 用 '-' 將字串合併
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	// 重複
	p("Repeat:    ", s.Repeat("a", 5))
	// 替換所有 o -> 0
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	// 指替換第一個 o -> 0
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	// 遇到 '-' 就分開，組成 array
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	// 轉小寫
	p("ToLower:   ", s.ToLower("TEST"))
	// 轉大寫
	p("ToUpper:   ", s.ToUpper("test"))
	p()

	// Not part of `strings`, but worth mentioning here, are
	// the mechanisms for getting the length of a string in
	// bytes and getting a byte by index.
	p("Len: ", len("hello"))
	p("Char:", "hello"[1])
}

// Note that `len` and indexing above work at the byte level.
// Go uses UTF-8 encoded strings, so this is often useful
// as-is. If you're working with potentially multi-byte
// characters you'll want to use encoding-aware operations.
// See [strings, bytes, runes and characters in Go](https://blog.golang.org/strings)
// for more information.
