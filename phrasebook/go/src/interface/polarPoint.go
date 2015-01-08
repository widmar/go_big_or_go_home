package ifeg

import "fmt"
import "math"

type polarPoint struct {
  r, t float64
}

// implement Point interface for polarPoint
func (p polarPoint) X() float64 {
  return p.r * math.Cos(p.t)
}

func (p polarPoint) Y() float64 {
  return p.r * math.Sin(p.t)
}

// implement Printer interface for Point interface for polarPoint
func (self polarPoint) Print() {
  fmt.Printf("(%f, %f degrees)\n", self.r, self.t)
}
