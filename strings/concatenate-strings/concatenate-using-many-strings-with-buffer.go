package main

import (
  "fmt"
  "bytes"
)

func main() {

  var buffer bytes.Buffer

  for i := 0; i < 10000; i++ {
    buffer.WriteString("a")
  }
  fmt.Println(buffer.String())
}