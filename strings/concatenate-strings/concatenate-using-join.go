package main

import "fmt"
import "strings"

func main() {

	//Texts to be concatenated
	firstname := "Alexandre"
	lastname := "Gama"

	//Creating a slice with texts
	text := []string{firstname, lastname}
	//Using strings.Join() method to join texts. Notice that this method will return a new string object
	newText := strings.Join(text, "")
	//This should print just the slice, without joining the texts
	fmt.Println(text)
	//This should print the joined text
	fmt.Println(newText)
}
