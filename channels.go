package main

import (
  "fmt"
  "time"
  "sync/atomic"
)

var counter int
var myChannel = make(chan int)
var flagChannel = make(chan bool)
var x int64

func main() {
  go addNum("First")
  go addNum("Second")
  go printChannelValue()
  <-flagChannel
}

func addNum(name string)  {
  for i := 0;i < 10; i++ {
    time.Sleep(time.Millisecond*500)
    fmt.Println(name , i)
    myChannel <- 1
    if i==9{
      atomic.AddInt64(&x, 1)
      fmt.Println(name, "is over and the flag value is",x)
    }
    if atomic.LoadInt64(&x) == 2{
      close(myChannel)
    }
  }
}

func printChannelValue()  {
  for{
    i, more := <- myChannel
    if more{
      counter += i
      fmt.Println("Counter value is",counter)
    } else {
      flagChannel <- true
      close(flagChannel)
      return
    }
  }
}
