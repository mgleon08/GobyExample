// Sometimes we'll want to sort a collection by something
// other than its natural order. For example, suppose we
// wanted to sort strings by their length instead of
// alphabetically. Here's an example of custom sorts
// in Go.

package main

import "sort"
import "fmt"

// In order to sort by a custom function in Go, we need a
// corresponding type. Here we've created a `byLength`
// type that is just an alias for the builtin `[]string`
// type.
type byLength []string

// We implement `sort.Interface` - `Len`, `Less`, and
// `Swap` - on our type so we can use the `sort` package's
// generic `Sort` function. `Len` and `Swap`
// will usually be similar across types and `Less` will
// hold the actual custom sorting logic. In our case we
// want to sort in order of increasing string length, so
// we use `len(s[i])` and `len(s[j])` here.

// byLength 實作了 sort 的 interface
func (s byLength) Len() int {
	return len(s)
}

// 排序交換方式 return true 就執行兩個交換
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// 後面長度大於前面長度 return true
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

// With all of this in place, we can now implement our
// custom sort by casting the original `fruits` slice to
// `byLength`, and then use `sort.Sort` on that typed
// slice.
func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}
