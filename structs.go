package main

import (
  "fmt"
  "math"
)

func main() {
  myCircle := new(Circle)
  myCircle.rad = 7
  myCircle.Name = "Circle1"
  fmt.Println(myCircle)
  fmt.Printf("%T\n" ,myCircle)
  fmt.Println(myCircle.rad)
  fmt.Println(myCircle.Name)
}

type Circle struct{
  rad float64
  Name string
}

func (circle Circle) area() float64 {
  return math.Pi*math.Pow(circle.rad,2)
}
