// `for` is Go's only looping construct. Here are
// three basic types of `for` loops.

package main

import "fmt"

func main() {

	// The most basic type, with a single condition.
	i := 1 // 聲明並賦值
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// A classic initial/condition/after `for` loop.
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// `for` without a condition will loop repeatedly
	// until you `break` out of the loop or `return` from
	// the enclosing function.
	for {
		fmt.Println("loop")
		break
	}

	// You can also `continue` to the next iteration of
	// the loop.
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	// 定義一個 string 的 array
	data := []string{"a", "b", "c"}

	// 用 range 做迴圈，第一個變數會是 index, 第二個是 value
	for index, value := range data {
		fmt.Printf("%d%s, ", index, value) // 0a, 1b, 2c,
	}

	for index := range data {
		fmt.Printf("%d, ", index) // 0, 1, 2,
	}

	for _, value := range data {
		fmt.Printf("%s, ", value) // a, b, c,
	}
}
