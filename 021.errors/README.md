# Go by Example: Errors

[Go by Example: Errors](https://gobyexample.com/errors)

Go 和其他語言，像是 ruby 做 exceptions handing 處理錯誤方式很不一樣，Go 會在 function 明確的指定要 return 的 error 是什麼。

* Go 有一個專門處理 error 的 package `errors`，可以透過 import 來使用
* 在慣例中 error 會擺在最後一個 return value
* 若是沒有 error 則會回傳 nil
* 也很容易可以自己寫一個客製的 errors，只要透過實作 `Errors()` 的 method
* 在 go 中也很常見，在 if 用 inline check error，`if r, e := f1(i); e != nil`

```go
// In Go it's idiomatic to communicate errors via an
// explicit, separate return value. This contrasts with
// the exceptions used in languages like Java and Ruby and
// the overloaded single result / error value sometimes
// used in C. Go's approach makes it easy to see which
// functions return errors and to handle them using the
// same language constructs employed for any other,
// non-error tasks.

package main

import "errors"
import "fmt"

// By convention, errors are the last return value and
// have type `error`, a built-in interface.
func f1(arg int) (int, error) {
	if arg == 42 {

		// `errors.New` constructs a basic `error` value
		// with the given error message.
		return -1, errors.New("can't work with 42")

	}

	// A `nil` value in the error position indicates that
	// there was no error.
	return arg + 3, nil
}

// It's possible to use custom types as `error`s by
// implementing the `Error()` method on them. Here's a
// variant on the example above that uses a custom type
// to explicitly represent an argument error.
type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {

		// In this case we use `&argError` syntax to build
		// a new struct, supplying values for the two
		// fields `arg` and `prob`.
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {

	// The two loops below test out each of our
	// error-returning functions. Note that the use of an
	// inline error check on the `if` line is a common
	// idiom in Go code.
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	// If you want to programmatically use the data in
	// a custom error, you'll need to get the error as an
	// instance of the custom error type via type
	// assertion.
	_, e := f2(42)

	// Type assertion. 確認 type 是不是這個 interface
	// 用 *argError 是因為 pointer 才有實作這個 interface
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
```
