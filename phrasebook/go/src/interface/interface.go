package main

import "fmt"
import "math"

type cartesianPoint struct {
  x, y float64
}

type polarPoint struct {
  r, t float64
}

func (p cartesianPoint) X() float64 { return p.x }
func (p cartesianPoint) Y() float64 { return p.y }

func (p polarPoint) X() float64 {
  return p.r * math.Cos(p.t)
}

func (p polarPoint) Y() float64 {
  return p.r * math.Sin(p.t)
}

func (self cartesianPoint) Print() {
  fmt.Printf("(%f, %f)\n", self.x, self.y)
  // fmt.Printf("(%f, %f)\n", self.X(), self.Y())
}

func (self polarPoint) Print() {
  fmt.Printf("(%f, %f degrees)\n", self.r, self.t)
}

type Point interface {
  Printer
  X() float64
  Y() float64
}

type Printer interface {
  Print()
}

func main() {
  var p1 cartesianPoint
  p1.x = 3.1
  p1.y = 0.0
  fmt.Printf("p1...")
  p1.Print()

  var p2 Point = p1
  fmt.Printf("p2...")
  p2.Print()

  p1.x = 0
  fmt.Printf("p2, after changing p1.x -> 0")
  p2.Print()
  fmt.Printf("p1, after changing p1.x -> 0")
  p1.Print()


}
