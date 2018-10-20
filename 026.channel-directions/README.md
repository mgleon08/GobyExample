# Go by Example: Channel Directions

[Go by Example: Channel Directions](https://gobyexample.com/channel-directions)

channel 也可以是單向的，only send or receive data

```go
ch := make(chan int) // 雙向 channel
sendch := make(chan<- int) // 單向 channel
pings chan<- string //只能發送 data 到 channel
pings <-chan string //只能接收 chaneel 裡的 data
```

```go
// When using channels as function parameters, you can
// specify if a channel is meant to only send or receive
// values. This specificity increases the type-safety of
// the program.

package main

import "fmt"

// This `ping` function only accepts a channel for sending
// values. It would be a compile-time error to try to
// receive on this channel.

// 只能發送 data 到 pings
func ping(pings chan<- string, msg string) {
    pings <- msg
}

// The `pong` function accepts one channel for receives
// (`pings`) and a second for sends (`pongs`).

// 只能取出 pings 裡的 data
// 只能發送 data 到 pongs
func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}

func main() {
// 一開始定義就是雙向的，但也可以定義單向
    pings := make(chan string, 1)
    pongs := make(chan string, 1)
    ping(pings, "passed message")
    pong(pings, pongs)
    fmt.Println(<-pongs)
}
```
