# Go by Example: Hello World

[Go by Example: Hello World](https://gobyexample.com/hello-world)

```go
// 宣告程式屬於哪個 package，所有的 go 檔案都必須聲明
package main

// 引入套件，多個可以加括號 ()
import "fmt"

// 程式執行入口，main 在 golang 中是特殊的 function，每個執行檔只能有一個
func main() {
// 使用 fmt 套件印出字串 word
 fmt.Println("hello world")
}

// hello world
```
