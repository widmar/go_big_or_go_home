package main

import "fmt"
import "math"

type cartesianPoint struct {
  x, y float64
}

type polarPoint struct {
  r, t float64
}

// implement Point interface for cartesianPoint
func (p cartesianPoint) X() float64 { return p.x }

func (p cartesianPoint) Y() float64 { return p.y }

// implement Printer interface for Point interface for cartesianPoint
func (self cartesianPoint) Print() {
  fmt.Printf("(%f, %f)\n", self.x, self.y)
  // fmt.Printf("(%f, %f)\n", self.X(), self.Y())
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

/**
 * Type (contains another Type)
 */
type Point interface {
  Printer
  X() float64
  Y() float64
}

/**
 * Type capable of printing... notice Print/Printer (*er) association.
 */
type Printer interface {
  Print()
}

// show off the types
func main() {
  s := "string"
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

  var printer Printer = p1
  fmt.Printf("p1 via printer")
  printer.Print()

  // error - 
  // printer = s
  // produces:
//# interface
//./interface.go:78: cannot use s (type string) as type Printer in assignment:
//	string does not implement Printer (missing Print method)
  
}
