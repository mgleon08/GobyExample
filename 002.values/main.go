package main

import "fmt"

func main() {
	// Strings 可以用 + 來相加
	fmt.Println("go" + "lang")

	// Integers 透過 + 相加
	fmt.Printf("1+1 = %v, Type = %T\n", 1+1, 1+1)

	// 當有一邊數字是 float64，結果也會是 float64
	fmt.Printf("4/2 = %v, Type = %T\n", 4/2, 4/2)
	fmt.Printf("4.0/2.0 = %v, Type = %T\n", 4.0/2.0, 4.0/2.0)
	fmt.Printf("4.0/2 = %v, Type = %T\n", 4.0/2, 4.0/2)
	fmt.Printf("7.0/3.0 = %v, Type = %T\n", 7.0/3.0, 7.0/3.0)
	fmt.Printf("7.0/3 = %v, Type = %T\n", 7.0/3, 7.0/3)

	// boolean
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}

// golang
// 1+1 = 2, Type = int
// 4/2 = 2, Type = int
// 4.0/2.0 = 2, Type = float64
// 4.0/2 = 2, Type = float64
// 7.0/3.0 = 2.3333333333333335, Type = float64
// 7.0/3 = 2.3333333333333335, Type = float64
// false
// true
// false
