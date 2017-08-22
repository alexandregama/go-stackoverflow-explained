package main

import "fmt"

func main() {

  text := ""
  for i := 0; i < 10000; i++ {
    text += getTextToBeConcatenated()
  }
  fmt.Println(text)
}

func getTextToBeConcatenated() string {
  return "a"
}