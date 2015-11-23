package main

import (
  "fmt"
  "time"
  "sync"
  "sync/atomic"
)

var wg sync.WaitGroup

func main() {
  wg.Add(2)
  go count1()
  go count2()
  wg.Wait()
}

func count1()  {
  for i := 0;i < 10;i++ {
    fmt.Println("Count1 :" , i)
    time.Sleep(time.Millisecond*500)
  }
  wg.Done()
}

func count2()  {
  for i := 0;i < 10;i++ {
    fmt.Println("Count2 :" , i)
    time.Sleep(time.Millisecond*600)
  }
  wg.Done()
}
