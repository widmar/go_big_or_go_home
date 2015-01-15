package main

import (
  "os"
  "fmt"
)

type Logger struct {
  out *os.File
}

func (l Logger) Log(msg string) {
  out := l.out

  if( out == nil ) {
    out = os.Stderr
  }
  
  fmt.Fprintf(out, "%s [%d]: %s\n", os.Args[0], os.Getpid(), msg)
}

func (l *Logger) SetOutput(out *os.File) {
  l.out = out
}

func main() {
  var l Logger
  l.Log("hi")
}
