package main

import (
	"bytes"
	"fmt"
)

func main() {

	//In this example we wil use Buffer that will not create a new string object for each concatenation
	//Remember that for each + in the concatenation Go will create a new string object
	var buffer bytes.Buffer

	buffer.WriteString("Ale")
	buffer.WriteString("xand")
	buffer.WriteString("re")

	buffer.WriteString("Ga")
	buffer.WriteString("ma")

	//To print the awesome message just use the String() method from buffer
	fmt.Println(buffer.String())
}
