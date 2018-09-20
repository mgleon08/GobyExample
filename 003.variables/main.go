// In Go, _variables_ are explicitly declared and used by
// the compiler to e.g. check type-correctness of function
// calls.

package main

import "fmt"

func main() {

	// 變數必須使用 var 聲明
	var a = "hi"
	fmt.Println("a =", a)

	// 可以一次聲明多個，宣告成 int，沒有給值 int 預設是 0
	var b, c int = 1, 2
	fmt.Println("b, c =", b, c)

	// golang 是強型別，當宣告成 int，就不能給 int 以外的 type，一但定義就無法轉換成其他型別
	// var test int = "string"

	// 如果沒有宣告 type go 會自動推斷型別
	var d = true
	fmt.Println("d =", d)

	// 宣告 int 卻沒給值，預設會是 0
	var e int
	fmt.Println("e =", e)

	var g, h = true, false
	fmt.Println("g, h =", g, h)

	// 多個同時宣告和給值，可以用括號包再一起 (不能同行，不然會錯)
	var (
		i = false
		j int
		k = "hello"
	)

	fmt.Println("i, j, k =", i, j, k)

	// := 是聲明並賦值，並且系統自動推斷類型，必須放在 main function 裡面
	// := only works in functions and the lower case 't' is so that it is only visible to the package (unexported).
	f := "short" // var f string = "short"
	fmt.Println("f =", f)

	// 重複宣告，因為 c 是新的變數，因此可以透過
	aa, bb := 1, 2
	bb, cc := 3, 4
	fmt.Println(aa, bb, cc)

	// 重複宣告，因為沒有新的參數，因此會報錯
	// aa, bb := 1, 2
	// aa, bb := 3, 4
	// no new variables on left side of :=

	// 宣告並給值
	sum := float64(0)
	fmt.Println("sum =", sum)

	// 當不同 type 相加，會 error，必須先轉型
	ii := 55           //int
	jj := 67.8         //float64
	ss := ii + int(jj) //j is converted to int
	fmt.Println("ss =", ss)
}

// a = hi
// b, c = 1 2
// d = true
// e = 0
// g, h = true false
// i, j, k = false 0 hello
// f = short
// 1 3 4
// sum = 0
// ss = 122
