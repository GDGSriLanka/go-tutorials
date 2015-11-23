package main

import (
  "fmt"
  "time"
)

func main() {
  count1()
  count2()
}

func count1()  {
  for i := 0;i < 10;i++ {
    fmt.Println("Count1 :" , i)
    time.Sleep(time.Millisecond*500)
  }
}

func count2()  {
  for i := 0;i < 10;i++ {
    fmt.Println("Count2 :" , i)
    time.Sleep(time.Millisecond*600)
  }
}
