package main

type Point interface {
  Printer
  X() float64
  Y() float64
}
