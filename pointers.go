package main

import (
  "fmt"
)

func main() {
  x := 8
  fmt.Println("X is",x)
  changeX(&x)
  fmt.Println("Now X is",x)
  fmt.Println(x)
}

func changeX(x *int) {
  *x = 10
}
