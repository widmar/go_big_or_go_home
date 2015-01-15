package main

import "fmt"

// show off the types
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

  var printer Printer = p1
  fmt.Printf("p1 via printer")
  printer.Print()

  // error case:
  //
  // s := "string"
  // printer = s
  // produces:
//# interfaceExample
//./interfaceExample.go:78: cannot use s (type string) as type Printer in assignment:
//	string does not implement Printer (missing Print method)
}
