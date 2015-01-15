package main

import "fmt"

type cartesianPoint struct {
  x, y float64
}

// implement Point interface for cartesianPoint
func (p cartesianPoint) X() float64 { return p.x }

func (p cartesianPoint) Y() float64 { return p.y }

// implement Printer interface for Point interface for cartesianPoint
func (self cartesianPoint) Print() {
  fmt.Printf("(%f, %f)\n", self.x, self.y)
  // fmt.Printf("(%f, %f)\n", self.X(), self.Y())
}
